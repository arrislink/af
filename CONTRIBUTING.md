# Contributing

## Project Goals

AI Factory is a toolchain orchestration based on ACD (Atomic Context Development). Core principles:

- **No wheel reinventing**: Use existing tools, don't build new systems
- **Protocol over code**: `.ai/` shared memory + `AGENTS.md` is the core — keep it minimal
- **Stay lean**: Only do what truly needs to be done

## How to Contribute

### 1. Report Issues

- Open an Issue describing the problem or suggestion
- Check for duplicates first

### 2. Submit Changes

```bash
git clone git@github.com:arrislink/af.git
git checkout -b feat/your-feature
# Make changes
git commit -m "feat: add something"
git push -u origin HEAD
# Open PR
```

### 3. PR Standards

- Title concise, follow [Conventional Commits](https://www.conventionalcommits.org/)
- Describe what and why
- Ensure `go build ./...` and `go test ./...` pass

## Directory Overview

```
.ai/          → Shared memory (memory_bank / decision_log)
context/      → Global rules
tasks/        → Task templates
AGENTS.md     → AI collaboration protocol (core, don't modify casually)
GUIDE.md      → Workflow guide
PRD.md        → Product design
```

## Code Standards

- Go code follows `go fmt` + `go vet`
- Markdown files stay concise, no duplicate content
