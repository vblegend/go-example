package models

import (
	"backend/common/models"
)

type SysJobIndex struct {
	JobId int `uri:"JobId" json:"jobId" gorm:"primaryKey;autoIncrement"`
}

type SysJob struct {
	SysJobIndex
	JobName        string `json:"jobName" gorm:"size:255;"`                  // 名称
	JobGroup       string `form:"jobGroup" json:"jobGroup" gorm:"size:255;"` // 任务分组
	JobType        int    `json:"jobType" gorm:"size:1;"`                    // 任务类型
	CronExpression string `json:"cronExpression" gorm:"size:255;"`           // cron表达式
	InvokeTarget   string `json:"invokeTarget" gorm:"size:255;"`             // 调用目标
	Args           string `json:"args" gorm:"size:255;"`                     // 目标参数
	MisfirePolicy  int    `json:"misfirePolicy" gorm:"size:255;"`            // 执行策略
	Concurrent     int    `json:"concurrent" gorm:"size:1;"`                 // 是否并发
	Enabled        bool   `json:"enabled" gorm:"size:1;"`                    // 状态开关
	models.ModelTime
}

func (SysJob) TableName() string {
	return "sys_job"
}
