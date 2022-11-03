package commands

import (
	"fmt"
	"os"
	"server/common/assembly"
	"server/sugar/echo"

	"github.com/spf13/cobra"
)

var (
	StatusCmd = &cobra.Command{
		Use:     "status",
		Short:   fmt.Sprintf("get %s service status", assembly.AppFileName),
		Example: fmt.Sprintf("%s service status", assembly.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			pid := 0
			if assembly.IsRuning(&pid) {
				fmt.Printf("服务[%s]...\n", echo.Green("已运行"))
				os.Exit(0)
			} else {
				fmt.Printf("服务[%s]...\n", echo.Red("未运行"))
				os.Exit(1)
			}
		},
	}
)
