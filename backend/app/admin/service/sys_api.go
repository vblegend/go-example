package service

import (
	"errors"
	"fmt"

	"backend/core/sdk/runtime"
	"backend/core/sdk/service"

	"gorm.io/gorm"

	"backend/app/admin/models"
	"backend/app/admin/service/dto"
	cDto "backend/common/dto"
)

type SysApi struct {
	service.Service
}

// GetPage 获取SysApi列表
func (e *SysApi) GetPage(c *dto.SysApiGetPageReq, list *[]models.SysApi, count *int64) error {
	var err error
	var data models.SysApi

	err = e.Orm.Debug().Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysApiPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysApi对象with id
func (e *SysApi) Get(d *dto.SysApiGetReq, model *models.SysApi) *SysApi {
	var data models.SysApi
	err := e.Orm.Model(&data).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysApi error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	return e
}

// Update 修改SysApi对象
func (e *SysApi) Update(c *dto.SysApiUpdateReq) error {
	var model = models.SysApi{}
	db := e.Orm.Debug().First(&model, c.GetId())
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	c.Generate(&model)
	db = e.Orm.Save(&model)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysApi error:%s", err)
		return err
	}

	return nil
}

// Remove 删除SysApi
func (e *SysApi) Remove(d *dto.SysApiDeleteReq) error {
	var data models.SysApi

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysApi error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// CheckStorageSysApi 创建SysApi对象
func (e *SysApi) CheckStorageSysApi(c *[]runtime.Router) error {
	for _, v := range *c {
		err := e.Orm.Debug().Where(models.SysApi{Path: v.RelativePath, Action: v.HttpMethod}).
			Attrs(models.SysApi{Handle: v.Handler}).
			FirstOrCreate(&models.SysApi{}).Error
		if err != nil {
			err := fmt.Errorf("Service CheckStorageSysApi error: %s \r\n ", err.Error())
			return err
		}
	}
	return nil
}
