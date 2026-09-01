# Foundry baseline evaluation

This branch captures the Foundry-tagged baseline evaluation, including raw
evaluation data and summarized results.

The evaluation runs one language and case first, with a review gate before the
remaining cases and languages. The baseline configuration provides no
Azure-specific tools or skills to the coding agent.

## Progress

| Language | Configuration | Status | Reports | Notes |
|---|---|---|---:|---|
| Python | Baseline | In progress | 6/10 | Smoke plus 5 remaining cases complete |
| .NET | Baseline | In progress | 8/10 | Language run active |
| Java | Baseline | In progress | 5/10 | Language run active |
| JavaScript/TypeScript | Baseline | In progress | 6/10 | Language run active |

## Prompt checks

Interim: 215/217 passed (99.1%).

## Language checks

Interim: 124/169 passed (73.4%).

## Program checks

Interim: 31/31 passed (100%).

These rates cover the 25 reports preserved at this checkpoint and will change
as the remaining evaluations complete.

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

## Interpretation limits

- Each prompt and configuration combination is a single trial and is subject
  to model variance.
- Prompt checks, language checks, and program checks are reported separately.
- Cross-language scores are not directly comparable because prompt inventories
  and criteria differ.
