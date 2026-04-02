# AI Factory — 产品设计

## 愿景

用现有工具组合成高效的 AI 协作开发流水线，不重复造轮子。

---

## ACD — Atomic Context Development

**核心理念：代码即逻辑，文档即记忆，工具即执行单元。**

### 三层结构

```
规划层  → Antigravity（需求拆解 / UI 验收）
执行层  → OpenCode + OMO（写代码 / 测试 / 修复）
审查层  → Trae AI（diff review / 精修）
并行层  → Supacode（多项目同时推进）
```

### 共享记忆系统

每个项目有独立的 `.ai/`：
- `memory_bank.md` — 当前状态、进度、风险
- `decision_log.md` — 重大决策与原因

全局 `context/global.md` 提供跨项目规则。

### 隔离机制

Git worktree 天然支持并行隔离：
- 每个任务一个独立分支
- 不污染主分支
- 合并顺序可控制

---

## 工具定位

| 工具 | 角色 | 使用场景 |
|------|------|---------|
| **OpenCode + OMO** | 执行引擎 | 写代码、补测试、修复失败、复杂重构 |
| **Supacode** | 并行工作台 | 同时操作多个项目/任务 |
| **Trae AI** | 审查与精修 | diff review、小范围修改 |
| **Antigravity** | 规划层 | 需求模糊时拆分任务、UI 验收（可选） |
| **Git worktree** | 隔离层 | 任务隔离、并行不冲突 |

---

## 协议层（AGENTS.md）

所有 AI 工具共享同一套约束：

1. **Read First**：执行前读 `.ai/memory_bank.md`
2. **Context Awareness**：遵守 `context/global.md` 规则
3. **Log or Die**：重大变更写入 `.ai/decision_log.md`
4. **可验证**：所有改动必须运行测试或说明验证方式

---

## 目录结构

```
ai-factory/           # Git 仓库
├── .ai/              # 共享记忆（memory_bank / decision_log）
├── context/          # 全局规则
├── tasks/            # 任务模板
├── AGENTS.md         # AI 协作协议
├── GUIDE.md          # 工作流指南
└── README.md         # 入口
```

---

## 与传统方案的区别

| 项目 | 传统方案 | AI Factory |
|------|---------|-----------|
| 上下文 | 每次会话重新描述 | 共享记忆持久化 |
| 多工具协同 | 各自为战 | AGENTS.md 协议统一 |
| 并行 | 容易冲突 | Git worktree 物理隔离 |
| 知识传承 | LLM 即用即忘 | decision_log 沉淀 |
