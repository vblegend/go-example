package commands

import (
	"backend/common/global"
	"backend/core/console"
	"backend/core/sdk/pkg"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "start",
		Short:   fmt.Sprintf("start %s service", global.AppFileName),
		Example: fmt.Sprintf("%s service start", global.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			pid := 0
			if pkg.IsRuning(&pid) {
				fmt.Printf("服务[%s]...\n", console.Red("已运行"))
				os.Exit(0)
			}
			exe := exec.Command(global.AppFileName, "server")
			exe.Dir = pkg.AssemblyDir()
			exe.Start()
			fmt.Printf("服务[%s]...\n", console.Green("已启动"))
			os.Exit(0)
		},
	}
)
