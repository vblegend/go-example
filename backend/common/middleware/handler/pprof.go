package handler

import (
	"backend/core/sdk/config"
	"backend/core/sdk/pkg"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

const (
	pprof_User     = "admin"
	pprof_Password = "123456"
)

func InitPPROF(r gin.IRouter) {
	//性能调优监视 TODO Gin自主隐藏，待优化
	authStr := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", pprof_User, pprof_Password))))
	pprofGroup := r.Group("/debug", func(c *gin.Context) {
		if config.ApplicationConfig.Mode == pkg.Production {
			auth := c.Request.Header.Get("Authorization")
			if auth != authStr {
				c.Header("www-Authenticate", "Basic")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		c.Next()
	})
	pprof.RouteRegister(pprofGroup, "/")
}
