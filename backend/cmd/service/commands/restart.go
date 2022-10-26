package commands

import (
	"backend/common/global"
	"backend/core/console"
	"backend/core/sdk/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	RestartCmd = &cobra.Command{
		Use:     "restart",
		Short:   fmt.Sprintf("restart %s service", global.AppFileName),
		Example: fmt.Sprintf("%s service restart", global.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			pid := 0
			if pkg.IsRuning(&pid) {
				bash := fmt.Sprintf("kill -2 %d\nkill -9 %d", pid, pid)
				pkg.ExeCommand("/bin/bash", "-c", bash)
				fmt.Printf("服务[%s]...\n", console.Green("已停止"))
			}
			bash := fmt.Sprintf("cd %s \nnohup ./%s server", pkg.AssemblyDir(), global.AppFileName)
			pkg.ExeCommand("/bin/bash", "-c", bash)
			fmt.Printf("服务[%s]...\n", console.Green("已启动"))
			os.Exit(0)
		},
	}
)
