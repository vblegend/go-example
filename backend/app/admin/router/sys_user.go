package router

import (
	"backend/app/admin/apis"
	jwt "backend/core/sdk/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
}

// 需认证的路由代码
func registerSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysUser{}
	r := v1.Group("/sys-user").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}

	user := v1.Group("/user").Use(authMiddleware.MiddlewareFunc())
	{
		user.GET("/profile", api.GetProfile)
		user.POST("/avatar", api.InsetAvatar)
		user.PUT("/pwd/set", api.UpdatePwd)
		user.PUT("/pwd/reset", api.ResetPwd)
		user.PUT("/status", api.UpdateStatus)
	}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		v1auth.GET("/getinfo", api.GetInfo)
	}
}
