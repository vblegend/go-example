package plugs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

func Cache(cacheSecond uint32) gin.HandlerFunc {
	return func(c *gin.Context) {
		if cacheSecond > 0 {
			c.Header("Cache-Control", fmt.Sprintf("max-age=%d, public", cacheSecond))
			// c.Header("Expires", time.Now().UTC().Format(http.TimeFormat))
			// c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		}
		c.Next()
	}

}
