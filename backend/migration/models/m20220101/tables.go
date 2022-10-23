package m20220101

import (
	syscommon "backend/common"
	"backend/migration/models/common"
	"time"
)

//sys_role_dept
type SysRoleDept struct {
	RoleId int `gorm:"size:11;primaryKey"`
	DeptId int `gorm:"size:11;primaryKey"`
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}

type SysApi struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Handle string `json:"handle" gorm:"size:128;comment:handle"`
	Title  string `json:"title" gorm:"size:128;comment:标题"`
	Path   string `json:"path" gorm:"size:128;comment:地址"`
	Type   string `json:"type" gorm:"size:16;comment:接口类型"`
	Action string `json:"action" gorm:"size:16;comment:请求类型"`
	common.ModelTime
	common.ControlBy
}

func (SysApi) TableName() string {
	return "sys_api"
}

type SysColumns struct {
	ColumnId           int    `gorm:"primaryKey;autoIncrement" json:"columnId"`
	TableId            int    `gorm:"" json:"tableId"`
	ColumnName         string `gorm:"size:128;" json:"columnName"`
	ColumnComment      string `gorm:"column:column_comment;size:128;" json:"columnComment"`
	ColumnType         string `gorm:"column:column_type;size:128;" json:"columnType"`
	GoType             string `gorm:"column:go_type;size:128;" json:"goType"`
	GoField            string `gorm:"column:go_field;size:128;" json:"goField"`
	JsonField          string `gorm:"column:json_field;size:128;" json:"jsonField"`
	IsPk               string `gorm:"column:is_pk;size:4;" json:"isPk"`
	IsIncrement        string `gorm:"column:is_increment;size:4;" json:"isIncrement"`
	IsRequired         string `gorm:"column:is_required;size:4;" json:"isRequired"`
	IsInsert           string `gorm:"column:is_insert;size:4;" json:"isInsert"`
	IsEdit             string `gorm:"column:is_edit;size:4;" json:"isEdit"`
	IsList             string `gorm:"column:is_list;size:4;" json:"isList"`
	IsQuery            string `gorm:"column:is_query;size:4;" json:"isQuery"`
	QueryType          string `gorm:"column:query_type;size:128;" json:"queryType"`
	HtmlType           string `gorm:"column:html_type;size:128;" json:"htmlType"`
	DictType           string `gorm:"column:dict_type;size:128;" json:"dictType"`
	Sort               int    `gorm:"column:sort;" json:"sort"`
	List               string `gorm:"column:list;size:1;" json:"list"`
	Pk                 bool   `gorm:"column:pk;size:1;" json:"pk"`
	Required           bool   `gorm:"column:required;size:1;" json:"required"`
	SuperColumn        bool   `gorm:"column:super_column;size:1;" json:"superColumn"`
	UsableColumn       bool   `gorm:"column:usable_column;size:1;" json:"usableColumn"`
	Increment          bool   `gorm:"column:increment;size:1;" json:"increment"`
	Insert             bool   `gorm:"column:insert;size:1;" json:"insert"`
	Edit               bool   `gorm:"column:edit;size:1;" json:"edit"`
	Query              bool   `gorm:"column:query;size:1;" json:"query"`
	Remark             string `gorm:"column:remark;size:255;" json:"remark"`
	FkTableName        string `gorm:"" json:"fkTableName"`
	FkTableNameClass   string `gorm:"" json:"fkTableNameClass"`
	FkTableNamePackage string `gorm:"" json:"fkTableNamePackage"`
	FkLabelId          string `gorm:"" json:"fkLabelId"`
	FkLabelName        string `gorm:"size:255;" json:"fkLabelName"`
	common.ModelTime
	common.ControlBy
}

func (SysColumns) TableName() string {
	return "sys_columns"
}

