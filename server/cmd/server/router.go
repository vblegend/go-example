package server

import (
	adminApis "server/app/admin/apis"
	"server/app/jobs"
	jobModels "server/app/jobs/models"
	"server/common/config"
	"server/common/middleware"
	"server/common/middleware/jwt"
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
	mid, err := jwtauth.NewJWT(&jwt.Standard{}, config.Jwt.Timeout, config.Jwt.Secret)
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
			plugs.StaticFileServe("/", config.Web.Root),
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
			Url: "/api",
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
					// Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						api := adminApis.MenuApi{}
						r.GET("/menutree", api.GetMenuTree)
						// user := adminapi.SysUser{}
						// r.GET("/getinfo", user.GetInfo)
					},
				},
				g.Router{ //菜单管理
					Url: "/menu",
					Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						api := adminApis.MenuApi{}
						// api := adminapi.SysMenu{}
						r.GET("/", api.GetMenuTree)
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
						// 列出所有Job模型
						r.GET("/list", restful.ListHander(&jobModels.SysJob{}))
						// 分页+条件查询Job模型
						r.GET("", restful.WherePageHander(&jobModels.SysJob{}, &model.Pagination{}, &jobModels.SysJob{}))
						// 根据ID 获取Job模型
						r.GET("/:JobId", restful.WhereFirstHander(&jobModels.SysJobIndex{}, &jobModels.SysJob{}))
						// 创建 job
						r.POST("", restful.CreateHander(func(create restful.HandlerActionCallBack, api *restful.Api) {
							model := jobModels.SysJob{}
							err := create(&model)
							if err == nil {
								jobs.ConfigJob(model)
							}
						}))
						// 修改job
						r.PUT("", restful.UpdateHander(func(update restful.HandlerActionCallBack, api *restful.Api) {
							model := jobModels.SysJob{}
							err := update(&model)
							if err == nil {
								jobs.ConfigJob(model)
							}
						}))
						// 删除job
						r.DELETE("/:JobId", restful.DeleteHander(func(delete restful.HandlerQueryCallBack, api *restful.Api) {
							query := jobModels.SysJobIndex{}
							model := jobModels.SysJob{}
							err := delete(query, model)
							if err == nil {
								jobs.StopJob(model.JobID)
							}
						}))
					},
				},
				g.Router{ // 定时任务操作
					Url: "/task",
					// Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						r.GET("/stop/:JobId", restful.ActionHander(func(call restful.HandlerActionCallBack, api *restful.Api) {
							model := jobModels.SysJobIndex{}
							err := call(&model)
							if err == nil {
								jobs.StopJob(model.JobID)
							}
						}))
						r.GET("/start/:JobId", restful.ActionHander(func(call restful.HandlerActionCallBack, api *restful.Api) {
							model := jobModels.SysJobIndex{}
							err := call(&model)
							if err == nil {
								// api.OK()
								// jobs.StartJob(model.JobID)
							}
						}))
					},
				},
			},
		},
	}

}
