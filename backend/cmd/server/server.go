package server

import (
	"context"
	"fmt"

	"backend/core/console"
	g "backend/core/groute"
	"backend/core/log"
	"backend/core/sdk/config"
	"backend/core/sdk/pkg"
	"backend/core/system"
	"backend/migration"
	_ "backend/migration/versions" // 这个是数据迁移的 不能删
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"backend/app/jobs"
	"backend/common/database"
	"backend/common/global"
	"backend/common/storage"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "backend server -c config/settings.yml",
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
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")
}

func printLogo() {
	log.Print(console.Yellow(strings.Join(global.LogoContent, "\n")) + fmt.Sprintf(" %s %s (%s)\n", console.Green(config.ApplicationConfig.Mode), console.Red("V"+global.Version), global.BuildTime))

}

func setup() {
	err := pkg.RunOfOnec()
	if err != nil {
		log.Print(console.Red("service instance is running..."))
		log.Print(console.Red("service exit, code = 100"))
		os.Exit(100)
	}
	//1. 读取配置
	config.Setup(
		configYml,
		printLogo,
		// 1.初始化SQLite连接
		database.InitDatabase,
		// 2.执行数据升级&迁移
		migration.DataBaseMigrate,
		// 开发模式显示菜单
		database.Development,
		// 初始化 redis
		database.InitRedisDB,
		// 时序数据库、队列、缓存 初始化
		storage.Setup,
		// 任务自动化
		jobs.Setup,
	)

}

func run() error {
	engine := gin.New()
	if config.ApplicationConfig.Mode == pkg.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	// 注册路由
	g.Register(engine, GetRootRouter())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: engine,
	}
	log.Info(console.Green("Server run at:"))
	for _, ip := range pkg.GetIpAddress() {
		log.Infof("- %s://%s:%d/", config.ApplicationConfig.GetHttpProtocol(), ip, config.ApplicationConfig.Port)
	}
	go func() {
		// 服务连接
		if config.ApplicationConfig.Https {
			if err := srv.ListenAndServeTLS(config.ApplicationConfig.CertFile, config.ApplicationConfig.KeyFile); err != nil && err != http.ErrServerClosed {
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
	log.Info(console.Yellow("The server is shut down .... "))
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")
	return nil
}
