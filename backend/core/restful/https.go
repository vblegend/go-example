package restful

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"

	"backend/core/sdk/config"
)

func NewHttpsHandler(https bool) gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: https,
		SSLHost:     fmt.Sprintf("%s:%d", config.ApplicationConfig.Domain, config.ApplicationConfig.Port),
	})
	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
