package middleware

import (
	"backend/common/database"
	"backend/core/sdk"

	"github.com/gin-gonic/gin"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", sdk.Runtime.GetDb(database.Default))
	c.Next()
}
