package config

import (
	"backend/cmd/config/commands"
	"backend/common/global"
	"os"

	"github.com/spf13/cobra"
)

var (
	ConfigCmd = &cobra.Command{
		Use:               "config",
		SilenceUsage:      true,
		Short:             "siteweb-manager config",
		Example:           "siteweb-manager config",
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				global.PrintCobraHelp()
				os.Exit(1)
			}
			cmds := cmd.Commands()
			for _, cmd := range cmds {
				if cmd.Use == args[0] {
					return nil
				}
			}
			global.PrintCobraHelp()
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
	ConfigCmd.AddCommand(commands.ResetConfig)
}
