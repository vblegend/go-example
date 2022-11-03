package models

import (
	"server/common/models"
)

type SysJobIndex struct {
	JobId int `uri:"JobId" json:"jobId" gorm:"primaryKey;autoIncrement"`
}

type SysJob struct {
	SysJobIndex
	JobName        string `json:"jobName" gorm:"size:255;"`                 // 名称
	JobGroup       string `form:"jobGroup" json:"jobGroup" gorm:"size:32;"` // 任务分组
	CronExpression string `json:"cronExpression" gorm:"size:64;"`           // cron表达式
	InvokeTarget   string `json:"invokeTarget" gorm:"size:64;"`             // 调用目标
	Args           string `json:"args" gorm:"size:4096;"`                   // 目标参数
	WorkDir        string `json:"workDir" gorm:"size:255;"`                 // 工作路径
	MisfirePolicy  int    `json:"misfirePolicy" gorm:"size:255;"`           // 执行策略
	Concurrent     int    `json:"concurrent" gorm:"size:1;"`                // 是否并发
	Enabled        bool   `json:"enabled" gorm:"size:1;"`                   // 状态开关
	models.ModelTime
}

func (SysJob) TableName() string {
	return "sys_job"
}
