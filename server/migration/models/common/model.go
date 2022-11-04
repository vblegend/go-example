package common

import (
	"time"

	"gorm.io/gorm"
)

type ModelIndex struct {
	Id int `uri:"id" json:"id" gorm:"primaryKey;autoIncrement"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
