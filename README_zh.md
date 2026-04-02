# AI Factory

[![CI](https://github.com/arrislink/af/actions/workflows/CI.yml/badge.svg)](https://github.com/arrislink/af/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

整合 OpenCode + OMO、Supacode、Trae、Antigravity 的 AI 开发工作流。

不重复造轮子。核心是 `.ai/` 共享记忆 + `AGENTS.md` 协作协议，让所有 AI 工具在同一套上下文和规则下工作。

[English](./README.md)

---

## 工具链

| 工具 | 职责 |
|------|------|
| **OpenCode + OMO** | 执行引擎：写代码 / 测试 / 修复 |
| **Supacode** | 并行工作台：多项目同时推进 |
| **Trae AI** | 代码审查与精修 |
| **Antigravity** | 需求规划与 UI 验收（可选） |

## 快速开始

```bash
git clone git@github.com:arrislink/af.git && cd af
./af init

# 打开 Supacode，加载项目目录
opencode -p "基于 .ai/ 和 context/ 实现积分扣减接口"
```

## 工作流

```
Supacode 窗格 1 → api-v6 项目
  opencode -p "实现扣减接口"

Supacode 窗格 2 → admin-portal 项目
  opencode -p "接入扣减接口"

每个任务完成后：
  → git commit → 合并
  → 重大决策写入 .ai/decision_log.md
```

详见 [GUIDE_zh.md](./GUIDE_zh.md)。

## 项目结构

```
ai-factory/
├── .ai/              # 共享记忆
│   ├── memory_bank.md
│   └── decision_log.md
├── context/          # 全局规则
│   └── global.md
├── tasks/            # 任务模板
│   └── template.md
├── AGENTS.md         # AI 协作协议
├── GUIDE_zh.md       # 完整工作流指南
└── PRD.md            # 产品设计
```

## AI 协作协议

所有 AI 工具必须遵守（详见 `AGENTS.md`）：

1. **Read First**：修改前读 `.ai/memory_bank.md`
2. **Context Awareness**：遵守 `context/global.md` 规则
3. **Log or Die**：重大变更写入 `.ai/decision_log.md`
4. **可验证**：所有改动必须运行测试

## 安装

### 方式 1: 下载二进制（最快）

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

### 方式 2: Go Install

```bash
go install github.com/arrislink/af/cmd/af@latest
./af init
```

### 方式 3: 源码构建

```bash
git clone git@github.com:arrislink/af.git && cd af
make install
./af init
```

## 开发

```bash
make all      # fmt + vet + test + build
make test     # 运行测试
make vet      # 代码检查
```

## License

MIT
