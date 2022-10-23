package commands

import (
	"backend/core/sdk/console"
	"backend/core/sdk/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "start",
		Short:   "start siteweb-manager service",
		Example: "siteweb-manager service start",
		Run: func(cmd *cobra.Command, args []string) {
			pid := 0
			if pkg.IsRuning(&pid) {
				fmt.Printf("服务[%s]...\n", console.Red("已运行"))
				os.Exit(0)
			}
			bash := fmt.Sprintf("cd %s ; nohup ./siteweb-manager server > logs.log 2>&1 &", pkg.AssemblyDir())
			pkg.ExeCommand("/bin/bash", "-c", bash)
			fmt.Printf("服务[%s]...\n", console.Green("已启动"))
			os.Exit(0)
		},
	}
)
