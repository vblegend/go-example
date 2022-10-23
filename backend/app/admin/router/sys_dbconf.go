package router

import (
	"backend/app/admin/apis"
	jwt "backend/core/sdk/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSyMysqlConfigureRouter)
}

// 需认证的路由代码
func registerSyMysqlConfigureRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysDBConfig{}
	r := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("/mysqlconf", api.GetMYSQLConfigure)
		r.PUT("/mysqlconf", api.SaveMYSQLConfigure)

		r.GET("/redisconf", api.GetRedisConfigure)
		r.PUT("/redisconf", api.SaveRedisConfigure)

	}

}
