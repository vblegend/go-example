package service

import (
	"backend/core/sdk/service"
)

type SysJobService struct {
	service.Service
}

// func (e *SysJobService) GetList(tx *gorm.DB, list []models.SysJob) (err error) {

// 	return e.Orm.Table(e.TableName()).Where("enabled = ?", true).Find(list).Error
// }

// func (e *SysJobService) Create(tx *gorm.DB, id models.SysJob) (err error) {
// 	return e.Orm.Table(e.TableName()).Where(id).Updates(&e).Error
// }

// func (e *SysJobService) Update(tx *gorm.DB, id models.SysJob) (err error) {
// 	return e.Orm.Table(e.TableName()).Where(id).Updates(&e).Error
// }
