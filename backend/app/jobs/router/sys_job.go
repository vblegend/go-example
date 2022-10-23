package router

import (
	"backend/app/jobs/apis"
	models2 "backend/app/jobs/models"
	dto2 "backend/app/jobs/service/dto"
	"backend/common/actions"
	jwt "backend/core/sdk/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysJobRouter)
}

// 需认证的路由代码
func registerSysJobRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/sysjob").Use(authMiddleware.MiddlewareFunc())
	{
		sysJob := &models2.SysJob{}
		r.GET("", actions.IndexAction(sysJob, new(dto2.SysJobSearch), func() interface{} {
			list := make([]models2.SysJob, 0)
			return &list
		}))
		r.GET("/:id", actions.ViewAction(new(dto2.SysJobById), func() interface{} {
			return &dto2.SysJobItem{}
		}))
		r.POST("", actions.CreateAction(new(dto2.SysJobControl)))
		r.PUT("", actions.UpdateAction(new(dto2.SysJobControl)))
		r.DELETE("", actions.DeleteAction(new(dto2.SysJobById)))
	}
	sysJob := apis.SysJob{}

	v1.GET("/job/remove/:id", sysJob.RemoveJobForService)
	v1.GET("/job/start/:id", sysJob.StartJobForService)
}
