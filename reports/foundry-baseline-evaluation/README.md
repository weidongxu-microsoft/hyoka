# Foundry baseline evaluation

This branch captures the Foundry-tagged baseline evaluation, including raw
evaluation data and summarized results.

The evaluation runs one language and case first, with a review gate before the
remaining cases and languages. The baseline configuration provides no
Azure-specific tools or skills to the coding agent.

## Progress

| Language | Configuration | Status | Reports | Notes |
|---|---|---|---:|---|
| Python | Baseline | In progress | 1/10 | Smoke case completed |
| .NET | Baseline | In progress | 0/10 | Language run started |
| Java | Baseline | In progress | 0/10 | Language run started |
| JavaScript/TypeScript | Baseline | In progress | 0/10 | Language run started |

## Prompt checks

Python smoke: 15/16 reviewer checks passed. Proper Exception Handling failed.

## Language checks

Python smoke: included in the 15/16 reviewer result.

## Program checks

Python smoke: passed `python -m compileall -q .`.

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
