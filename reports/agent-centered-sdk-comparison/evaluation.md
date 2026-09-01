# Agent-centered SDK evaluation

## Objective

This evaluation tests whether Azure client libraries are agent-friendly: can an
AI coding agent produce correct, idiomatic, and buildable Azure SDK projects
from representative developer prompts?

Baseline results are the primary evidence. Configurations with MCP and skills
show whether external guidance changes the outcome.

## Configurations

| Configuration | Coding-agent assistance |
|---|---|
| Baseline | No Azure-specific tools or skills |
| Azure MCP and skills | Azure MCP and general Azure skills |
| Azure SDK skills | Azure MCP, general Azure skills, and language-specific SDK skills |

.NET, Java, JavaScript/TypeScript, and Python use all three configurations. Go
uses two because no equivalent Go SDK skills plugin is available.

## Methodology

The evaluation measures this through three independent dimensions:

- **Prompt Checks:** Measure task completion and API usability using each
  prompt's `## Evaluation Criteria`.
- **Language Checks:** Measure idiomatic SDK integration using
  `criteria/language/<language>.yaml`.
- **Program Checks:** Restore and compile the project to verify package choices,
  API signatures, dependencies, and project structure.

The prompts represent common developer tasks, including authentication, client
construction, CRUD operations, pagination, error handling, retries, polling,
batch operations, and event processing. They test the complete developer
experience rather than isolated API calls.

Results are reported separately for each dimension. This prevents strong
performance in one area from hiding failures in another. Coverage across
services, tasks, and languages indicates whether behavior is broadly consistent.

## Grading and scoring

Each applicable criterion produces an equally weighted pass or fail point. The
pass rate is passed checks divided by applicable checks. Check categories remain
separate, skipped checks are excluded, and infrastructure failures are reported
separately.

## Results

Evaluation data, findings, and configuration comparisons will be added after the
run is complete and the expected reports are reconciled.

## Assumptions and limitations

- Each prompt runs once per configuration; model variability is not measured.
- LLM-reviewed checks are nondeterministic and use one reviewer model.
- Build results depend on installed toolchains and dependency availability.
- Many prompts specify the expected package, client, or SDK pattern. They provide
  stronger evidence of correct implementation than independent API discovery.
- The prompts cover common tasks, not every library or usage pattern.
- The prompts focus on local code generation and prohibit live Azure operations,
  limiting opportunities for MCP to help.
- Configuration comparisons do not measure general tool effectiveness or prove
  that the underlying library is agent-friendly.
