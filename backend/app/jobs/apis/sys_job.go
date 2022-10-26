package apis

import (
	"net/http"

	"backend/core/api"
	"backend/core/log"

	"github.com/gin-gonic/gin"

	"backend/app/jobs/service"
	"backend/common/dto"
)

type SysJob struct {
	api.Api
}

// RemoveJobForService 调用service实现
func (e SysJob) RemoveJobForService(c *gin.Context) {
	v := dto.GeneralDelDto{}
	s := service.SysJob{}
	err := e.Make(c, &s.Service).Bind(&v).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	err = s.RemoveJob(&v)
	if err != nil {
		e.Logger.Errorf("RemoveJob error, %s", err.Error())
		e.Error(500, err, "")
		return
	}
	e.OK(nil, s.Msg)
}

// StartJobForService 启动job service实现
func (e SysJob) StartJobForService(c *gin.Context) {
	s := service.SysJob{}
	var v dto.GeneralGetDto
	err := e.Make(c, &s.Service).BindUris(&v)
	if err != nil {
		log.Warnf("error: %s", err)
		e.Error(http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	// .Bind(&v)
	err = s.StartJob(&v)
	if err != nil {
		log.Errorf("GetCrontabKey error, %s", err.Error())
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, s.Msg)
}
