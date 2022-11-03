package plugs

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WWWAuthenticator(user string, passwd string) gin.HandlerFunc {
	//性能调优监视 TODO Gin自主隐藏，待优化
	authStr := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, passwd))))
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth != authStr {
			c.Header("www-Authenticate", "Basic")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
		// restful.Error(c, http.StatusBadRequest, err, "")
	}
}
