package cmd

import (
	"backend/cmd/api"
	"backend/cmd/config"
	"backend/cmd/service"
	"backend/common/global"
	"backend/core/sdk/console"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var showVersion bool = false

var rootCmd = &cobra.Command{
	Use:               "siteweb-manager",
	Short:             "服务器管理台",
	SilenceUsage:      true,
	Long:              `siteweb-manager`,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	PreRun: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf("%s version %s build %s commit %s\n", console.Green(global.AppName), console.Green(global.Version), console.Green(global.BuildTime), console.Green(global.CommitID))
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		global.PrintCobraHelp()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show the current version of the application")
	rootCmd.AddCommand(api.StartCmd)       // 启动主API服务
	rootCmd.AddCommand(service.ServiceCmd) // 服务控制CIL
	rootCmd.AddCommand(config.ConfigCmd)   // 配置相关CIL
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
