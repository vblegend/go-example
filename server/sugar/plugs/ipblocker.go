package plugs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IPBlockerHandler IP拦截器处理函数
type IPBlockerHandler func(string) bool

// NewIPBlockerHandler 创建一个IP拦截器
func NewIPBlockerHandler(handler IPBlockerHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ok := handler(c.ClientIP())
		if !ok {
			c.AbortWithStatus(http.StatusBadGateway)
			return
		}
		c.Next()
	}
}
