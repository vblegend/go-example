package plugs

import (
	"backend/core/sdk"

	"github.com/gin-gonic/gin"
)

func WithContextDB(dbname string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", sdk.Runtime.GetDb(dbname).WithContext(c))
		c.Next()
	}
}
