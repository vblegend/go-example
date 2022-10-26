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
	UnInstallCmd = &cobra.Command{
		Use:     "uninstall",
		Short:   "uninstall app service",
		Example: fmt.Sprintf("%s service uninstall", global.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {

			bashFile := ""
			if pkg.IsUbuntu() {
				bashFile = fmt.Sprintf(ubuntuPath, global.AppFileName)
			} else if pkg.IsCentOS() {
				bashFile = fmt.Sprintf(centosPath, global.AppFileName)
			}
			if !pkg.FileExist(bashFile) {
				fmt.Printf("服务[%s]...\n", console.Red("未部署"))
				os.Exit(0)
			}
			pkg.ExeCommand("systemctl", "stop", global.AppFileName)    // 停止服务
			pkg.ExeCommand("systemctl", "disable", global.AppFileName) // 取消自启动
			os.Remove(bashFile)                                        // 删除服务配置
			os.Remove(fmt.Sprintf("/etc/systemd/system/multi-user.target.wants/%s.service", global.AppFileName))
			os.Remove(fmt.Sprintf("/etc/systemd/system/%s.service", global.AppFileName))
			pkg.ExeCommand("systemctl", "daemon-reload") // 重新加载服务配置
			fmt.Printf("服务[%s]...\n", console.Green("已卸载"))
			os.Exit(0)
		},
	}
)
