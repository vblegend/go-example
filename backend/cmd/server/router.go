package server

import (
	jobModels "backend/app/jobs/models"
	"backend/common/config"
	"backend/common/middleware"
	g "backend/core/groute"
	"backend/core/jwtauth"
	"backend/core/log"
	"backend/core/plugs"
	"backend/core/restful"
	"backend/core/ws"

	"backend/core/model"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createAuthMiddleware() *jwtauth.GinJWTMiddleware {
	mid, err := jwtauth.NewJWT(&jwtauth.Standard{}, config.Jwt.Timeout, config.Jwt.Secret)
	if err != nil {
		panic("初始化身份认证中间件失败。")
	}
	return mid
}

func GetRootRouter() g.Routers {
	var authMiddleware = createAuthMiddleware()
	return g.Routers{
		g.Router{ // 公共根路由
			Url: "",
			Use: g.Use(
				middleware.CustomError, // 自定义异常处理
				plugs.NewHttpsHandler(config.Application.Https, config.Application.Domain, uint(config.Application.Port)),
				plugs.RequestLogOut(log.GetLogger(), log.TraceLevel), // 请求日志
				plugs.TraceId("requestId", uuid.NewString),           // 请求UUID
				plugs.WithContextDB("default"),                       // 数据连接
				plugs.NoCache,                                        // 禁用缓存
				plugs.Options,                                        // 跨域请求
				plugs.Secure,                                         // https相关
			),
			Children: GetApiRouter(authMiddleware),
		},
	}
}

func GetApiRouter(authMiddleware *jwtauth.GinJWTMiddleware) g.Routers {
	return g.Routers{
		g.Router{ // websocket 连接
			Url: "debug",
			Use: g.Use(plugs.WWWAuthenticator("admin", "admin")),
			Handle: func(r gin.IRoutes) {
				pprof.RouteRegister(r.(*gin.RouterGroup), "/")
			},
		},
		g.Router{ // websocket 连接
			Url: "/ws",
			Handle: func(r gin.IRoutes) {
				r.GET("", ws.Default.AcceptHandler)
			},
		},
		g.Router{
			Url: "/api/v1",
			Handle: func(r gin.IRoutes) {
				// 登录接口
				r.POST("login", authMiddleware.LoginHandler)
				// 刷新Token
				r.GET("/refresh_token", authMiddleware.RefreshHandler)
				// 注销登出
				r.POST("logout", jwtauth.LogOut)
			},
			Children: g.Routers{
				g.Router{ //左侧菜单
					Url: "",
					Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						// api := adminapi.SysMenu{}
						// user := adminapi.SysUser{}
						// r.GET("/menutree", api.GetMenuTreeSelect)
						// r.GET("/getinfo", user.GetInfo)
					},
				},
				g.Router{ //菜单管理
					Url: "/menu",
					Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						// api := adminapi.SysMenu{}
						// r.GET("", api.GetPage)
						// r.GET("/:id", api.Get)
						// r.POST("", api.Insert)
						// r.PUT("/:id", api.Update)
						// r.DELETE("", api.Delete)
					},
				},
				g.Router{ // 定时任务相关
					Url: "/job",
					// Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						r.GET("", restful.WherePageHander(&jobModels.SysJob{}, &model.Pagination{}, &jobModels.SysJob{}))
						r.GET("/:JobId", restful.WhereFirstHander(&jobModels.SysJobIndex{}, &jobModels.SysJob{}))
						r.POST("", restful.CreateHander(&jobModels.SysJob{}, func(model interface{}) {}))
						r.PUT("", restful.UpdateHander(&jobModels.SysJob{}, func(model interface{}) {}))
						r.DELETE("/:JobId", restful.DeleteHander(&jobModels.SysJobIndex{}, &jobModels.SysJob{}, func(id interface{}) {}))
					},
				},
				g.Router{ // 定时任务操作
					Url: "/task",
					// Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						r.GET("/stop/:JobId", restful.ActionHander(&jobModels.SysJobIndex{}, func(object interface{}) {
							// 停止任务
						}))
						r.GET("/start/:JobId", restful.ActionHander(&jobModels.SysJobIndex{}, func(object interface{}) {
							// 开始任务
						}))
					},
				},
			},
		},
		g.Router{
			Url: "",
			Handle: func(r gin.IRoutes) {
				r.GET("/", func(c *gin.Context) {
					c.Redirect(http.StatusMovedPermanently, "/login")
				})
				index := func(c *gin.Context) { c.File("./static/www/index.html") }
				r.GET("/401", index)
				r.GET("/404", index)
				r.GET("/login", index)

				r.Static("/js", "static/www/js")
				r.Static("/css", "static/www/css")
				r.Static("/fonts", "static/www/fonts")
				r.Static("/img", "static/www/img")
			},
		},
	}

}
