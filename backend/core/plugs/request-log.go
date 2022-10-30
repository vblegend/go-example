package plugs

import (
	"backend/core/log"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func formatDuration(d time.Duration) string {
	var v float64
	var u string
	if d > time.Hour {
		v = float64(d) / float64(time.Hour)
		u = "h"
	} else if d > time.Minute {
		v = float64(d) / float64(time.Minute)
		u = "m"
	} else if d > time.Second {
		v = float64(d) / float64(time.Second)
		u = "s"
	} else if d > time.Millisecond {
		v = float64(d) / float64(time.Millisecond)
		u = "ms"
	} else if d > time.Microsecond {
		v = float64(d) / float64(time.Microsecond)
		u = "us"
	} else if d > time.Nanosecond {
		v = float64(d) / float64(time.Nanosecond)
		u = "ns"
	}
	return fmt.Sprintf("%.2f%s", v, u)
}

func RequestLogOut(loger log.Logger, level log.Level) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		defer func() {
			endTime := time.Now()
			responseTime := endTime.Sub(startTime)
			outtime := fmt.Sprintf("% 8s", formatDuration(responseTime))
			loger.Logf(level, "[%s] [%s] [%s] [%s] [%v] %s", outtime, c.ClientIP(), c.Request.Proto, c.Request.Method, c.Writer.Status(), c.Request.URL.Path)
		}()
		c.Next()
	}
}
