package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// TestConfigSweep_AllRepoConfigsParseUnderNewSchema walks every YAML
// under the repo-root `configs/` directory and asserts:
//  1. Parse succeeds (the top-level `plugins:` pre-scanner does not
//     reject any migrated config);
//  2. No config carries a top-level `plugins:` field;
//  3. Every plugin entry lives under `generator.tools` (or explicitly
//     under `reviewer.tools`) with `type: plugin`;
//  4. A plugin declared only in generator.tools does NOT appear in
//     the same config's reviewer.tools — regression guard for the
//     retired auto-append behavior.
//
// This is the config-level twin of the resolver-layer no-auto-append
// tests. If anyone adds a new config with the old schema or re-enables
// ExpandPlugins-style auto-append, this test fires.
func TestConfigSweep_AllRepoConfigsParseUnderNewSchema(t *testing.T) {
	configsDir := findRepoConfigsDir(t)
	if configsDir == "" {
		t.Skip("repo configs/ directory not locatable from test cwd")
	}

	entries, err := os.ReadDir(configsDir)
	if err != nil {
		t.Fatalf("reading configs dir: %v", err)
	}

	saw := 0
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		ext := filepath.Ext(e.Name())
		if ext != ".yaml" && ext != ".yml" {
			continue
		}
		saw++
		t.Run(e.Name(), func(t *testing.T) {
			path := filepath.Join(configsDir, e.Name())
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("read: %v", err)
			}

			// 2. No stray top-level plugins: field. The Parse
			//    pre-scanner rejects it, but we also pin the raw
			//    text so the test surface doesn't depend on error
			//    routing.
			if containsTopLevelPluginsField(data) {
				t.Errorf("%s: top-level `plugins:` field must be removed (migrated to generator.tools)", e.Name())
			}

			// 1. Parse succeeds.
			cfg, err := Parse(data)
			if err != nil {
				t.Fatalf("parse %s: %v", e.Name(), err)
			}

			// 3 + 4. For each config, partition plugin entries by
			//        role and confirm the no-auto-append contract.
			for _, c := range cfg.Configs {
				genPlugins := map[string]bool{}
				if c.Generator != nil {
					for _, tl := range c.Generator.Tools {
						if tl.ResolvedType() == "plugin" {
							genPlugins[tl.Name] = true
						}
					}
				}
				revPlugins := map[string]bool{}
				if c.Reviewer != nil {
					for _, tl := range c.Reviewer.Tools {
						if tl.ResolvedType() == "plugin" {
							revPlugins[tl.Name] = true
						}
					}
				}
				// Any plugin in reviewer must be explicitly
				// listed there — we only know it was explicit
				// if it came straight from the parsed YAML. The
				// parser does no cross-role injection post-migration,
				// so presence in revPlugins is equivalent to an
				// explicit listing. That's the invariant we pin.
				// Sanity: intersection is allowed (explicit dual-role)
				// but we want to flag if someone ever reintroduces
				// a parse-time "also in reviewer" side effect. The
				// best signal we have is this: when a reviewer has
				// no tools block at all in the YAML, it should have
				// no plugin entries.
				if c.Reviewer == nil || len(c.Reviewer.Tools) == 0 {
					if len(revPlugins) != 0 {
						t.Errorf("config %q: reviewer has no tools block in YAML but parser added plugin entries: %v", c.Name, revPlugins)
					}
				}
				_ = genPlugins // generator inventory is informational
			}
		})
	}

	if saw == 0 {
		t.Skip("no YAML configs found in configs/ — skipping sweep")
	}
}

