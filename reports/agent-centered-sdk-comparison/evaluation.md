# Agent-centered SDK evaluation

## Objective

This evaluation tests whether Azure client libraries are agent-friendly: can an
AI coding agent produce correct, idiomatic, and buildable Azure SDK projects
from representative developer prompts?

The baseline configuration is the primary evidence because it measures how well
the agent can use the libraries without Azure-specific tools or skills. The
other configurations show whether additional guidance changes the outcome, but
that comparison is secondary to the agent-friendliness assessment.

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

Together, the checks measure three aspects of agent-friendliness:

- **Task completion:** The project satisfies the prompt-specific requirements.
- **SDK usability:** The project follows language and Azure SDK conventions.
- **Buildability:** The generated project restores and compiles.

## Grading and scoring

Each applicable criterion produces a pass or fail point. The current Prompt and
Language Checks use equal point weights. Results are reported as:

```text
pass rate = passed checks / applicable checks
```

Prompt, Language, and Program Checks are not combined into one headline score.
Skipped checks are excluded. Generation, reviewer, and infrastructure failures
are reported separately because they are not evidence of code quality.

## Results

Evaluation data, findings, and configuration comparisons will be added after the
run is complete and the expected reports are reconciled.

## Assumptions and limitations

- Each prompt runs once per configuration, so model variability is not measured.
- LLM-reviewed checks are nondeterministic and depend on one reviewer model.
- Build results depend on installed toolchains and dependency availability.
- Reviewer and generation failures reduce comparability and require separate
  reporting.
- The selected prompts represent common Azure SDK tasks but do not cover every
  library or usage pattern.
- Tools and skills may improve agent performance, but they do not by themselves
  demonstrate that the underlying library is agent-friendly.
- Final findings require all expected reports and complete comparison sets.
