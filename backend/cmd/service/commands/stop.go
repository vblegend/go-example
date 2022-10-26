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
	StopCmd = &cobra.Command{
		Use:     "stop",
		Short:   fmt.Sprintf("stop %s service", global.AppFileName),
		Example: fmt.Sprintf("%s service stop", global.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			pid := 0
			if !pkg.IsRuning(&pid) {
				fmt.Printf("服务[%s]...\n", console.Red("未运行"))
				os.Exit(0)
			}
			pkg.ExeCommand("/bin/bash", "-c", fmt.Sprintf("kill -2 %d ; kill -9 %d", pid, pid))
			pkg.RemovePidFile()
			fmt.Printf("服务[%s]...\n", console.Green("已停止"))
			os.Exit(0)
		},
	}
)
