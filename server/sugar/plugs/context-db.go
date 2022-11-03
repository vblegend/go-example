package plugs

import (
	"server/sugar/sdk"

	"github.com/gin-gonic/gin"
)

func WithContextDB(dbname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", sdk.Runtime.GetDb(dbname).WithContext(c))
		c.Next()
	}
}
