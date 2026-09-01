# Foundry baseline evaluation

## Objective

This evaluation tests whether Azure client libraries used by Foundry-related
developer scenarios are agent-friendly: can an AI coding agent produce correct,
idiomatic, and buildable Azure SDK projects from representative prompts without
Azure-specific assistance?

The baseline configuration is the primary evidence because the coding agent
receives no Azure-specific tools or skills.

## Configuration

- **Baseline:** No Azure-specific tools or skills are provided to the coding
  agent.

.NET, Java, JavaScript/TypeScript, and Python each use their language-specific
baseline configuration. Ten Foundry-tagged prompts run once per language, for
40 expected reports.

## Methodology

An agent-friendly SDK enables a coding agent to translate a normal developer
request into a correct, idiomatic, and buildable project without requiring
Azure-specific assistance.

The evaluation measures this through three independent dimensions:

- **Prompt Checks - task completion and API usability:** Derived from each
  prompt's `## Evaluation Criteria`, these checks determine whether the agent
  selects the appropriate packages, clients, authentication mechanisms, and
  operations.
- **Language Checks - idiomatic SDK integration:** Defined in
  `criteria/language/<language>.yaml`, these checks evaluate language-specific
  Azure SDK conventions, including authentication, asynchronous usage,
  exception handling, and resource management.
- **Program Checks - project buildability:** These checks restore and compile
  the generated project. Passing provides objective evidence that the selected
  packages, API signatures, project structure, and dependencies work together.

Before each language run:

1. Validate the checked-in prompts, criteria, and configuration.
2. Confirm the expected prompt and report counts with a dry run.
3. Run one prompt as a smoke check.
4. Run the remaining language set and preserve every result, including
   failures.
5. Reconcile generated reports against the expected language set.

Results are reported separately for each dimension. This prevents strong
performance in one area, such as compilation, from hiding failures in task
completion or idiomatic SDK usage. Coverage across services, tasks, and
languages indicates whether observed behavior is broadly consistent rather
than limited to one API or scenario.

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

Evaluation data and findings will be added after the run is complete and all 40
expected reports are reconciled.

## Assumptions and limitations

- Each prompt runs once, so model variability is not measured.
- LLM-reviewed checks are nondeterministic and depend on one reviewer model.
- Build results depend on installed toolchains and dependency availability.
- Reviewer, generation, MCP, and other infrastructure failures reduce
  comparability and require separate reporting.
- The selected prompts represent Foundry-related Azure SDK tasks but do not
  cover every library or usage pattern.
- Final findings require all expected reports and complete language sets.
