package migration

import (
	"backend/common/database"
	"backend/common/models"
	"backend/core/console"
	"backend/core/log"
	"backend/core/sdk"
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
	log.Info(console.Green("Start Patch Upgrade..."))
	db = sdk.Runtime.GetDb(database.Default)
	err := db.Debug().AutoMigrate(&models.Migration{})
	if err != nil {
		return
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
			log.Errorf("Init Error %s\n", console.Red(err.Error()))
			break
		}
		if count > 0 {
			continue
		}
		log.Info(console.Green(fmt.Sprintf("Apply Patch %s", v)))
		err = migrator.Migrate(db)
		if err != nil {
			log.Errorf("Data Migrate Error %s\n", console.Red(err.Error()))
			break
		}
	}
	log.Info(console.Green("End of patching ..."))
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
