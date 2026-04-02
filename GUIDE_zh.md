# AI Factory 工作流指南

## 核心理念

AI Factory 不重复造轮子。核心是 `.ai/` 共享记忆 + `AGENTS.md` 协作协议，让所有 AI 工具在同一套上下文和规则下工作。

[English](./GUIDE.md)

---

## 工具职责

| 工具 | 做什么 | 不做什么 |
|------|--------|---------|
| **Supacode** | 多项目并行工作台，直接加载任意 git 目录 | 不生成任务、不调度 |
| **OpenCode + OMO** | 执行引擎：读上下文、写代码、补测试、修失败 | 不做 UI、不做规划 |
| **Trae AI** | 实时代码补全、diff review、精准小改 | 不做大范围重构 |
| **Antigravity** | 需求拆解、任务规划、Browser 截图验收 | 不做日常开发 |

---

## 初始化

```bash
./af init
```

创建：
- `.ai/memory_bank.md` — 共享记忆
- `.ai/decision_log.md` — 决策日志
- `context/global.md` — 全局规则
- `tasks/template.md` — 任务模板

---

## 日常开发流程

### 单项目开发

```
1. Supacode 打开项目目录
2. opencode -p "
   基于以下内容完成任务：
   $(cat .ai/memory_bank.md)
   $(cat context/global.md)
   [任务描述]
   "
3. go test / npm test 验证
4. Trae review diff
5. 重大决策写入 .ai/decision_log.md
```

### 多项目并行（核心场景）

```
Supacode 窗格 1 → api-v6 项目
  opencode -p "实现积分扣减接口..."

Supacode 窗格 2 → admin-portal 项目
  opencode -p "接入 api-v6 扣减接口..."

每个任务完成后：
  git commit → 合并
  跨项目影响写入 .ai/decision_log.md
```

### 复杂任务（启用 OMO）

```bash
opencode -p "使用多 agent 分工：
  - agent1: 写接口实现
  - agent2: 写单元测试
  - agent3: 做 code review
  基于 context/global.md 规则执行"
```

### 失败修复循环

```bash
go test ./... 2>&1 | tee test.log
opencode -p "测试失败，请修复：
  $(cat test.log)"
# 重复直到通过
```

---

## 共享记忆维护

| 时机 | 写入位置 | 内容 |
|------|---------|------|
| 任务开始 | `.ai/memory_bank.md` | 当前进度、风险 |
| 任务完成 | `.ai/memory_bank.md` | 完成状态、下一步 |
| 重大决策 | `.ai/decision_log.md` | 决策原因、权衡 |
| 跨项目影响 | 所有涉及项目的 `.ai/decision_log.md` | 接口变更、依赖关系 |

---

## 项目结构约定

每个项目（api-v6、admin-portal 等）都可以有自己的 `.ai/`：

```
~/Projects/api-v6/
├── .ai/
│   ├── memory_bank.md
│   └── decision_log.md
├── context/         # 可选
└── src/

~/Projects/admin-portal/
├── .ai/
│   ├── memory_bank.md
│   └── decision_log.md
└── src/
```

AI 执行时会优先读项目自身的 `.ai/`，再读全局 `context/`。

---

## AI 协作协议

所有 AI 工具在执行任务前必须：

1. **Read First**：读取 `.ai/memory_bank.md`
2. **Context Awareness**：遵守 `context/global.md` 中的规则
3. **Log or Die**：重构、Bug 修复、依赖变更写入 `.ai/decision_log.md`
4. **可验证**：所有改动必须运行测试或明确说明验证方式
5. **Atomic Commit**：每次 commit 说明对上下文的影响

---

## 任务模板

```markdown
# 任务：[标题]

## 背景
[为什么做这个任务]

## 目标
[完成后状态]

## 具体要求
1. [要求 1]
2. [要求 2]

## 边界条件
- [边界情况]

## 验收标准
- [ ] [测试/检查项]
```

---

## 模型选型建议

| 场景 | 推荐模型 | 成本 |
|------|---------|------|
| 日常执行 | `opencode/gpt-5-nano` | 低 |
| 复杂重构（OMO 多 agent） | `anthropic/claude-sonnet-4-6` | 中 |
| 规划（Antigravity） | Gemini 3 Pro | 免费额度 |
| Trae 补全 | 内置模型 | 包月 |

**费用控制**：日常任务用低成本模型；只在真正需要 OMO 多 agent 时才用高级模型。

---

## 常见问题

**Q: 多项目并行时如何保证不混乱？**
A: 每个项目有独立的 `.ai/memory_bank.md`；Supacode 窗格之间完全隔离；跨项目影响手动写入 `decision_log.md`。

**Q: Antigravity 什么时候用？**
A: 需求模糊时、跨多个项目需要协调时、需要 UI 截图验收时。不做日常代码开发。

**Q: 测试一直失败怎么办？**
A: 用 `opencode -p "修复：$(cat test.log)"` 循环修复，每次修复后重跑测试，直到通过。

**Q: 多人如何共享 `.ai/`？**
A: 放在网络盘或通过 Git 管理（加到 `.gitignore` 反向排除）。
