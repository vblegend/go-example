package groute

import (
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
