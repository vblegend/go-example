package service

import (
	"backend/app/admin/models"

	"errors"

	"backend/core/log"
	"backend/core/service"

	"gorm.io/gorm"
)

type SysUser struct {
	service.Service
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *models.SysUser) error {
	var err error
	var model models.SysUser
	db := e.Orm.First(&model, c.UserId)
	if err = db.Error; err != nil {
		log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
	if err != nil {
		log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	return nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) ResetPwd(id int, newPassword string) error {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = e.Orm.Model(c).Select("UserId", "Password", "Salt").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("无权更新该数据")
		}
		log.Errorf("db error: %s", err)
		return err
	}
	// var ok bool
	// ok, err = pkg.CompareHashAndPassword(c.Password, oldPassword)
	// if err != nil {
	// 	e.Log.Errorf("CompareHashAndPassword error, %s", err.Error())
	// 	return err
	// }
	// if !ok {
	// 	err = errors.New("incorrect Password")
	// 	e.Log.Warnf("user[%d] %s", id, err.Error())
	// 	return err
	// }
	c.Password = newPassword
	db := e.Orm.Model(c).Where("user_id = ?", id).Select("Password", "Salt").Updates(c)
	if err = db.Error; err != nil {
		log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("set password error")
		log.Warnf("db update error")
		return err
	}
	return nil
}
