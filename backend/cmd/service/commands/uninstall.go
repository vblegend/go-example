package commands

import (
	"backend/core/sdk/console"
	"backend/core/sdk/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	UnInstallCmd = &cobra.Command{
		Use:     "uninstall",
		Short:   "uninstall siteweb-manager service",
		Example: "siteweb-manager service uninstall",
		Run: func(cmd *cobra.Command, args []string) {

			bashFile := ""
			if pkg.IsUbuntu() {
				bashFile = ubuntuPath
			} else if pkg.IsCentOS() {
				bashFile = centosPath
			}
			if !pkg.FileExist(bashFile) {
				fmt.Printf("服务[%s]...\n", console.Red("未部署"))
				os.Exit(0)
			}
			pkg.ExeCommand("systemctl", "stop", "siteweb-manager")    // 停止服务
			pkg.ExeCommand("systemctl", "disable", "siteweb-manager") // 取消自启动
			os.Remove(bashFile)                                       // 删除服务配置
			os.Remove("/etc/systemd/system/multi-user.target.wants/siteweb-manager.service")
			os.Remove("/etc/systemd/system/siteweb-manager.service")
			pkg.ExeCommand("systemctl", "daemon-reload") // 重新加载服务配置
			fmt.Printf("服务[%s]...\n", console.Green("已卸载"))
			os.Exit(0)
		},
	}
)
