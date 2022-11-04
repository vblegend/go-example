package versions

import (
	"fmt"

	"server/migration"
	"server/sugar/tools"

	"server/migration/models/m20220101"

	"gorm.io/gorm"
)

func init() {
	migration.SetVersion(&v20220101{}) //
}

type v20220101 struct {
	tools.DataMigrator `version:"2022-01-01"`
}

// 初始化需要迁移的表对象
func (v *v20220101) Init() {
	v.Apply(m20220101.SysMenu{})
	v.Apply(m20220101.SysUser{})
	v.Apply(m20220101.SysJob{})
}

// 执行初始化脚本 或 gorm 的update/insert/delete等操作
// 升级过程不允许失败，数据脚本执行失败则会触发panic，导致此版本号升级失败。
func (v *v20220101) Patch(db *gorm.DB) {
	err := v.ExecSqlFile(db, "config/db.sql")
	if err != nil {
		fmt.Printf("%s", err)
	}
}
