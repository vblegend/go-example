package jwtauth

import (
	"server/app/admin/models"
	"server/sugar/log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginModel struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
}

func (u *LoginModel) Verify(tx *gorm.DB) (MapClaims, error) {
	user := models.SysUser{}
	err := tx.Table("sys_user").Where("username = ?  and status = 2", u.Username).First(&user).Error
	if err != nil {
		log.Errorf("get user error, %s", err.Error())
		return nil, err
	}
	_, err = u.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		log.Errorf("user login error, %s", err.Error())
		return nil, err
	}
	return MapClaims{
		IdentityKey: user.UserId,
		NiceKey:     user.NickName,
	}, nil
}

func (u *LoginModel) CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}
