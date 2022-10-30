package plugs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TraceGenerator func() string

const TraceIdKey = "{[trace-id]}"

// 接收来自客户端header中的 requestIdKey， 如果没有则调用 fn()生成一个key
// 当接口返回时 返回该key字段至 traceId
func TraceId(requestIdKey string, fn TraceGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		id := c.GetHeader(requestIdKey)
		if id == "" {
			id = fn()
			c.Request.Header.Set(TraceIdKey, id)
		}
		c.Set(TraceIdKey, id)
		c.Next()
	}
}
