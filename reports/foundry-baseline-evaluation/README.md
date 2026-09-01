# Foundry baseline evaluation

This branch captures the Foundry-tagged baseline evaluation, including raw
evaluation data and summarized results.

The evaluation runs one language and case first, with a review gate before the
remaining cases and languages. The baseline configuration provides no
Azure-specific tools or skills to the coding agent.

## Progress

| Language | Configuration | Status | Reports | Notes |
|---|---|---|---:|---|
| Python | Baseline | Smoke test in progress | 0/10 | Basic agent lifecycle runs first |
| .NET | Baseline | Pending | 0/10 | Wait for smoke-test review |
| Java | Baseline | Pending | 0/10 | Wait for smoke-test review |
| JavaScript/TypeScript | Baseline | Pending | 0/10 | Wait for smoke-test review |

## Prompt checks

Pending.

## Language checks

Pending.

## Program checks

Pending.

## Run health

The initial launches failed before evaluation because the log directory,
configuration names, or Foundry prompt checkout were incorrect. They produced
no reports or Copilot sessions and are not evaluation attempts. The isolated
run worktree now uses the Foundry comparison commit, checksum-verified Copilot
CLI `v1.0.81-11`, and the matching baseline configuration.

## Interpretation limits

- Each prompt and configuration combination is a single trial and is subject
  to model variance.
- Prompt checks, language checks, and program checks are reported separately.
- Cross-language scores are not directly comparable because prompt inventories
  and criteria differ.
