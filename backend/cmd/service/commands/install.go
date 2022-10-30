package commands

import (
	"backend/common/assembly"
	"backend/core/echo"
	"backend/core/env"
	"backend/core/futils"
	"backend/core/shell"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const (
	ubuntuPath = "/etc/init.d/%s"
	centosPath = "/etc/systemd/system/%s.service"
)

const ubuntuBash = `#!/bin/bash
### BEGIN INIT INFO
# Provides:          {{.AppName}}
# Required-Start:    $local_fs $network $remote_fs
# Required-Stop:     $local_fs $network $remote_fs
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Start {{.AppName}} daemon at boot time
# Description:       Start {{.AppName}} daemon at boot time
### END INIT INFO

usage() {
  echo " usage:$0 {start|stop|restart|status} "
}

start() {
	{{.AppName}} service start
}

stop() {
	{{.AppName}} service stop
}

status() {
    {{.AppName}} service status
}

restart() {
	{{.AppName}} service restart
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
Description={{.AppName}}
After=network.target

[Service]
Type=forking

#启动脚本路径
ExecStart={{.AppPath}} service start

#重启脚本路径
ExecReload={{.AppPath}} service restart

#停止脚本路径
ExecStop={{.AppPath}} service top

# 停止超时时间，如果不能在指定时间内停止，将通过SIGKILL强制终止
KillSignal=SIGQUIT
TimeoutStopSec=5

# systemd停止服务的方式
KillMode=mixed

# 服务不正常退出后重启
Restart=on-failure

# 表示给服务分配独立的临时空间
PrivateTmp=false

[Install]
WantedBy=multi-user.target
`

func parseScriptTemplate(stemplate string) string {
	t, err := template.New("").Parse(stemplate)
	if err != nil {
		panic("error")
	}
	buf := bytes.Buffer{}
	t.Execute(&buf, map[string]interface{}{
		"AppName": assembly.AppFileName,
		"AppPath": env.AssemblyFile,
	})
	return buf.String()
}

var (
	InstallCmd = &cobra.Command{
		Use:     "install",
		Short:   "install app service",
		Example: fmt.Sprintf("%s service install", assembly.AppFileName),
		Run: func(cmd *cobra.Command, args []string) {
			script := ""
			bashFile := ""
			if env.System() == env.Linux_Ubuntu {
				script = parseScriptTemplate(ubuntuBash)
				bashFile = fmt.Sprintf(ubuntuPath, assembly.AppFileName)
			} else if env.System() == env.Linux_CentOS {
				script = parseScriptTemplate(centosBash)
				bashFile = fmt.Sprintf(centosPath, assembly.AppFileName)
			}
			if futils.FileExist(bashFile) {
				fmt.Printf("服务[%s]...\n", echo.Red("已部署"))
				os.Exit(0)
			}
			shell.ExeCommand("ln", "-s", "-f", env.AssemblyFile, "/usr/bin/"+assembly.AppFileName) // 创建全局命令
			os.WriteFile(bashFile, []byte(script), 0777)                                           //写入服务配置文件
			shell.ExeCommand("chmod", "777", bashFile)                                             // 给777权限，其实上一步已经给了
			time.Sleep(time.Second)
			shell.ExeCommand("systemctl", "daemon-reload") //重新加载配置
			time.Sleep(time.Second)
			shell.ExeCommand("systemctl", "enable", assembly.AppFileName) //设置开机启动
			fmt.Printf("服务[%s]...\n", echo.Green("已部署"))
			os.Exit(0)
		},
	}
)
