package state

import (
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

// NewState 新建一个状态
func NewState() IState {
	return &state{}
}

// Default 默认状态管理器
var Default = NewState()

// DataBaseKey 数据库 key
type DataBaseKey string

// IState 状态管理器
type IState interface {
	// GetDB 获取指定数据库的GORM对象 如果没有则返回nil
	GetDB(key DataBaseKey) *gorm.DB
	// 设置数据库Gorm对象，如果为nil则删除
	SetDB(key DataBaseKey, value *gorm.DB)

	GetCrontab(key string) *cron.Cron

	SetCrontab(key string, value *cron.Cron) *cron.Cron
}
