package service

import (
	"fmt"
	"os"
	"runtime"
	"server/cmd/service/commands"
	"server/common/assembly"
	"server/sugar/log"

	"github.com/spf13/cobra"
)

var (
	ServiceCmd = &cobra.Command{
		Use:               "service",
		SilenceUsage:      true,
		Short:             fmt.Sprintf("%s control service", assembly.AppFileName),
		Example:           fmt.Sprintf("%s service start/stop/restart/install/uninstall/status", assembly.AppFileName),
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		Args: func(cmd *cobra.Command, args []string) error {
			if runtime.GOOS != "linux" {
				log.Error("error: The service command supports only linux.")
				os.Exit(1)
			}
			if len(args) < 1 {
				assembly.PrintCobraHelp()
				os.Exit(1)
			}
			cmds := cmd.Commands()
			for _, cmd := range cmds {
				if cmd.Use == args[0] {
					return nil
				}
			}
			assembly.PrintCobraHelp()
			os.Exit(1)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Execute() != nil {
				os.Exit(1)
			}
		},
	}
)

func init() {
	ServiceCmd.AddCommand(commands.InstallCmd)
	ServiceCmd.AddCommand(commands.UnInstallCmd)
	ServiceCmd.AddCommand(commands.StartCmd)
	ServiceCmd.AddCommand(commands.StopCmd)
	ServiceCmd.AddCommand(commands.RestartCmd)
	ServiceCmd.AddCommand(commands.StatusCmd)
	ServiceCmd.AddCommand(commands.LogsCmd)
}
