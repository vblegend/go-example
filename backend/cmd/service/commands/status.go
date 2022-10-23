package commands

import (
	"backend/core/sdk/console"
	"backend/core/sdk/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	StatusCmd = &cobra.Command{
		Use:     "status",
		Short:   "get siteweb-manager service status",
		Example: "siteweb-manager service status",
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
