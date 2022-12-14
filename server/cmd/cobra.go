package cmd

import (
	"fmt"
	"os"
	"server/cmd/server"
	"server/cmd/service"
	"server/common/assembly"
	"server/sugar/echo"

	"github.com/spf13/cobra"
)

var showVersion bool = false

var rootCmd = &cobra.Command{
	Use:               assembly.AppFileName,
	Short:             "go-example",
	SilenceUsage:      true,
	Long:              assembly.AppFileName,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	PreRun: func(_ *cobra.Command, _ []string) {
		if showVersion {
			fmt.Printf("%s version %s build %s commit %s\n", echo.Green(assembly.AppFileName), echo.Green(assembly.Version), echo.Green(assembly.BuildTime), echo.Green(assembly.CommitID))
			os.Exit(0)
		}
	},
	Run: func(_ *cobra.Command, _ []string) {
		assembly.PrintCobraHelp()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show the current version of the application")
	rootCmd.AddCommand(server.StartCmd)    // 启动主API服务
	rootCmd.AddCommand(service.ServiceCmd) // 服务控制CIL
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
