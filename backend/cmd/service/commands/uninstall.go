package commands

import (
	"backend/common/assembly"
	"backend/core/echo"
	"backend/core/env"
	"backend/core/futils"
	"backend/core/shell"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	UnInstallCmd = &cobra.Command{
		Use:     "uninstall",
		Short:   "uninstall app service",
		Example: fmt.Sprintf("%s service uninstall", assembly.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {

			bashFile := ""
			if env.System() == env.Linux_Ubuntu {
				bashFile = fmt.Sprintf(ubuntuPath, assembly.AppFileName)
			} else if env.System() == env.Linux_CentOS {
				bashFile = fmt.Sprintf(centosPath, assembly.AppFileName)
			}
			if !futils.FileExist(bashFile) {
				fmt.Printf("服务[%s]...\n", echo.Red("未部署"))
				os.Exit(0)
			}
			shell.ExeCommand("systemctl", "stop", assembly.AppFileName)    // 停止服务
			shell.ExeCommand("systemctl", "disable", assembly.AppFileName) // 取消自启动
			os.Remove(bashFile)                                            // 删除服务配置
			os.Remove(fmt.Sprintf("/etc/systemd/system/multi-user.target.wants/%s.service", assembly.AppFileName))
			os.Remove(fmt.Sprintf("/etc/systemd/system/%s.service", assembly.AppFileName))
			shell.ExeCommand("systemctl", "daemon-reload") // 重新加载服务配置
			fmt.Printf("服务[%s]...\n", echo.Green("已卸载"))
			os.Exit(0)
		},
	}
)
