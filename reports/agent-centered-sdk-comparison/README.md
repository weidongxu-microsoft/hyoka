# Agent-centered SDK comparison

This evaluation tests whether Azure MCP and skills improve generated Azure SDK
projects over the baseline configuration. The run is still in progress, so all
findings are preliminary.

## Method

Each prompt runs with the baseline and applicable tool and skill configurations.
Prompt Checks come from the prompt file, Language Checks come from
`criteria/language/<language>.yaml`, and Program Checks run the language build.
Each applicable check passes or fails. Results are reported as passed checks
over applicable checks for each category and configuration; the categories are
not combined.

The evaluation runs one language at a time with a human review gate. Reviewer,
generation, and infrastructure failures are preserved and reported separately
from code-quality findings.

## Current .NET checkpoint

- 25 of 60 reports and 8 of 20 complete comparison sets
- Two reviewer HTTP 400 failures and one generation without a report

| Configuration | Prompt Checks | Language Checks | Program Checks |
|---|---:|---:|---:|
| Baseline | 46/53 | 22/24 | 8/8 |
| Azure MCP and skills | 40/46* | 23/24 | 7/8 |
| Azure MCP, skills, and .NET SDK skills | 46/53 | 23/24 | 8/8 |

\* One Prompt Check review failed with HTTP 400, so its denominator is not
comparable.

## Preliminary finding

The baseline remains competitive. Against the .NET SDK skills configuration,
five of eight Prompt Check results tie, the baseline leads two, and the SDK
skills configuration leads one. Language Checks improve by only one pass, while
Program Checks do not improve. The results so far do not show a consistent or
material benefit from adding MCP or skills.

Only eight .NET comparison sets are complete. Model variability, reviewer
failures, and the remaining languages may change the final conclusion.
