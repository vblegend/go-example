package restful

import (
	"strings"

	"github.com/gin-gonic/gin"

	"backend/core/log"
	"backend/core/sdk/pkg"
)

type loggerKey struct{}

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *log.Helper {
	var logv *log.Helper
	l, ok := c.Get(pkg.LoggerKey)
	if ok {
		ok = false
		logv, ok = l.(*log.Helper)
		if ok {
			return logv
		}
	}
	return nil
}

// SetRequestLogger 设置logger中间件
func SetRequestLogger(c *gin.Context) {
	requestId := pkg.GenerateMsgIDFromContext(c)
	log := log.NewHelper(log.GetLogger()).WithFields(map[string]interface{}{
		strings.ToLower(pkg.TrafficKey): requestId,
	})
	c.Set(pkg.LoggerKey, log)
}
