package migration

import (
	"backend/common/database"
	"backend/common/models"
	"backend/core/logger"
	"backend/core/sdk"
	"backend/core/sdk/console"
	"fmt"
	"path/filepath"
	"sort"
	"sync"

	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	mutex   sync.Mutex
	version map[string]database.IDataMigrator = make(map[string]database.IDataMigrator)
)

func DataBaseMigrate() {
	logger.Info(console.Green("Start Patch Upgrade..."))
	db = sdk.Runtime.GetDbByKey(database.SQLite)
	err := db.Debug().AutoMigrate(&models.Migration{})
	if err != nil {
		return
	}
	// 兼容旧版本的 版本号 1599190683659 重置为 2022-01-01
	db.Model(models.Migration{}).Where("version = ?", "1599190683659").Updates(models.Migration{Version: "2022-01-01"})
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
			logger.Errorf("Init Error %s\n", console.Red(err.Error()))
			break
		}
		if count > 0 {
			continue
		}
		logger.Info(console.Green(fmt.Sprintf("Apply Patch %s", v)))
		err = migrator.Migrate(db)
		if err != nil {
			logger.Errorf("Data Migrate Error %s\n", console.Red(err.Error()))
			break
		}
	}
	logger.Info(console.Green("End of patching ..."))
}

func GetFilename(s string) string {
	s = filepath.Base(s)
	return s[:len(s)-3]
}

func SetVersion(k string, dm database.IDataMigrator) {
	dm.SetVersion(k)
	dm.SetPatcher(dm)
	mutex.Lock()
	defer mutex.Unlock()
	if version[k] != nil {
		panic(fmt.Sprintf("检测到重复的版本迁移补丁“%s”，请修正后再使用...", k))
	}
	version[k] = dm
}
