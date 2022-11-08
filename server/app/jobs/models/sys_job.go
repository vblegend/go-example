package models

import (
	"server/common/models"
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

// OnQueryAfter 数据从数据表读取之后触发
func (job *SysJob) OnQueryAfter() {
	// job.JobName = "xxoo"
}

// OnInsertBefore 数据插入数据表之前触发
func (job *SysJob) OnInsertBefore() error {
	return nil
}

// OnInsertAfter 数据插入数据表之后触发
func (job *SysJob) OnInsertAfter() {
}

// OnUpdateBefore 数据更新至数据表之前触发
func (job *SysJob) OnUpdateBefore() error {
	return nil
}

// OnUpdateAfter 数据更新至数据表之后触发
func (job *SysJob) OnUpdateAfter() {

}

// OnDeleteBefore 数据从数据表删除之前触发
func (job *SysJob) OnDeleteBefore() error {
	return nil
}

// OnDeleteAfter 数据从数据表删除之后触发
func (job *SysJob) OnDeleteAfter() {

}
