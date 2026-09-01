# Foundry baseline evaluation

This branch captures the Foundry-tagged baseline evaluation, including raw
evaluation data and summarized results.

The evaluation runs one language and case first, with a review gate before the
remaining cases and languages. The baseline configuration provides no
Azure-specific tools or skills to the coding agent.

## Progress

| Language | Configuration | Status | Reports | Notes |
|---|---|---|---:|---|
| Python | Baseline | Complete | 10/10 | 1 passed, 8 grader-failed, 1 review error |
| .NET | Baseline | Complete | 10/10 | 6 passed, 4 grader-failed evaluations |
| Java | Baseline | Complete | 10/10 | 10 grader-failed evaluations |
| JavaScript/TypeScript | Baseline | Complete | 10/10 | 2 controlled retries selected |

## Prompt checks

Final selected: 344/348 passed (98.9%).

## Language checks

Final selected: 209/290 passed (72.1%).

## Program checks

Final selected: 50/50 passed (100%).

These rates cover the final 40-report dataset. The original action-limited
JavaScript/TypeScript attempts remain preserved but are replaced by their
controlled reruns in these totals.

## Run health

The Python smoke evaluation completed generation and review, produced nine
files, and encountered no timeouts or malformed tool invocations. The process
exited successfully; the evaluation failed only the Proper Exception Handling
grader.

Initial harness launches failed before evaluation because the log directory,
configuration names, or Foundry prompt checkout were incorrect. They produced
no reports or Copilot sessions and are not evaluation attempts. Language runs
use isolated worktrees at the Foundry comparison commit and checksum-verified
Copilot CLI `v1.0.81-11`.

The .NET stream completed without generation, session, review, MCP, action-limit,
or turn-limit failures. A debug log entry showed discovery of an `azure-python`
plugin definition, but generation sessions were created with zero configured
skills and zero configured MCP servers; the plugin was not exposed to the
coding agent.

The Python stream completed all 10 generations with final responses and
generated files. The `evaluation-run` case generated 60 files, causing all six
review requests to exceed the model token limit. Its Program Check passed, but
its Prompt and Language Checks are unavailable; the preserved result is the
only current infrastructure retry candidate.

The JavaScript/TypeScript stream preserved 10/10 primary reports. `file-search`
and `evaluation-run` reached 51/50 session actions and returned partial results.
Both were rerun once with a 100-action limit, completed in 14 and 20 actions,
and passed both Program Checks. The healthy reruns replace both primary attempts
in the final dataset, irrespective of score direction. Transliteration issued
one malformed PowerShell command and corrected it on the next action.

The Java stream preserved 10/10 reports with all generations, reviews, generated
file sets, tool calls, and Program Checks completing successfully. Its failures
are grader outcomes rather than infrastructure failures.

## Interpretation limits

- Each prompt and configuration combination is a single trial and is subject
  to model variance.
- Prompt checks, language checks, and program checks are reported separately.
- Cross-language scores are not directly comparable because prompt inventories
  and criteria differ.
