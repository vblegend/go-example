package restful

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthWithWWW(user string, passwd string, gcall func(c *gin.Context)) gin.HandlerFunc {
	//性能调优监视 TODO Gin自主隐藏，待优化
	if gcall == nil {
		panic("method gcall cannot be nil")
	}
	authStr := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, passwd))))
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth != authStr {
			c.Header("www-Authenticate", "Basic")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		gcall(c)
		// restful.Error(c, http.StatusBadRequest, err, "")
	}
}
