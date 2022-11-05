package server

import (
	"fmt"
	jobModels "server/app/jobs/models"
	"server/common/config"
	"server/common/middleware"
	g "server/sugar/groute"
	"server/sugar/jwtauth"
	"server/sugar/log"
	"server/sugar/plugs"
	"server/sugar/restful"
	"server/sugar/ws"

	"server/sugar/model"

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

func GetRootRouter() g.Router {
	var authMiddleware = createAuthMiddleware()
	return g.Router{
		Url: "", // 根路由
		Use: g.Use(
			middleware.CustomError, // 自定义异常处理
			plugs.NewIPBlockerHandler(func(s string) bool {
				fmt.Println(s)
				return true
			}), //IP拦截器
			plugs.NewHttpsHandler(config.Web.Https, config.Web.Domain, uint(config.Web.Port)), // https
			plugs.RequestLogOut(log.GetLogger(), log.TraceLevel),                              // 请求日志
			plugs.TraceID("requestId", uuid.NewString),                                        // 请求UUID
			plugs.WithContextDB(config.DefaultDB),                                             // 数据连接
			plugs.NoCache,                                                                     // 禁用缓存
			plugs.Options,                                                                     // 跨域请求
			plugs.Secure,                                                                      // 安全相关
			plugs.Limit(10),                                                                   // 并发数控制
			plugs.StaticFileServe("/", "./static/www"),
		),
		Children: GetApiRouter(authMiddleware),
	}
}

func GetApiRouter(authMiddleware *jwtauth.GinJWTMiddleware) g.Routers {
	return g.Routers{
		g.Router{
			Url: "debug", // pprof debug
			Use: g.Use(plugs.WWWAuthenticator("admin", "admin")),
			Handle: func(r gin.IRoutes) {
				pprof.RouteRegister(r.(*gin.RouterGroup), "/")
			},
		},
		g.Router{
			Url: "/ws", // websocket 连接
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
					Url: "/",
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
	}

}
