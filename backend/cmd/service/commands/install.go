package commands

import (
	"backend/core/sdk/console"
	"backend/core/sdk/pkg"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	ubuntuPath = "/etc/init.d/siteweb-manager"
	centosPath = "/etc/systemd/system/siteweb-manager.service"
)

const ubuntuBash = `#!/bin/bash
### BEGIN INIT INFO
# Provides:          Siteweb-Manager
# Required-Start:    $local_fs $network $remote_fs
# Required-Stop:     $local_fs $network $remote_fs
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Start Siteweb-Manager daemon at boot time
# Description:       Start Siteweb-Manager daemon at boot time
### END INIT INFO

usage() {
  echo " usage:$0 {start|stop|restart|status} "
}

start() {
	siteweb-manager service start
}

stop() {
	siteweb-manager service stop
}

status() {
    siteweb-manager service status
}

restart() {
	siteweb-manager service restart
}

#main function
case $1 in
  start)
     start
     ;;
  stop)
     stop
     ;;
  restart)
     restart
     ;;
  status)
     status
     ;;
  *)
     usage
     ;;
esac
exit 0
`

const centosBash = `
[Unit]
Description=siteweb-manager
After=network.target

[Service]
Type=forking

#启动脚本路径
ExecStart=$(__ASSEMBLY) service start

#重启脚本路径
ExecReload=$(__ASSEMBLY) service restart

#停止脚本路径
ExecStop=$(__ASSEMBLY) service top

# 停止超时时间，如果不能在指定时间内停止，将通过SIGKILL强制终止
KillSignal=SIGQUIT
TimeoutStopSec=5

# systemd停止服务的方式
KillMode=mixed

# 服务不正常退出后重启
Restart=on-failure

# 表示给服务分配独立的临时空间
PrivateTmp=true

[Install]
WantedBy=multi-user.target
`

var (
	InstallCmd = &cobra.Command{
		Use:     "install",
		Short:   "install siteweb-manager service",
		Example: "siteweb-manager service install",
		Run: func(cmd *cobra.Command, args []string) {
			script := ""
			bashFile := ""
			if pkg.IsUbuntu() {
				script = ubuntuBash
				bashFile = ubuntuPath
			} else if pkg.IsCentOS() {
				script = strings.Replace(centosBash, "$(__ASSEMBLY)", pkg.AssemblyFile(), -1)
				bashFile = centosPath
			}
			if pkg.FileExist(bashFile) {
				fmt.Printf("服务[%s]...\n", console.Red("已部署"))
				os.Exit(0)
			}
			pkg.ExeCommand("ln", "-s", "-f", pkg.AssemblyFile(), "/usr/bin/siteweb-manager") // 创建全局命令
			os.WriteFile(bashFile, []byte(script), 0777)                                     //写入服务配置文件
			pkg.ExeCommand("chmod", "777", bashFile)                                         // 给777权限，其实上一步已经给了
			time.Sleep(time.Second)
			pkg.ExeCommand("systemctl", "daemon-reload") //重新加载配置
			time.Sleep(time.Second)
			pkg.ExeCommand("systemctl", "enable", "siteweb-manager") //设置开机启动
			fmt.Printf("服务[%s]...\n", console.Green("已部署"))
			os.Exit(0)
		},
	}
)
