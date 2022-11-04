package migration

import (
	"fmt"
	"path/filepath"
	"runtime"
	"server/common/config"
	"server/common/models"
	"server/sugar/echo"
	"server/sugar/log"
	"server/sugar/state"
	"server/sugar/tools"
)

var (
	version tools.DataMigratePair = tools.DataMigratePair{}
)

func DataBaseMigrate() {
	migrations := models.Migrations{}

	log.Info(echo.Green("Start Patch Upgrade..."))
	db := state.Default.GetDB(config.DefaultDB)
	err := db.AutoMigrate(&models.Migration{})
	if err != nil {
		log.Errorf("Table Migration Migrate Fail.. %s\n", echo.Red(err.Error()))
		panic(0)
	}
	// get all Migration record
	err = db.Table("sys_migration").Find(&migrations).Error
	if err != nil {
		log.Errorf("Table Migration Migrate Fail.. %s\n", echo.Red(err.Error()))
		panic(0)
	}
	recordMigration := migrations.Map()
	versions := version.SortKeys()
	for _, key := range versions {
		if recordMigration[key] == nil {
			migrator := version[key]
			log.Info(echo.Green(fmt.Sprintf("Apply Patch %s", key)))
			err = migrator.Migrate(db)
			if err != nil {
				log.Errorf("Data Migrate Error %s\n", echo.Red(err.Error()))
				break
			}
		}
	}
	log.Info(echo.Green("End of patching ..."))
}

func GetFilename(s string) string {
	s = filepath.Base(s)
	return s[:len(s)-3]
}

func SetVersion(dm tools.IDataMigrator) {

	_, fileName, _, _ := runtime.Caller(1)
	k := GetFilename(fileName)

	dm.SetVersion(k)
	dm.SetPatcher(dm)
	if version[k] != nil {
		panic(fmt.Sprintf("检测到重复的版本迁移补丁“%s”，请修正后再使用...", k))
	}
	version.Set(k, dm)
}
