package router

import (
	jwt "backend/core/sdk/pkg/jwtauth"

	"github.com/gin-gonic/gin"

	"backend/app/admin/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysApiRouter)
}

// registerSysApiRouter
func registerSysApiRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysApi{}
	r := v1.Group("/sys-api").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.PUT("/:id", api.Update)
	}
}
