package server

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"server/common/assembly"
	"server/common/config"
	_ "server/migration/versions" // 这个是数据迁移的 不能删
	"server/sugar/echo"
	"server/sugar/env"
	g "server/sugar/groute"
	"server/sugar/log"
	"server/sugar/network"
	"server/sugar/system"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"server/app/jobs"
	"server/common/initialize"
)

var (
	configYml string
	// StartCmd 命令行
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      assembly.AppFileName + " server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	err := assembly.RunOfOnec()
	if err != nil {
		log.Print(echo.Red("service instance is running..."))
		log.Print(echo.Red("service exit, code = 100"))
		os.Exit(100)
	}
	//1. 读取配置
	config.Setup(
		configYml,
		// 初始化运行模式
		initialize.InitRunMode,
		// 初始化日志
		initialize.InitLogger,
		// 打印logo
		initialize.PrintLogo,
		// 1.初始化SQLite连接
		initialize.InitSQLDB,
		// 1.初始化SQLite连接
		initialize.InitRedisDB,
		// 2.执行数据升级&迁移
		initialize.InitMigration,
		// 开发模式显示菜单
		initialize.InitDevelopmentMenu,
		// 初始化 redis
		initialize.InitRedisDB,
		// 时序数据库、队列、缓存 初始化
		initialize.InitCache,
		// 任务自动化
		jobs.Setup,
	)

}

func run() error {
	engine := gin.New()
	if env.ModeIs(env.Production) {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	// 注册路由
	g.Register(engine, GetRootRouter())
	log.Info(echo.Green("Server run at:"))
	for _, ip := range network.LocalIpAddres() {
		log.Infof("- %s://%s:%d/", config.Application.GetHttpProtocol(), ip, config.Application.Port)
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port),
		Handler: engine,
	}
	go func() {
		// 服务连接
		if config.Application.Https {
			if err := srv.ListenAndServeTLS(config.Application.CertFile, config.Application.KeyFile); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		}
	}()
	log.Info("Enter Control + C Shutdown Server")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	system.WaitQuitSignal()
	log.Info(echo.Yellow("The server is shut down .... "))
	// 等待上下文退出
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")
	return nil
}
