package models

import (
	"fmt"
	"server/common/models"

	"gorm.io/gorm"
)

// SysJobIndex 系统任务ID索引模型
type SysJobIndex struct {
	JobID int `uri:"JobId" json:"jobId" gorm:"primaryKey;autoIncrement"`
}

// TableName 模型对应表明
func (SysJobIndex) TableName() string {
	return "job"
}

// SysJob 系统任务模型
type SysJob struct {
	SysJobIndex
	JobName        string `json:"jobName" gorm:"size:255;" binding:"required"`                 // 名称
	JobGroup       string `form:"jobGroup" json:"jobGroup" gorm:"size:32;" binding:"required"` // 任务分组
	CronExpression string `json:"cronExpression" gorm:"size:64;" binding:"required"`           // cron表达式
	InvokeTarget   string `json:"invokeTarget" gorm:"size:64;" binding:"required"`             // 调用目标
	Args           string `json:"args" gorm:"size:4096;" binding:"required"`                   // 目标参数
	WorkDir        string `json:"workDir" gorm:"size:255;"`                                    // 工作路径
	MisfirePolicy  int    `json:"misfirePolicy" gorm:"size:255;"`                              // 执行策略
	Concurrent     int    `json:"concurrent" gorm:"size:1;"`                                   // 是否并发
	Enabled        bool   `json:"enabled" gorm:"size:1;"`                                      // 状态开关
	models.ModelTime
}

// TableName 模型对应表明
func (SysJob) TableName() string {
	return "job"
}

func (u *SysJob) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println(u)
	return
}
