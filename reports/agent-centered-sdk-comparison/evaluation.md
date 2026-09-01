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

- **Baseline:** No Azure-specific tools or skills are provided to the coding
  agent.
- **Azure MCP and skills:** Azure MCP plus the general Azure skills plugin.
- **Azure SDK skills:** Azure MCP, general Azure skills, and the applicable
  language-specific SDK skills plugin.

.NET, Java, JavaScript/TypeScript, and Python use all three configurations. Go
uses the baseline and Azure MCP configurations because no equivalent Go SDK
skills plugin is available.

## Methodology

An agent-friendly SDK enables a coding agent to translate a normal developer
request into a correct, idiomatic, and buildable project without requiring
Azure-specific assistance.

The evaluation measures this through three independent dimensions:

- **Prompt Checks — task completion and API usability:** Derived from each
  prompt's `## Evaluation Criteria`, these checks determine whether the agent
  selects the appropriate packages, clients, authentication mechanisms, and
  operations. Passing indicates that the SDK's APIs and usage patterns can be
  applied correctly from the developer request.
- **Language Checks — idiomatic SDK integration:** Defined in
  `criteria/language/<language>.yaml`, these checks evaluate language-specific
  Azure SDK conventions, including authentication, asynchronous usage,
  exception handling, and resource management. Passing indicates that the SDK
  can be used correctly within the conventions of its target language.
- **Program Checks — project buildability:** These checks restore and compile
  the generated project. Passing provides objective evidence that the selected
  packages, API signatures, project structure, and dependencies work together.

The baseline configuration is the primary measure of agent-friendliness because
the coding agent receives no Azure-specific tools or skills. The additional
configurations are diagnostic: they show whether external guidance can address
failures, but they do not independently demonstrate that the SDK itself is
agent-friendly.

Results are reported separately for each dimension. This prevents strong
performance in one area, such as compilation, from hiding failures in task
completion or idiomatic SDK usage. Coverage across services, tasks, and
languages indicates whether observed behavior is broadly consistent rather than
limited to one API or scenario.

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
