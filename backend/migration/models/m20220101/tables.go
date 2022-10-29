package m20220101

import (
	"backend/migration/models/common"
)

type SysJob struct {
	JobId          int    `json:"jobId" gorm:"primaryKey;autoIncrement"` // 编码
	JobName        string `json:"jobName" gorm:"size:255;"`              // 名称
	JobGroup       string `json:"jobGroup" gorm:"size:255;"`             // 任务分组
	JobType        int    `json:"jobType" gorm:"size:1;"`                // 任务类型
	CronExpression string `json:"cronExpression" gorm:"size:255;"`       // cron表达式
	InvokeTarget   string `json:"invokeTarget" gorm:"size:255;"`         // 调用目标
	Args           string `json:"args" gorm:"size:255;"`                 // 目标参数
	MisfirePolicy  int    `json:"misfirePolicy" gorm:"size:255;"`        // 执行策略
	Concurrent     int    `json:"concurrent" gorm:"size:1;"`             // 是否并发
	Enabled        bool   `json:"enabled" gorm:"size:1;"`                // 状态开关
	common.ModelTime
}

func (SysJob) TableName() string {
	return "sys_job"
}

type SysMenu struct {
	MenuId     int    `json:"menuId" gorm:"primaryKey;autoIncrement"`
	MenuName   string `json:"menuName" gorm:"size:128;"`
	Title      string `json:"title" gorm:"size:128;"`
	Icon       string `json:"icon" gorm:"size:128;"`
	Path       string `json:"path" gorm:"size:128;"`
	Paths      string `json:"paths" gorm:"size:128;"`
	MenuType   string `json:"menuType" gorm:"size:1;"`
	Action     string `json:"action" gorm:"size:16;"`
	ParentId   int    `json:"parentId" gorm:"size:11;"`
	NoCache    bool   `json:"noCache" gorm:"size:8;"`
	Breadcrumb string `json:"breadcrumb" gorm:"size:255;"`
	Component  string `json:"component" gorm:"size:255;"`
	Sort       int    `json:"sort" gorm:"size:4;"`
	Visible    string `json:"visible" gorm:"size:1;"`
	IsFrame    string `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	common.ControlBy
	common.ModelTime
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string `json:"username" gorm:"type:varchar(64);comment:用户名"`
	Password string `json:"-" gorm:"type:varchar(128);comment:密码"`
	NickName string `json:"nickName" gorm:"type:varchar(128);comment:昵称"`
	Phone    string `json:"phone" gorm:"type:varchar(11);comment:手机号"`
	Salt     string `json:"-" gorm:"type:varchar(255);comment:加盐"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:头像"`
	Sex      string `json:"sex" gorm:"type:varchar(255);comment:性别"`
	Email    string `json:"email" gorm:"type:varchar(128);comment:邮箱"`
	Remark   string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	Status   string `json:"status" gorm:"type:varchar(4);comment:状态"`
	common.ControlBy
	common.ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}
