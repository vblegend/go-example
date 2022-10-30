package commands

import (
	"backend/common/assembly"
	"backend/core/echo"
	"backend/core/env"
	"backend/core/shell"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	RestartCmd = &cobra.Command{
		Use:     "restart",
		Short:   fmt.Sprintf("restart %s service", assembly.AppFileName),
		Example: fmt.Sprintf("%s service restart", assembly.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			pid := 0
			if assembly.IsRuning(&pid) {
				bash := fmt.Sprintf("kill -2 %d\nkill -9 %d", pid, pid)
				shell.ExeCommand("/bin/bash", "-c", bash)
				fmt.Printf("服务[%s]...\n", echo.Green("已停止"))
			}
			exe := exec.Command(assembly.AppFileName, "server")
			exe.Dir = env.AssemblyDir
			exe.Start()
			fmt.Printf("服务[%s]...\n", echo.Green("已启动"))
			os.Exit(0)
		},
	}
)
