package plugs

import (
	"server/sugar/state"

	"github.com/gin-gonic/gin"
)

// WithContextDB db中间件
func WithContextDB(db state.DataBaseKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := state.Default.GetDB(db)
		c.Set("db", db.WithContext(c))
		c.Next()
	}
}
