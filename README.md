# AI Factory

[![CI](https://github.com/arrislink/af/actions/workflows/CI.yml/badge.svg)](https://github.com/arrislink/af/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

An AI-powered development workflow that orchestrates OpenCode + OMO, Supacode, Trae, and Antigravity.

No wheel-reinventing. Core is `.ai/` shared memory + `AGENTS.md` collaboration protocol — making every AI tool work from the same context and rules.

[中文](./README_zh.md)

---

## Tools

| Tool | Role |
|------|------|
| **OpenCode + OMO** | Execution engine: write code, tests, fix failures |
| **Supacode** | Parallel workspace: multiple projects simultaneously |
| **Trae AI** | Code review and polish |
| **Antigravity** | Planning and UI validation (optional) |

## Quick Start

```bash
git clone git@github.com:arrislink/af.git && cd af
./af init

# Open Supacode, load your project directories
opencode -p "Implement the points deduction API based on .ai/ and context/"
```

## Workflow

```
Supacode pane 1 → api-v6 project
  opencode -p "Implement deduction API"

Supacode pane 2 → admin-portal project
  opencode -p "Integrate deduction API"

After each task:
  → git commit → merge
  → Write significant decisions to .ai/decision_log.md
```

See [GUIDE.md](./GUIDE.md) for full workflow.

## Project Structure

```
ai-factory/
├── .ai/              # Shared memory
│   ├── memory_bank.md
│   └── decision_log.md
├── context/          # Global rules
│   └── global.md
├── tasks/            # Task templates
│   └── template.md
├── AGENTS.md         # AI collaboration protocol
├── GUIDE.md          # Full workflow guide
└── PRD.md            # Product design
```

## AI Collaboration Protocol

All AI tools must follow (see `AGENTS.md`):

1. **Read First**: Read `.ai/memory_bank.md` before any change
2. **Context Awareness**: Follow rules in `context/global.md`
3. **Log or Die**: Write significant changes to `.ai/decision_log.md`
4. **Verifiable**: All changes must run tests or explain verification

## Install

### Option 1: Download Binary (Fastest)

```bash
# macOS amd64
curl -L https://github.com/arrislink/af/releases/latest/download/af-darwin-amd64 -o af
chmod +x af
./af init

# macOS ARM64
curl -L https://github.com/arrislink/af/releases/latest/download/af-darwin-arm64 -o af
chmod +x af
./af init

# Linux amd64
curl -L https://github.com/arrislink/af/releases/latest/download/af-linux-amd64 -o af
chmod +x af
./af init
```

### Option 2: Go Install

```bash
go install github.com/arrislink/af/cmd/af@latest
./af init
```

### Option 3: Build from Source

```bash
git clone git@github.com:arrislink/af.git && cd af
make install
./af init
```

## Develop

```bash
make all      # fmt + vet + test + build
make test     # run tests
make vet      # go vet
```

## License

MIT
