package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/arrislink/af/internal/config"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化共享记忆与规则",
	RunE: func(cmd *cobra.Command, args []string) error {
		dirs := []string{
			config.AIDir(),
			config.ContextDir(),
			fmt.Sprintf("%s/backlog", config.TasksDir()),
			fmt.Sprintf("%s/doing", config.TasksDir()),
			fmt.Sprintf("%s/done", config.TasksDir()),
		}

		for _, d := range dirs {
			if err := os.MkdirAll(d, 0755); err != nil {
				return fmt.Errorf("create %s: %w", d, err)
			}
		}

		globalMd := fmt.Sprintf("%s/global.md", config.ContextDir())
		if _, err := os.Stat(globalMd); os.IsNotExist(err) {
			content := `# 全局上下文

## 代码规范
- 所有代码必须可运行
- 必须包含单元测试
- 不允许伪代码
- 接口必须有错误处理
- 所有数据库操作必须考虑事务

## 禁止事项
- 不得硬编码密码、密钥、IP
- 不得直接操作 main 分支
`
			if err := os.WriteFile(globalMd, []byte(content), 0644); err != nil {
				return err
			}
		}

		memoryBank := fmt.Sprintf("%s/memory_bank.md", config.AIDir())
		if _, err := os.Stat(memoryBank); os.IsNotExist(err) {
			content := `# 共享记忆

## 当前状态
- AI Factory 工具链整合完成

## 当前待办
- 持续将关键决策沉淀到 decision_log.md
`
			if err := os.WriteFile(memoryBank, []byte(content), 0644); err != nil {
				return err
			}
		}

		decisionLog := fmt.Sprintf("%s/decision_log.md", config.AIDir())
		if _, err := os.Stat(decisionLog); os.IsNotExist(err) {
			content := `# 决策日志
`
			if err := os.WriteFile(decisionLog, []byte(content), 0644); err != nil {
				return err
			}
		}

		templateMd := fmt.Sprintf("%s/template.md", config.TasksDir())
		if _, err := os.Stat(templateMd); os.IsNotExist(err) {
			content := `# 任务：[标题]

## 背景
[简述为什么做这个任务]

## 目标
[完成后的状态]

## 具体要求
1. [要求 1]
2. [要求 2]

## 边界条件
- [边界情况]

## 验收标准
- [ ] [测试/检查项]
`
			if err := os.WriteFile(templateMd, []byte(content), 0644); err != nil {
				return err
			}
		}

		printSuccess("初始化完成")
		fmt.Printf("  .ai/      → 共享记忆\n")
		fmt.Printf("  context/  → 全局规则\n")
		fmt.Printf("  tasks/    → 任务模板\n")
		fmt.Printf("\n下一步: 打开 Supacode，直接开始\n")

		return nil
	},
}
