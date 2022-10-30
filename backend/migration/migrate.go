package migration

import (
	"backend/common/models"
	"backend/core/echo"
	"backend/core/log"
	"backend/core/sdk"
	"backend/core/tools"
	"fmt"
	"path/filepath"
	"sort"
)

var (
	version map[string]tools.IDataMigrator = make(map[string]tools.IDataMigrator)
)

func DataBaseMigrate() {
	log.Info(echo.Green("Start Patch Upgrade..."))
	db := sdk.Runtime.GetDb("default")
	err := db.AutoMigrate(&models.Migration{})
	if err != nil {
		log.Errorf("Table Migration Migrate Fail.. %s\n", echo.Red(err.Error()))
		panic(0)
	}
	versions := make([]string, 0)
	for k := range version {
		versions = append(versions, k)
	}
	if !sort.StringsAreSorted(versions) {
		sort.Strings(versions)
	}
	for _, v := range versions {
		var count int64
		migrator := version[v]
		err = db.Table("sys_migration").Where("version = ?", v).Count(&count).Error
		if err != nil {
			log.Errorf("Init Error %s\n", echo.Red(err.Error()))
			break
		}
		if count > 0 {
			continue
		}
		log.Info(echo.Green(fmt.Sprintf("Apply Patch %s", v)))
		err = migrator.Migrate(db)
		if err != nil {
			log.Errorf("Data Migrate Error %s\n", echo.Red(err.Error()))
			break
		}
	}
	log.Info(echo.Green("End of patching ..."))
}

func GetFilename(s string) string {
	s = filepath.Base(s)
	return s[:len(s)-3]
}

func SetVersion(k string, dm tools.IDataMigrator) {
	dm.SetVersion(k)
	dm.SetPatcher(dm)
	if version[k] != nil {
		panic(fmt.Sprintf("检测到重复的版本迁移补丁“%s”，请修正后再使用...", k))
	}
	version[k] = dm
}
