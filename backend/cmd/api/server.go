package api

import (
	"context"
	"fmt"
	"path/filepath"

	"backend/core/logger"
	"backend/core/sdk"
	"backend/core/sdk/api"
	"backend/core/sdk/config"
	"backend/core/sdk/console"
	"backend/core/sdk/pkg"
	httpRuntime "backend/core/sdk/runtime"
	"backend/core/storage/queue"
	"backend/migration"
	_ "backend/migration/versions" // 这个是数据迁移的 不能删
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"backend/app/admin/models"
	"backend/app/admin/router"
	"backend/app/jobs"
	"backend/common/database"
	"backend/common/global"
	common "backend/common/middleware"
	"backend/common/middleware/handler"
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

var AppRouters = make([]func(), 0)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")
	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	AppRouters = append(AppRouters, router.InitRouter)
}

func printLogo() {
	fmt.Println(console.Yellow(strings.Join(global.LogoContent, "\n")) + fmt.Sprintf(" %s %s (%s)", console.Green(config.ApplicationConfig.Mode), console.Red("V"+global.Version), global.BuildTime))
}

func PatchSystemLibs(sourceDir string, destDir string) {
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		filename := filepath.Base(path)
		destfile := filepath.Join(destDir, filename)
		if !pkg.FileExist(destfile) {
			logger.Infof("Patch System Lib File %s Success...", console.Green(destfile))
			pkg.ExeCommand("/bin/cp", "-rf", path, destfile)
		}
		return nil
	})
}

func setup() {
	err := pkg.RunOfOnec()
	if err != nil {
		fmt.Println(console.Red("service instance is running..."))
		fmt.Println(console.Red("service exit, code = 100"))
		os.Exit(100)
	}
	if runtime.GOOS == "linux" {
		pkg.MkDirIfNotExist(filepath.Join(pkg.AssemblyDir(), "./temp"))
		pkg.MkDirIfNotExist(filepath.Join(pkg.AssemblyDir(), "./temp/access"))
		pkg.ExeCommand("chmod", "777", "-R", filepath.Join(pkg.AssemblyDir(), "./bin"))
		pkg.ExeCommand("chmod", "777", "-R", filepath.Join(pkg.AssemblyDir(), "./temp"))
	}
	// 安装依赖库补丁
	PatchSystemLibs(filepath.Join(pkg.AssemblyDir(), "bin/libs"), "/usr/lib/")

	//1. 读取配置
	config.Setup(
		configYml,
		printLogo,
		// 1.初始化SQLite连接
		database.InitSQLiteDB,
		// 2.执行数据升级&迁移
		migration.DataBaseMigrate,
		// 开发模式显示菜单
		database.Development,
		// 初始化mysql
		database.InitMysqlDB,
		// 初始化 redis
		database.InitRedisDB,
		// 时序数据库、队列、缓存 初始化
		storage.Setup,
	)
	//注册监听函数
	queue := sdk.Runtime.GetMemoryQueue("")
	queue.Register(global.LoginLog, models.SaveLoginLog)
	queue.Register(global.OperateLog, models.SaveOperaLog)
	queue.Register(global.ApiCheck, models.SaveSysApi)
	go queue.Run()

}

func run() error {
	if config.ApplicationConfig.Mode == pkg.Production {
		gin.SetMode(gin.ReleaseMode)
	}
	initRouter()

	for _, f := range AppRouters {
		f()
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: sdk.Runtime.GetEngine(),
	}

	// 设置容器
	db := sdk.Runtime.GetDbByKey("*")
	jobs.InitJob()
	jobs.Setup(db)

	// 监控数据初始化
	logger.Info(`starting api server...`)
	if apiCheck {
		var routers = sdk.Runtime.GetRouter()
		q := sdk.Runtime.GetMemoryQueue("")
		mp := make(map[string]interface{}, 0)
		mp["List"] = routers
		message := queue.NewQueueMessage("", global.ApiCheck, mp)
		err := q.Append(message)
		if err != nil {
			logger.Errorf("Append message error, %s \n", err.Error())
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		// 服务连接
		if config.SslConfig.Enable {
			if err := srv.ListenAndServeTLS(config.SslConfig.Pem, config.SslConfig.Key); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		}
	}()
	logger.Info(console.Green("Server run at:"))
	ipaddress := pkg.GetIpAddress()
	for _, ip := range ipaddress {
		logger.Infof("- http://%s:%d/", ip, config.ApplicationConfig.Port)
	}
	logger.Info("Enter Control + C Shutdown Server")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 5)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-quit
	logger.Info(console.Yellow(fmt.Sprintf("The server is shut down from [%v] .... ", s)))
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")

	return nil
}

var Router httpRuntime.Router

func initRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		sdk.Runtime.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	handler.InitPPROF(r)

	if config.SslConfig.Enable {
		r.Use(handler.TlsHandler())
	}
	//r.Use(middleware.Metrics())
	r.Use(common.RequestId(pkg.TrafficKey)).
		Use(api.SetRequestLogger)

	// =====================
	// Http Static Service
	// =====================
	initHttpStatic(r)
	common.InitMiddleware(r)
}

func initHttpStatic(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})
	r.Static("/js", "static/www/js")
	r.Static("/css", "static/www/css")
	r.Static("/fonts", "static/www/fonts")
	r.Static("/img", "static/www/img")
	// ====================
	r.StaticFile("/admin", "./static/www/index.html")
	r.StaticFile("/dashboard", "./static/www/index.html")
	r.StaticFile("/login", "./static/www/index.html")
	r.StaticFile("/profile", "./static/www/index.html")
	r.StaticFile("/auth-redirect", "./static/www/index.html")
	r.StaticFile("/redirect", "./static/www/index.html")
	r.StaticFile("/401", "./static/www/index.html")
	r.StaticFile("/404", "./static/www/index.html")

	// ====================
	index := func(c *gin.Context) {
		c.File("./static/www/index.html")
	}

	r.GET("/admin/*filepath", index)
	r.GET("/profile/*filepath", index)
	r.GET("/schedule/*filepath", index)
	r.GET("/sys-tools/*filepath", index)
	r.GET("/dev-tools/*filepath", index)
	r.GET("/redirect/*filepath", index)
	r.GET("/data/*filepath", index)
	r.GET("/service/*filepath", index)
	r.GET("/exception/*filepath", index)
	r.GET("/license", index)
	r.GET("/log/*filepath", index)
}