type SysConfig struct {
	common.Model
	ConfigName  string `json:"configName" gorm:"type:varchar(128);comment:ConfigName"`
	ConfigKey   string `json:"configKey" gorm:"type:varchar(128);comment:ConfigKey"`
	ConfigValue string `json:"configValue" gorm:"type:varchar(255);comment:ConfigValue"`
	ConfigType  string `json:"configType" gorm:"type:varchar(64);comment:ConfigType"`
	IsFrontend  int    `json:"isFrontend" gorm:"type:varchar(64);comment:是否前台"`
	Remark      string `json:"remark" gorm:"type:varchar(128);comment:Remark"`
	common.ControlBy
	common.ModelTime
}

func (SysConfig) TableName() string {
	return "sys_config"
}

type SysBackupRecord struct {
	RecordId  int                `json:"recordId" gorm:"primaryKey;autoIncrement"` // 记录ID
	BackupId  int                `json:"backupId" gorm:"size:11;"`                 // 备份任务ID
	TimeStamp syscommon.DateTime `json:"timeStamp" gorm:"comment:备份时间"`            // 备份时间
	FilePath  string             `json:"filePath" gorm:"size:512;"`                // 文件名
	FileSize  int64              `json:"fileSize" gorm:"size:16;"`                 // 文件大小
	UseTime   float64            `json:"useTime" gorm:"size:16;"`                  // 文件大小
	Status    int                `json:"status" gorm:"size:1;"`                    // 备份状态  0：正在备份  1：备份成功  -1：备份失败
	HasErrors int                `json:"hasErrors" gorm:"size:8;"`                 // 记录中存在错误。
}

func (SysBackupRecord) TableName() string {
	return "sys_backuprecords"
}

// 实体层模型
type SysDataBackup struct {
	BackupId       int    `json:"backupId" gorm:"primaryKey;autoIncrement"`                // 编码
	BackupName     string `json:"backupName" gorm:"type:varchar(128);comment:备份名称"`        // 名称
	CronExpression string `json:"cronExpression" gorm:"type:varchar(128);comment:Crom表达式"` // cron表达式
	IsEnabled      int    `json:"isEnabled" gorm:"type:int;comment:是否开启备份"`                // 状态
	BackupFiles    string `json:"backupFiles" gorm:"type:varchar(4096);comment:备份名称"`      // 备份文件
	BackupDirs     string `json:"backupDirs" gorm:"type:varchar(4096);comment:备份名称"`       // 备份目录
	BackupTables   string `json:"backupTables" gorm:"type:varchar(4096);comment:备份名称"`     // 备份表
	OutDir         string `json:"outDir" gorm:"type:varchar(512);comment:备份名称"`            // 备份输出目录
	common.ControlBy
	common.ModelTime
}

func (SysDataBackup) TableName() string {
	return "sys_backup"
}

type DictData struct {
	DictCode  int    `gorm:"primaryKey;autoIncrement;" json:"dictCode" example:"1"` //字典编码
	DictSort  int    `gorm:"" json:"dictSort"`                                      //显示顺序
	DictLabel string `gorm:"size:128;" json:"dictLabel"`                            //数据标签
	DictValue string `gorm:"size:255;" json:"dictValue"`                            //数据键值
	DictType  string `gorm:"size:64;" json:"dictType"`                              //字典类型
	CssClass  string `gorm:"size:128;" json:"cssClass"`                             //
	ListClass string `gorm:"size:128;" json:"listClass"`                            //
	IsDefault string `gorm:"size:8;" json:"isDefault"`                              //
	Status    int    `gorm:"size:4;" json:"status"`                                 //状态
	Default   string `gorm:"size:8;" json:"default"`                                //
	Remark    string `gorm:"size:255;" json:"remark"`                               //备注
	common.ControlBy
	common.ModelTime
}

func (DictData) TableName() string {
	return "sys_dict_data"
}

