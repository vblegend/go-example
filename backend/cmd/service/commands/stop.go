package commands

import (
	"backend/common/assembly"
	"backend/core/echo"
	"backend/core/shell"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	StopCmd = &cobra.Command{
		Use:     "stop",
		Short:   fmt.Sprintf("stop %s service", assembly.AppFileName),
		Example: fmt.Sprintf("%s service stop", assembly.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			pid := 0
			if !assembly.IsRuning(&pid) {
				fmt.Printf("服务[%s]...\n", echo.Red("未运行"))
				os.Exit(0)
			}
			shell.ExeCommand("/bin/bash", "-c", fmt.Sprintf("kill -2 %d ; kill -9 %d", pid, pid))
			fmt.Printf("服务[%s]...\n", echo.Green("已停止"))
			os.Exit(0)
		},
	}
)
