package restful

import (
	"net/http"
	"server/sugar/plugs"

	"github.com/gin-gonic/gin"
)

// Error 失败数据处理
func Error(c *gin.Context, code int, err error) {
	res := Response{
		Code:    code,
		TraceId: GetTraceID(c),
		Msg:     err.Error(),
		Data:    nil,
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	res := Response{
		Code:    0,
		TraceId: GetTraceID(c),
		Msg:     msg,
		Data:    data,
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res = Page{
		Count:     count,
		PageIndex: pageIndex,
		PageSize:  pageSize,
		List:      result,
	}
	OK(c, res, msg)
}

// Custum 兼容函数
func Custum(c *gin.Context, data gin.H) {
	data["traceId"] = GetTraceID(c)
	c.AbortWithStatusJSON(http.StatusOK, data)
}

// GetTraceID 从上下文获取TraceID
func GetTraceID(c *gin.Context) string {
	return c.MustGet(plugs.TraceIdKey).(string)
}
