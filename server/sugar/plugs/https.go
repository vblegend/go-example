package plugs

import (
	"fmt"
	"server/sugar/log"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func NewHttpsHandler(https bool, domain string, prot uint) gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: true,
		SSLHost:     fmt.Sprintf("%s:%d", domain, prot),
	})
	return func(c *gin.Context) {
		if https {
			err := secureMiddleware.Process(c.Writer, c.Request)
			if err != nil {
				log.Error(err)
				return
			}
		}
		c.Next()
	}
}