type DictType struct {
	DictId   int    `gorm:"primaryKey;autoIncrement;" json:"dictId"`
	DictName string `gorm:"size:128;" json:"dictName"` //字典名称
	DictType string `gorm:"size:128;" json:"dictType"` //字典类型
	Status   int    `gorm:"size:4;" json:"status"`     //状态
	Remark   string `gorm:"size:255;" json:"remark"`   //备注
	common.ControlBy
	common.ModelTime
}

func (DictType) TableName() string {
	return "sys_dict_type"
}

type Host struct {
	HostId        int       `json:"hostId" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	HostAlias     string    `json:"hostAlias" gorm:"type:varchar(128);comment:主机别名"`
	HostIpAddress string    `json:"hostIpAddress" gorm:"type:varchar(128);comment:主机IP地址"`
	HostVersion   string    `json:"hostVersion" gorm:"type:varchar(128);comment:版本"`
	LastDateTime  time.Time `json:"lastDateTime" gorm:"comment:最后一次运行时间"`
	common.ControlBy
	common.ModelTime
}

func (Host) TableName() string {
	return "sys_host"
}

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
	Status         int    `json:"status" gorm:"size:1;"`                 // 状态
	EntryId        int    `json:"entry_id" gorm:"size:11;"`              // job启动时返回的id
	common.ModelTime
	common.ControlBy
}

func (SysJob) TableName() string {
	return "sys_job"
}

type SysLoginLog struct {
	common.Model
	Username      string    `json:"username" gorm:"type:varchar(128);comment:用户名"`
	Status        string    `json:"status" gorm:"type:varchar(4);comment:状态"`
	Ipaddr        string    `json:"ipaddr" gorm:"type:varchar(255);comment:ip地址"`
	LoginLocation string    `json:"loginLocation" gorm:"type:varchar(255);comment:归属地"`
	Browser       string    `json:"browser" gorm:"type:varchar(255);comment:浏览器"`
	Os            string    `json:"os" gorm:"type:varchar(255);comment:系统"`
	Platform      string    `json:"platform" gorm:"type:varchar(255);comment:固件"`
	LoginTime     time.Time `json:"loginTime" gorm:"type:timestamp;comment:登录时间"`
	Remark        string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	Msg           string    `json:"msg" gorm:"type:varchar(255);comment:信息"`
	CreatedAt     time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	common.ControlBy
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}

type SysMenu struct {
	MenuId     int      `json:"menuId" gorm:"primaryKey;autoIncrement"`
	MenuName   string   `json:"menuName" gorm:"size:128;"`
	Title      string   `json:"title" gorm:"size:128;"`
	Icon       string   `json:"icon" gorm:"size:128;"`
	Path       string   `json:"path" gorm:"size:128;"`
	Paths      string   `json:"paths" gorm:"size:128;"`
	MenuType   string   `json:"menuType" gorm:"size:1;"`
	Action     string   `json:"action" gorm:"size:16;"`
	Permission string   `json:"permission" gorm:"size:255;"`
	ParentId   int      `json:"parentId" gorm:"size:11;"`
	NoCache    bool     `json:"noCache" gorm:"size:8;"`
	Breadcrumb string   `json:"breadcrumb" gorm:"size:255;"`
	Component  string   `json:"component" gorm:"size:255;"`
	Sort       int      `json:"sort" gorm:"size:4;"`
	Visible    string   `json:"visible" gorm:"size:1;"`
	IsFrame    string   `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	SysApi     []SysApi `json:"sysApi" gorm:"many2many:sys_menu_api_rule"`
	common.ControlBy
	common.ModelTime
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

type SysRole struct {
	RoleId    int       `json:"roleId" gorm:"primaryKey;autoIncrement"` // 角色编码
	RoleName  string    `json:"roleName" gorm:"size:128;"`              // 角色名称
	Status    string    `json:"status" gorm:"size:4;"`                  //
	RoleKey   string    `json:"roleKey" gorm:"size:128;"`               //角色代码
	RoleSort  int       `json:"roleSort" gorm:""`                       //角色排序
	Flag      string    `json:"flag" gorm:"size:128;"`                  //
	Remark    string    `json:"remark" gorm:"size:255;"`                //备注
	Admin     bool      `json:"admin" gorm:"size:4;"`
	DataScope string    `json:"dataScope" gorm:"size:128;"`
	SysMenu   []SysMenu `json:"sysMenu" gorm:"many2many:sys_role_menu;foreignKey:RoleId;joinForeignKey:role_id;references:MenuId;joinReferences:menu_id;"`
	common.ControlBy
	common.ModelTime
}

func (SysRole) TableName() string {
	return "sys_role"
}

type SoftWareTemplate struct {
	SoftwareTemplateId    int    `json:"softwareTemplateId" gorm:"primaryKey;autoIncrement;comment:服务模板ID"`
	SoftwareName          string `json:"softwareName" gorm:"type:varchar(255);comment:服务名称"`
	SoftwareAlias         string `json:"softwareAlias" gorm:"type:varchar(128);comment:服务别名"`
	HostId                int    `json:"hostId" gorm:"type:int;comment:主机ID"`
	SoftwareType          int    `json:"softwareType" gorm:"type:int;comment:服务类型"`
	SoftwarePath          string `json:"softwarePath" gorm:"type:varchar(255);comment:服务启动路径"`
	SoftwareArgs          string `json:"softwareArgs" gorm:"type:varchar(255);comment:服务启动参数"`
	SoftwareVersion       string `json:"softwareVersion" gorm:"type:varchar(128);comment:服务版本"`
	WatchEnabled          bool   `json:"watchEnabled" gorm:"type:boolean;comment:是否开启服务看门狗功能"`
	WatchType             int    `json:"watchType" gorm:"type:int;comment:状态查询接口类型(http-1 pgrep-2(需要配置software_path))"`
	WatchSource           string `json:"watchSource" gorm:"type:varchar(255);comment:状态查询源"`
	WatchFailNumber       int    `json:"watchFailNumber" gorm:"type:int;comment:状态次数判定"`
	AutoRestartEnabled    bool   `json:"autoRestartEnabled" gorm:"type:boolean;comment:是否开启定时重启功能"`
	AutoRestartExpression string `json:"autoRestartExpression" gorm:"type:varchar(128);comment:定时重启cron表达式"`
	LogDir                string `json:"logDir" gorm:"type:varchar(255);comment:指定日志"`
	common.ControlBy
	common.ModelTime
}

func (SoftWareTemplate) TableName() string {
	return "sys_softwaretemplate"
}

type SoftWare struct {
	SoftwareId            int       `json:"softwareId" gorm:"primaryKey;autoIncrement;comment:服务ID"`
	SoftwareName          string    `json:"softwareName" gorm:"type:varchar(255);comment:服务名称"`
	SoftwareAlias         string    `json:"softwareAlias" gorm:"type:varchar(128);comment:服务别名"`
	HostId                int       `json:"hostId" gorm:"type:int;comment:主机ID"`
	SoftwareType          int       `json:"softwareType" gorm:"type:int;comment:服务类型"`
	SoftwarePath          string    `json:"softwarePath" gorm:"type:varchar(255);comment:服务启动路径"`
	SoftwareArgs          string    `json:"softwareArgs" gorm:"type:varchar(255);comment:服务启动参数"`
	SoftwareVersion       string    `json:"softwareVersion" gorm:"type:varchar(128);comment:服务版本"`
	LastDateTime          time.Time `json:"lastDateTime" gorm:"comment:最后一次运行时间"`
	WatchEnabled          bool      `json:"watchEnabled" gorm:"type:boolean;comment:是否开启服务看门狗功能"`
	WatchType             int       `json:"watchType" gorm:"type:int;comment:状态查询接口类型"`
	WatchSource           string    `json:"watchSource" gorm:"type:varchar(255);comment:状态查询源"`
	WatchFailNumber       int       `json:"watchFailNumber" gorm:"type:int;comment:状态次数判定"`
	AutoRestartEnabled    bool      `json:"autoRestartEnabled" gorm:"type:boolean;comment:是否开启定时重启功能"`
	AutoRestartExpression string    `json:"autoRestartExpression" gorm:"type:varchar(128);comment:定时重启cron表达式"`
	LogDir                string    `json:"logDir" gorm:"type:varchar(255);comment:指定日志"`
	common.ControlBy
	common.ModelTime
}

func (SoftWare) TableName() string {
	return "sys_software"
}

type SysTables struct {
	TableId             int    `gorm:"primaryKey;autoIncrement" json:"tableId"`        //表编码
	TBName              string `gorm:"column:table_name;size:255;" json:"tableName"`   //表名称
	TableComment        string `gorm:"size:255;" json:"tableComment"`                  //表备注
	ClassName           string `gorm:"size:255;" json:"className"`                     //类名
	TplCategory         string `gorm:"size:255;" json:"tplCategory"`                   //
	PackageName         string `gorm:"size:255;" json:"packageName"`                   //包名
	ModuleName          string `gorm:"size:255;" json:"moduleName"`                    //go文件名
	ModuleFrontName     string `gorm:"size:255;comment:前端文件名;" json:"moduleFrontName"` //前端文件名
	BusinessName        string `gorm:"size:255;" json:"businessName"`                  //
	FunctionName        string `gorm:"size:255;" json:"functionName"`                  //功能名称
	FunctionAuthor      string `gorm:"size:255;" json:"functionAuthor"`                //功能作者
	PkColumn            string `gorm:"size:255;" json:"pkColumn"`
	PkGoField           string `gorm:"size:255;" json:"pkGoField"`
	PkJsonField         string `gorm:"size:255;" json:"pkJsonField"`
	Options             string `gorm:"size:255;" json:"options"`
	TreeCode            string `gorm:"size:255;" json:"treeCode"`
	TreeParentCode      string `gorm:"size:255;" json:"treeParentCode"`
	TreeName            string `gorm:"size:255;" json:"treeName"`
	Tree                bool   `gorm:"size:1;default:0;" json:"tree"`
	Crud                bool   `gorm:"size:1;default:1;" json:"crud"`
	Remark              string `gorm:"size:255;" json:"remark"`
	IsDataScope         int    `gorm:"size:1;" json:"isDataScope"`
	IsActions           int    `gorm:"size:1;" json:"isActions"`
	IsAuth              int    `gorm:"size:1;" json:"isAuth"`
	IsLogicalDelete     string `gorm:"size:1;" json:"isLogicalDelete"`
	LogicalDelete       bool   `gorm:"size:1;" json:"logicalDelete"`
	LogicalDeleteColumn string `gorm:"size:128;" json:"logicalDeleteColumn"`
	common.ModelTime
	common.ControlBy
}

func (SysTables) TableName() string {
	return "sys_tables"
}

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string `json:"username" gorm:"type:varchar(64);comment:用户名"`
	Password string `json:"-" gorm:"type:varchar(128);comment:密码"`
	NickName string `json:"nickName" gorm:"type:varchar(128);comment:昵称"`
	Phone    string `json:"phone" gorm:"type:varchar(11);comment:手机号"`
	RoleId   int    `json:"roleId" gorm:"type:bigint;comment:角色ID"`
	Salt     string `json:"-" gorm:"type:varchar(255);comment:加盐"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:头像"`
	Sex      string `json:"sex" gorm:"type:varchar(255);comment:性别"`
	Email    string `json:"email" gorm:"type:varchar(128);comment:邮箱"`
	DeptId   int    `json:"deptId" gorm:"type:bigint;comment:部门"`
	PostId   int    `json:"postId" gorm:"type:bigint;comment:岗位"`
	Remark   string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	Status   string `json:"status" gorm:"type:varchar(4);comment:状态"`
	common.ControlBy
	common.ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}