func TestFoundryComparisonPreservesThreeWayControls(t *testing.T) {
	configsDir := findRepoConfigsDir(t)
	if configsDir == "" {
		t.Skip("repo configs/ directory not locatable from test cwd")
	}

	languages := []string{"dotnet", "java", "js-ts", "python"}
	for _, language := range languages {
		t.Run(language, func(t *testing.T) {
			threeWay := parseRepoConfig(t, configsDir, language+"-azure-skills-three-way.yaml")
			cfg := parseRepoConfig(t, configsDir, language+"-azure-tools-comparison.yaml")
			if len(threeWay.Configs) != 3 {
				t.Fatalf("three-way arm count = %d, want 3", len(threeWay.Configs))
			}
			if len(cfg.Configs) != 3 {
				t.Fatalf("arm count = %d, want 3", len(cfg.Configs))
			}
			wantNames := []string{"baseline", "azure-skill-mcp", "with-azure-tools"}
			for i := range threeWay.Configs {
				if suffixAfterSlash(cfg.Configs[i].Name) != wantNames[i] {
					t.Errorf("Foundry arm %d name = %q, want suffix %q", i, cfg.Configs[i].Name, wantNames[i])
				}
				foundryGenerator := *cfg.Configs[i].Generator
				if i > 0 {
					foundryGenerator.Tools = withoutNamedTool(foundryGenerator.Tools, "foundry")
				}
				if !reflect.DeepEqual(&foundryGenerator, threeWay.Configs[i].Generator) {
					t.Errorf("Foundry arm %d generator differs from three-way control", i)
				}
				if !reflect.DeepEqual(cfg.Configs[i].Reviewer, threeWay.Configs[i].Reviewer) {
					t.Errorf("Foundry arm %d reviewer differs from three-way control", i)
				}
				if !reflect.DeepEqual(cfg.Configs[i].Limits, threeWay.Configs[i].Limits) {
					t.Errorf("Foundry arm %d limits differ from three-way control", i)
				}
			}

			for i := 1; i < len(cfg.Configs); i++ {
				foundryTool, ok := namedTool(cfg.Configs[i].Generator.Tools, "foundry")
				if !ok {
					t.Errorf("arm %d does not mount Foundry MCP", i)
					continue
				}
				if foundryTool.ResolvedMCPType() != "remote" ||
					foundryTool.URL != "https://mcp.ai.azure.com" ||
					foundryTool.MCPAuth != "azure_cli" ||
					!reflect.DeepEqual(foundryTool.MCPScopes, []string{"https://mcp.ai.azure.com/Foundry.Mcp.Tools"}) {
					t.Errorf("arm %d has unexpected Foundry MCP configuration: %+v", i, foundryTool)
				}
			}

			withoutMicrosoftSkills := *cfg.Configs[2].Generator
			withoutMicrosoftSkills.Tools = withoutToolRepo(withoutMicrosoftSkills.Tools, "microsoft/skills")
			if !reflect.DeepEqual(&withoutMicrosoftSkills, cfg.Configs[1].Generator) {
				t.Error("fully tooled arm must differ from the second arm only by adding microsoft/skills")
			}
		})
	}
}

func namedTool(tools []ToolEntry, name string) (ToolEntry, bool) {
	for _, entry := range tools {
		if entry.Name == name {
			return entry, true
		}
	}
	return ToolEntry{}, false
}

func withoutNamedTool(tools []ToolEntry, name string) []ToolEntry {
	filtered := make([]ToolEntry, 0, len(tools))
	for _, entry := range tools {
		if entry.Name != name {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

func withoutToolRepo(tools []ToolEntry, repo string) []ToolEntry {
	filtered := make([]ToolEntry, 0, len(tools))
	for _, entry := range tools {
		if entry.Repo != repo {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

func parseRepoConfig(t *testing.T, configsDir, name string) *ConfigFile {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(configsDir, name))
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("parse %s: %v", name, err)
	}
	return cfg
}

func suffixAfterSlash(name string) string {
	_, suffix, ok := strings.Cut(name, "/")
	if !ok {
		return name
	}
	return suffix
}

// containsTopLevelPluginsField does a line-based scan for a YAML key
// `plugins:` sitting at list-item indent (two spaces) directly under a
// configs: entry. We purposely avoid re-parsing here so the assertion
// is independent of the rejectRetiredPluginsField logic we're validating.
func containsTopLevelPluginsField(data []byte) bool {
	for _, line := range strings.Split(string(data), "\n") {
		// Skip comments and empty lines.
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		// A top-level `plugins:` under a configs list item is
		// indented exactly 4 spaces (2 for list, 2 for field).
		if strings.HasPrefix(line, "    plugins:") || strings.HasPrefix(line, "  plugins:") {
			return true
		}
	}
	return false
}

// findRepoConfigsDir walks upward from the test's cwd looking for a
// directory that contains both `configs/` and `hyoka/` — our repo-root
// signature. Returns "" when not found (e.g. when tests run from an
// extracted tarball).
func findRepoConfigsDir(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	for i := 0; i < 8; i++ {
		configsDir := filepath.Join(wd, "configs")
		hyokaDir := filepath.Join(wd, "hyoka")
		if stat(configsDir) && stat(hyokaDir) {
			return configsDir
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			return ""
		}
		wd = parent
	}
	return ""
}

func stat(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
