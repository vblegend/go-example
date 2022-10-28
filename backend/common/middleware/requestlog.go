package middleware

import (
	"backend/core/log"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger(c *gin.Context) {
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		responseTime := endTime.Sub(startTime)
		// logField := map[string]interface{}{
		// 	"uri":            c.Request.URL.Path,
		// 	"start":          startTime.Format("2006-01-02 15:04:05"),
		// 	"server":         c.Request.Host,
		// 	"client":         c.ClientIP(),
		// 	"proto":          c.Request.Proto,
		// 	"request_method": c.Request.Method,
		// 	"useTime":        responseTime.String(), // 毫秒
		// 	"content_type":   c.Request.Header.Get("Content-Type"),
		// 	"status":         c.Writer.Status(),
		// 	"user_agent":     c.Request.UserAgent(),
		// }
		content := fmt.Sprintf("[%s] [%s] [%s] [%s] [%v] %s", responseTime, c.ClientIP(), c.Request.Proto, c.Request.Method, c.Writer.Status(), c.Request.URL.Path)
		log.Trace(content)
	}()
	c.Next()
}
