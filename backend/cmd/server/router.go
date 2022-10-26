package server

import (
	adminapi "backend/app/admin/apis"
	jobApi "backend/app/jobs/apis"
	jobModel "backend/app/jobs/models"
	jboDto "backend/app/jobs/service/dto"
	"backend/common/actions"
	"backend/common/middleware"
	coreApi "backend/core/api"
	g "backend/core/groute"
	"backend/core/jwtauth"
	"backend/core/sdk/pkg"
	"backend/core/sdk/pkg/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createAuthMiddleware() *jwtauth.GinJWTMiddleware {
	mid, err := jwtauth.NewJWT(&jwtauth.Standard{})
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
			Use: g.Use(middleware.RequestId(pkg.TrafficKey), // 请求ID
				coreApi.SetRequestLogger, // 请求日志
				middleware.WithContextDb, // 数据连接
				middleware.NoCache,       // 禁用缓存
				middleware.Options,       // 跨域请求
				middleware.Secure,        // https相关
				middleware.CustomError,   // 自定义异常处理
			),
			Children: GetApiRouter(authMiddleware),
		},
	}
}

func GetApiRouter(authMiddleware *jwtauth.GinJWTMiddleware) g.Routers {
	return g.Routers{
		g.Router{ // websocket 连接
			Url: "/ws/:id/:channel",
			Handle: func(r gin.IRoutes) {
				r.GET("", ws.WebsocketManager.WsClient)
			},
		},
		g.Router{ // websocket 注销
			Url: "/wslogout/:id/:channel",
			Handle: func(r gin.IRoutes) {
				r.GET("", ws.WebsocketManager.UnWsClient)
			},
		},
		g.Router{
			Url: "/api/v1",
			Handle: func(r gin.IRoutes) {
				// 登录接口
				r.POST("login", authMiddleware.LoginHandler)
				// 刷新Token
				r.GET("/refresh_token", authMiddleware.RefreshHandler)
			},
			Children: g.Routers{
				g.Router{ //左侧菜单 登出
					Url: "",
					Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						api := adminapi.SysMenu{}
						r.GET("/roleMenuTreeselect/:roleId", api.GetMenuTreeSelect)
						r.POST("/logout", jwtauth.LogOut)
						user := adminapi.SysUser{}
						r.GET("/getinfo", user.GetInfo)
					},
				},
				g.Router{ // 系统API
					Url: "/sys-api",
					Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
					},
					Children: g.Routers{
						g.Router{ //菜单管理
							Url: "/menu",
							Handle: func(r gin.IRoutes) {
								api := adminapi.SysMenu{}
								r.GET("", api.GetPage)
								r.GET("/:id", api.Get)
								r.POST("", api.Insert)
								r.PUT("/:id", api.Update)
								r.DELETE("", api.Delete)
							},
						},
						g.Router{ //左侧菜单
							Url: "/menurole",
							Handle: func(r gin.IRoutes) {
								api := adminapi.SysMenu{}
								r.GET("", api.GetMenuRole)
							},
						},
						g.Router{ // 用户管理
							Url: "/user",
							Handle: func(r gin.IRoutes) {
								api := adminapi.SysUser{}
								r.GET("/profile", api.GetProfile)
								r.POST("/avatar", api.InsetAvatar)
								r.PUT("/pwd/set", api.UpdatePwd)
								r.PUT("/pwd/reset", api.ResetPwd)
								r.PUT("/status", api.UpdateStatus)
							},
						},
					},
				},
				g.Router{ // 定时任务相关
					Url: "/sysjob",
					Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						sysJob := &jobModel.SysJob{}
						r.GET("", actions.IndexAction(sysJob, new(jboDto.SysJobSearch), func() interface{} {
							list := make([]jobModel.SysJob, 0)
							return &list
						}))
						r.GET("/:id", actions.ViewAction(new(jboDto.SysJobById), func() interface{} {
							return &jboDto.SysJobItem{}
						}))
						r.POST("", actions.CreateAction(new(jboDto.SysJobControl)))
						r.PUT("", actions.UpdateAction(new(jboDto.SysJobControl)))
						r.DELETE("", actions.DeleteAction(new(jboDto.SysJobById)))
					},
				},
				g.Router{ // 定时任务操作
					Url: "/job",
					Use: g.Use(authMiddleware.MiddlewareFunc()),
					Handle: func(r gin.IRoutes) {
						sysJob := jobApi.SysJob{}
						r.GET("/remove/:id", sysJob.RemoveJobForService)
						r.GET("/start/:id", sysJob.StartJobForService)
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
