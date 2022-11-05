package groute

import (
	"server/sugar/echo"
	"server/sugar/log"

	"github.com/gin-gonic/gin"
)

type Router struct {
	// relative Path.  为空时不创建层级使用上一级路由
	Url      string
	Use      []gin.HandlerFunc
	Handle   func(r gin.IRoutes)
	Children Routers
}
type Routers = []Router

func Use(m ...gin.HandlerFunc) []gin.HandlerFunc {
	return m
}

func Register(root gin.IRouter, router Router) {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Tracef("[Register Router] [%v] [%v] [%v] [%v]", echo.Yellow(httpMethod), echo.Green(absolutePath), handlerName, nuHandlers)
	}
	if router.Url != "" {
		root = root.Group(router.Url)
	}
	if len(router.Use) > 0 {
		root.Use(router.Use...)
	}
	if router.Handle != nil {
		router.Handle(root)
	}
	for _, router := range router.Children {
		Register(root, router)
	}
}
