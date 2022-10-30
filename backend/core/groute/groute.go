package groute

import (
	"backend/core/echo"
	"backend/core/log"

	"github.com/gin-gonic/gin"
)

type Router struct {
	// relative Path.
	Url      string
	Use      []gin.HandlerFunc
	Handle   func(r gin.IRoutes)
	Children Routers
}
type Routers = []Router

func Use(m ...gin.HandlerFunc) []gin.HandlerFunc {
	return m
}

func Register(root gin.IRouter, routers Routers) {

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Tracef("[Register Router] [%v] [%v] [%v] [%v]", echo.Yellow(httpMethod), echo.Green(absolutePath), handlerName, nuHandlers)
	}

	for _, router := range routers {
		node := root.Group(router.Url)
		if len(router.Use) > 0 {
			node.Use(router.Use...)
		}
		if router.Handle != nil {
			router.Handle(node)
		}
		if len(router.Children) > 0 {
			Register(node, router.Children)
		}
	}
}
