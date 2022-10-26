package models

import (
	"gorm.io/gorm"

	"backend/core/sdk/pkg"
)

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string `json:"username" gorm:"size:64;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	Salt     string `json:"-" gorm:"size:255;comment:加盐"`
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex      string `json:"sex" gorm:"size:255;comment:性别"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	Remark   string `json:"remark" gorm:"size:255;comment:备注"`
	Status   string `json:"status" gorm:"size:4;comment:状态"`
	ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (e *SysUser) AfterFind(_ *gorm.DB) error {
	return nil
}

// BaseUser 密码登录基础用户
type BaseUser struct {
	Username     string `json:"username" gorm:"type:varchar(100);comment:用户名"`
	Salt         string `json:"-" gorm:"type:varchar(255);comment:加盐;<-"`
	PasswordHash string `json:"-" gorm:"type:varchar(128);comment:密码hash;<-"`
	Password     string `json:"password" gorm:"-"`
}

// SetPassword 设置密码
func (u *BaseUser) SetPassword(value string) {
	u.Password = value
	u.generateSalt()
	u.PasswordHash = u.GetPasswordHash()
}

// GetPasswordHash 获取密码hash
func (u *BaseUser) GetPasswordHash() string {
	passwordHash, err := pkg.SetPassword(u.Password, u.Salt)
	if err != nil {
		return ""
	}
	return passwordHash
}

// generateSalt 生成加盐值
func (u *BaseUser) generateSalt() {
	u.Salt = pkg.GenerateRandomKey16()
}

// Verify 验证密码
func (u *BaseUser) Verify(db *gorm.DB, tableName string) bool {
	db.Table(tableName).Where("username = ?", u.Username).First(u)
	return u.GetPasswordHash() == u.PasswordHash
}
