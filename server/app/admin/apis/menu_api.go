package apis

import (
	"server/app/admin/service"
	"server/sugar/restful"

	"github.com/gin-gonic/gin"
)

type MenuApi struct {
	restful.Api
}

func (api *MenuApi) GetMenuTree(c *gin.Context) {
	s := service.MenuService{}
	api.Make(c, &s)
	api.OK(s.GetMenuTree(), "ok")
}
