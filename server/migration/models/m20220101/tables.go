package m20220101

import (
	"server/migration/models/common"
)

type JobIndex struct {
	JobId int `uri:"JobId" json:"jobId" gorm:"primaryKey;autoIncrement"`
}
type Job struct {
	JobIndex
	JobName        string `json:"jobName" gorm:"size:255;"`                 // 名称
	JobGroup       string `form:"jobGroup" json:"jobGroup" gorm:"size:32;"` // 任务分组
	CronExpression string `json:"cronExpression" gorm:"size:64;"`           // cron表达式
	InvokeTarget   string `json:"invokeTarget" gorm:"size:64;"`             // 调用目标
	Args           string `json:"args" gorm:"size:4096;"`                   // 目标参数
	WorkDir        string `json:"workDir" gorm:"size:255;"`                 // 工作路径
	MisfirePolicy  int    `json:"misfirePolicy" gorm:"size:255;"`           // 执行策略
	Concurrent     int    `json:"concurrent" gorm:"size:1;"`                // 是否并发
	Enabled        bool   `json:"enabled" gorm:"size:1;"`                   // 状态开关
	common.ModelTime
}

func (Job) TableName() string {
	return "job"
}

type Menu struct {
	// 菜单ID
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	// 菜单标识符
	Name string `json:"name" gorm:"size:32;"`
	// 菜单类型  0 路由 1 IFrame
	Type int `json:"type" gorm:"size:2;DEFAULT:0;"`
	// 菜单标题
	Title string `json:"title" gorm:"size:64;"`
	// 菜单的图标
	Icon string `json:"icon" gorm:"size:32;"`
	// 菜单的url路径
	Path string `json:"path" gorm:"size:128;"`
	// 上级菜单ID
	ParentID int `json:"parentId" gorm:"size:11;"`
	// 菜单排序
	Sort int `json:"sort" gorm:"size:4;"`
	// 菜单是否可见
	Visible string `json:"visible" gorm:"size:1;"`
	common.ModelTime
}

func (Menu) TableName() string {
	return "menu"
}

type User struct {
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
	common.ModelTime
}

func (User) TableName() string {
	return "user"
}
