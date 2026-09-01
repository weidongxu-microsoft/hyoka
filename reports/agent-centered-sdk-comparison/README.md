# Agent-centered SDK comparison

This branch captures the full agent-centered SDK comparison, including raw
evaluation data and summarized results.

The evaluation runs one language at a time with a human review gate between
languages. The four plugin-supported languages use three comparison arms. Go
uses two arms because there is no equivalent Microsoft Azure SDK for Go plugin.

## Progress

| Language | Comparison | Status | Complete groups |
|---|---|---|---:|
| .NET | Three-way | In progress | 6/20 |
| Java | Three-way | Pending | 0 |
| JavaScript/TypeScript | Three-way | Pending | 0 |
| Python | Three-way | Pending | 0 |
| Go | Two-way | Pending | 0 |

## .NET checkpoint

- Reports: 20/60
- Program checks: 19/20 passed
- Retry candidates: one no-files generation and two reviewer HTTP 400 failures

Prompt checks, language checks, and program checks are reported separately.
