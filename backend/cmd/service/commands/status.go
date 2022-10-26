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
	StatusCmd = &cobra.Command{
		Use:     "status",
		Short:   fmt.Sprintf("get %s service status", global.AppFileName),
		Example: fmt.Sprintf("%s service status", global.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			pid := 0
			if pkg.IsRuning(&pid) {
				fmt.Printf("服务[%s]...\n", console.Green("已运行"))
				os.Exit(0)
			} else {
				fmt.Printf("服务[%s]...\n", console.Red("未运行"))
				os.Exit(1)
			}
		},
	}
)
