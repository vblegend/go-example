package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"server/common/assembly"
	"server/common/config"
	"server/sugar/env"
	"strings"

	"github.com/spf13/cobra"
)

var (
	lastLine int = 9999999
	LogsCmd      = &cobra.Command{
		Use:     "logs",
		Short:   fmt.Sprintf("view %s service logs", assembly.AppFileName),
		Example: fmt.Sprintf("%s service logs -l20", assembly.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			// servicelog := filepath.Join(config.LoggerConfig.Path, fmt.Sprintf("%s.%s", config.LoggerConfig.FileName, config.LoggerConfig.FileSuffix))
			// data, err := os.ReadFile(servicelog)
			// if err != nil {
			// 	fmt.Println(err)
			// 	os.Exit(0)
			// }
			ymal := filepath.Join(env.AssemblyDir, "/config/settings.yml")
			config.Setup(ymal)
			Rpath := filepath.Join(env.AssemblyDir, config.Logger.Path, "*-*-*.log")
			files, err := filepath.Glob(Rpath)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			filename := filepath.Join(env.AssemblyDir, "logs.log")
			if len(files) > 0 {
				filename = files[len(files)-1]
			}
			data, err := os.ReadFile(filename)
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
