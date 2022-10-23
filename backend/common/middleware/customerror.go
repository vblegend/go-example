package middleware

import (
	"backend/core/logger"
	"backend/core/sdk/api"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

func WriteError(c *gin.Context, err interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()

	if c.IsAborted() {
		c.Status(200)
	}
	switch errStr := err.(type) {
	case api.AssertInterrupter:
		c.JSON(http.StatusOK, gin.H{
			"code": errStr.Code,
			"msg":  errStr.Messsage,
			"data": nil,
		})
	default:
		callStack := string(debug.Stack())
		lines := strings.Split(callStack, "\n")
		content := fmt.Sprintf("%s\n%s", err, strings.Join(lines[7:], "\n"))
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": content})
		logger.Error(content)
	}
}

func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			WriteError(c, err)
		}
	}()
	c.Next()
}
