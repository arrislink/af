# AI Factory Workflow Guide

## Philosophy

AI Factory doesn't reinvent wheels. Core is `.ai/` shared memory + `AGENTS.md` collaboration protocol — making OpenCode, Trae, and Antigravity all work from the same context and rules.

---

## Tools

| Tool | Does | Doesn't Do |
|------|------|------------|
| **Supacode** | Parallel workspace, load any git directory | Task generation, scheduling |
| **OpenCode + OMO** | Execution engine: read context, write code, tests, fix failures | UI, planning |
| **Trae AI** | Real-time completion, diff review, precise small changes | Large-scale refactoring |
| **Antigravity** | Requirement breakdown, task planning, Browser screenshot validation | Daily development |

---

## Init

```bash
./af init
```

Creates:
- `.ai/memory_bank.md` — shared memory
- `.ai/decision_log.md` — decision log
- `context/global.md` — global rules
- `tasks/template.md` — task template

---

## Daily Workflow

### Single Project

```
1. Open Supacode, load project directory
2. opencode -p "
   Based on the following context:
   $(cat .ai/memory_bank.md)
   $(cat context/global.md)
   [task description]
   "
3. go test / npm test to verify
4. Trae review diff
5. Write significant decisions to .ai/decision_log.md
```

### Multi-Project Parallel (Core Scenario)

```
Supacode pane 1 → api-v6 project
  opencode -p "Implement points deduction API..."

Supacode pane 2 → admin-portal project
  opencode -p "Integrate api-v6 deduction API..."

After each task:
  git commit → merge
  Cross-project impacts → .ai/decision_log.md
```

### Complex Task (OMO Multi-Agent)

```bash
opencode -p "Use multi-agent:
  - agent1: write interface
  - agent2: write tests
  - agent3: code review
  Based on context/global.md rules"
```

### Fix-Fail Loop

```bash
go test ./... 2>&1 | tee test.log
opencode -p "Test failed, fix:
  $(cat test.log)"
# Repeat until pass
```

---

## Shared Memory Maintenance

| When | Where | What |
|------|-------|------|
| Task start | `.ai/memory_bank.md` | Progress, risks |
| Task done | `.ai/memory_bank.md` | Status, next steps |
| Significant decision | `.ai/decision_log.md` | Reason, tradeoffs |
| Cross-project impact | All involved projects' `.ai/decision_log.md` | Interface changes, dependencies |

---

## Project Structure Convention

Each project (api-v6, admin-portal, etc.) can have its own `.ai/`:

```
~/Projects/api-v6/
├── .ai/
│   ├── memory_bank.md
│   └── decision_log.md
├── context/         # optional
└── src/

~/Projects/admin-portal/
├── .ai/
│   ├── memory_bank.md
│   └── decision_log.md
└── src/
```

AI execution reads project-level `.ai/` first, then global `context/`.

---

## AI Collaboration Protocol

All AI tools must before executing any task:

1. **Read First**: Read `.ai/memory_bank.md`
2. **Context Awareness**: Follow rules in `context/global.md`
3. **Log or Die**: Write refactoring, bug fixes, dependency changes to `.ai/decision_log.md`
4. **Verifiable**: All changes must run tests or explain verification
5. **Atomic Commit**: Each commit describes impact on context

---

## Task Template

```markdown
# Task: [Title]

## Background
[Why this task]

## Goal
[State after completion]

## Requirements
1. [Requirement 1]
2. [Requirement 2]

## Edge Cases
- [Boundary conditions]

## Acceptance Criteria
- [ ] [Test/check item]
```

---

## Model Selection

| Scenario | Model | Cost |
|----------|-------|------|
| Daily execution | `opencode/gpt-5-nano` | Low |
| Complex refactoring (OMO) | `anthropic/claude-sonnet-4-6` | Medium |
| Planning (Antigravity) | Gemini 3 Pro | Free quota |
| Trae completion | Built-in | Subscription |

**Cost control**: Use low-cost models for daily tasks; only use advanced models when OMO multi-agent is truly needed.

---

## FAQ

**Q: How to avoid chaos with multi-project parallel?**
A: Each project has independent `.ai/memory_bank.md`; Supacode panes are fully isolated; cross-project impacts manually written to `decision_log.md`.

**Q: When to use Antigravity?**
A: When requirements are vague, multiple projects need coordination, or UI screenshot validation is needed. Not for daily coding.

**Q: Test keeps failing?**
A: Use `opencode -p "Fix: $(cat test.log)"` in a loop, re-running tests after each fix until passing.

**Q: How to share `.ai/` across team?**
A: Place on network drive or manage via Git (add to `.gitignore` in reverse).
