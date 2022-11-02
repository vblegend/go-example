package plugs

import "github.com/gin-gonic/gin"

func Limit(count int) gin.HandlerFunc {
	sem := make(chan bool, count)
	acquire := func() { sem <- true }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}
