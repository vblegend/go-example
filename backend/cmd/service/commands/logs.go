package commands

import (
	"backend/core/sdk/pkg"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	lastLine int = 9999999
	LogsCmd      = &cobra.Command{
		Use:     "logs",
		Short:   "view siteweb-manager logs",
		Example: "siteweb-manager service logs -l20",
		Run: func(cmd *cobra.Command, args []string) {
			// systemctl status siteweb-manager.service
			servicelog := filepath.Join(pkg.AssemblyDir(), "logs.log")
			data, err := os.ReadFile(servicelog)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			lines := strings.Split(string(data), "\n")
			linelength := len(lines)
			limit := lastLine
			if linelength > limit {
				lines = lines[linelength-limit:]
			}
			for _, line := range lines {
				fmt.Println(line)
			}
			os.Exit(0)
		},
	}
)

func init() {
	LogsCmd.PersistentFlags().IntVarP(&lastLine, "line", "l", 100000, "display the latest log line limit, default(100000)")
}
