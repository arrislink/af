package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "af",
	Short: "AI Factory — 工具链整合 CLI",
	Long: `af 整合 OpenCode/OMO + Supacode + Trae + Antigravity 的 AI 开发工作流。

  af init           初始化共享记忆与规则（.ai/ + context/ + tasks/）`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Version = "0.1.0"

	rootCmd.AddCommand(initCmd)
}

func printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func printSuccess(msg string) {
	fmt.Printf("✅ %s\n", msg)
}
