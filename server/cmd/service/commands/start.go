package commands

import (
	"fmt"
	"os"
	"os/exec"
	"server/common/assembly"
	"server/sugar/echo"
	"server/sugar/env"

	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "start",
		Short:   fmt.Sprintf("start %s service", assembly.AppFileName),
		Example: fmt.Sprintf("%s service start", assembly.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			pid := 0
			if assembly.IsRuning(&pid) {
				fmt.Printf("服务[%s]...\n", echo.Red("已运行"))
				os.Exit(0)
			}
			exe := exec.Command(assembly.AppFileName, "server")
			exe.Dir = env.AssemblyDir
			exe.Start()
			fmt.Printf("服务[%s]...\n", echo.Green("已启动"))
			os.Exit(0)
		},
	}
)
