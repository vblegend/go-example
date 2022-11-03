package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"server/sugar/log"
	"server/sugar/restful"
	"strings"

	"github.com/gin-gonic/gin"
)

//
func writeError(c *gin.Context, err interface{}) {
	if c.IsAborted() {
		return // c.Status(200)

	}
	switch errStr := err.(type) {
	case restful.AssertInterrupter:
		restful.Error(c, errStr.Code, errStr.Error)
	default:
		callStack := string(debug.Stack())
		lines := strings.Split(callStack, "\n")
		content := fmt.Sprintf("%s\n%s", err, strings.Join(lines[7:], "\n"))
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": content})
		log.Error(content)
	}
}

// CustomError 自定义错误处理中间件
func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			writeError(c, err)
		}
	}()
	c.Next()
}
