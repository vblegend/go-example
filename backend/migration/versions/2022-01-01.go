package versions

import (
	"fmt"
	"runtime"

	"gorm.io/gorm"

	"backend/common/database"
	"backend/migration"

	"backend/migration/models/m20220101"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.SetVersion(migration.GetFilename(fileName), &v20220101{}) //
}

type v20220101 struct {
	database.DataMigrator
}

// 初始化需要迁移的表对象
func (v *v20220101) Init() {
	v.Apply(new(m20220101.SysMenu))
	v.Apply(new(m20220101.SysUser))
	v.Apply(new(m20220101.SysJob))
}

// 执行初始化脚本 或 gorm 的update/insert/delete等操作
// 升级过程不允许失败，数据脚本执行失败则会触发panic，导致此版本号升级失败。
func (v *v20220101) Patch(db *gorm.DB) {
	err := v.ExecSqlFile(db, "config/db.sql")
	if err != nil {
		fmt.Printf("%s", err)
	}
}
