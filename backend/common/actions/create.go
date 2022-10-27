package actions

import (
	"net/http"

	"backend/core/restful"
	"backend/core/sdk/pkg"

	"github.com/gin-gonic/gin"

	"backend/common/dto"
	"backend/common/models"
)

// CreateAction 通用新增动作
func CreateAction(control dto.Control) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := restful.GetRequestLogger(c)
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		//新增操作
		req := control.Generate()
		err = req.Bind(c)
		if err != nil {
			restful.Error(c, http.StatusUnprocessableEntity, err, err.Error())
			return
		}
		var object models.ActiveRecord
		object, err = req.GenerateM()
		if err != nil {
			restful.Error(c, 500, err, "模型生成失败")
			return
		}
		err = db.WithContext(c).Create(object).Error
		if err != nil {
			log.Errorf("Create error: %s", err)
			restful.Error(c, 500, err, "创建失败")
			return
		}
		restful.OK(c, object.GetId(), "创建成功")
		c.Next()
	}
}
