# AI Collaboration Protocol

All AI tools must follow these rules when executing any task.

1. **Read First**: Read `.ai/memory_bank.md` before any code change
2. **Context Awareness**: Follow rules defined in `context/global.md`
3. **Log or Die**: Refactoring, bug fixes, dependency changes must be written to `.ai/decision_log.md` (with reason and tradeoffs)
4. **Verifiable**: All changes must run tests or explicitly explain verification method
5. **Atomic Commit**: Each commit must describe the impact on project context
