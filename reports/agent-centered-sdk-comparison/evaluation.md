# Agent-centered SDK evaluation

## Objective

This evaluation tests whether Azure MCP, general Azure skills, or
language-specific Azure SDK skills improve generated projects over the baseline
configuration. The baseline is the primary reference because it uses the same
model and prompts without the added tools or skills.

The run is incomplete. Current findings are preliminary and cover only eight
complete .NET comparison sets.

## Configurations

- **Baseline:** No Azure MCP or generator skills.
- **Azure MCP and skills:** Azure MCP plus the general Azure skills plugin.
- **Azure SDK skills:** Azure MCP, general Azure skills, and the applicable
  language-specific SDK skills plugin.

.NET, Java, JavaScript/TypeScript, and Python use all three configurations. Go
uses the baseline and Azure MCP configurations because no equivalent Go SDK
skills plugin is available.

## Methodology

Each prompt runs once with every applicable configuration. The prompt, generator
model, system instructions, reviewer model, and limits remain the same; only the
available tools and skills differ.

Before each language run:

1. Validate the checked-in prompts, criteria, and configurations.
2. Confirm the expected prompt and report counts with a dry run.
3. Run one prompt across all configurations as a smoke check.
4. Run the full language set and preserve every result, including failures.
5. Reconcile generated reports against the expected comparison sets.

### Prompt Checks

Prompt Checks come from each prompt's `## Evaluation Criteria` section. The
reviewer evaluates every listed requirement independently against the generated
project. Examples include required packages, client construction, requested
operations, authentication, and error handling.

### Language Checks

Language Checks come from `criteria/language/<language>.yaml` and apply through
the file's language condition. They evaluate language-specific Azure SDK
practices such as authentication, exception handling, and asynchronous usage.

Program Checks remain separate. They run the configured build command in the
generated project; exit code 0 passes and any other result fails.

## Grading and scoring

Each applicable criterion produces a pass or fail point. The current Prompt and
Language Checks use equal point weights. Results are reported as:

```text
pass rate = passed checks / applicable checks
```

Prompt, Language, and Program Checks are not combined into one headline score.
Skipped checks are excluded. Generation, reviewer, and infrastructure failures
are reported separately because they are not evidence of code quality.

## Current data

Raw smoke and full-run data are stored under:

- `reports/agent-centered-sdk-comparison/smoke/`
- `reports/agent-centered-sdk-comparison/full/`

The raw reports include generated files, grader points and reasoning, tool
calls, session events, timelines, configuration details, and environment data.

### .NET checkpoint

- 25 of 60 reports
- 8 of 20 complete comparison sets
- Two reviewer HTTP 400 failures
- One generation without a report

| Configuration | Prompt Checks | Language Checks | Program Checks |
|---|---:|---:|---:|
| Baseline | 46/53 | 22/24 | 8/8 |
| Azure MCP and skills | 40/46* | 23/24 | 7/8 |
| Azure SDK skills | 46/53 | 23/24 | 8/8 |

\* One Prompt Check review failed with HTTP 400, so this result is not directly
comparable.

## Preliminary findings

The baseline remains competitive. Compared with the Azure SDK skills
configuration, five of eight Prompt Check results tie, the baseline leads two,
and the SDK skills configuration leads one. Each skills configuration improves
Language Checks by only one pass, and neither improves Program Checks.

The available results do not show a consistent or material improvement from
adding MCP or skills.

## Assumptions and limitations

- Only eight .NET comparison sets are complete; other languages may differ.
- Each prompt runs once per configuration, so model variability is not measured.
- LLM-reviewed checks are nondeterministic and depend on one reviewer model.
- Build results depend on installed toolchains and dependency availability.
- Reviewer and generation failures reduce comparability and require separate
  reporting.
- Final findings require all expected reports and complete comparison sets.
