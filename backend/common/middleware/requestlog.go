package middleware

import (
	"backend/core/log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogOut(loger log.Logger, level log.Level) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		defer func() {
			endTime := time.Now()
			responseTime := endTime.Sub(startTime)
			// 	"server":         c.Request.Host,
			// 	"content_type":   c.Request.Header.Get("Content-Type"),
			// 	"user_agent":     c.Request.UserAgent(),
			loger.Logf(level, "[%s] [%s] [%s] [%s] [%v] %s", responseTime, c.ClientIP(), c.Request.Proto, c.Request.Method, c.Writer.Status(), c.Request.URL.Path)
		}()
		c.Next()
	}
}
