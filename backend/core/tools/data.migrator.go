package tools

import (
	"backend/common/models"
	"backend/core/echo"
	"backend/core/log"
	"fmt"
	"io/ioutil"
	"strings"

	"gorm.io/gorm"
)

type IDataPather interface {
	// 迁移前置，正式迁移之前预处理
	// 初始化需要迁移的表对象
	Init()
	// 执行初始化脚本 或 gorm 的update/insert/delete等操作
	// 升级过程不允许失败，失败处理请跳过继续下一步
	Patch(db *gorm.DB)
}

type IDataMigrator interface {
	Migrate(db *gorm.DB) error
	SetVersion(string)
	SetPatcher(dst IDataPather)
	IDataPather
}

type DataMigrator struct {
	Version string
	tables  []interface{}
	patch   IDataPather
}
type PatchException struct {
	script   string
	excption error
}

func (dm *DataMigrator) SetPatcher(dst IDataPather) {
	dm.patch = dst
}

func (dm *DataMigrator) SetVersion(dst string) {
	dm.Version = dst
}

func (dm *DataMigrator) Apply(dst ...interface{}) {
	for i := 0; i < len(dst); i++ {
		dm.tables = append(dm.tables, dst[i])
	}
}

func (dm *DataMigrator) Migrate(db *gorm.DB) error {
	defer func() {
		err := recover()
		if err != nil {
			switch errStr := err.(type) {
			case PatchException:
				log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
				log.Error(echo.Yellow(errStr.script))
				log.Error(echo.Green("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓"))
				log.Error(echo.Red(errStr.excption.Error()))
				log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
			case error:
				log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
				log.Error(echo.Red(errStr.Error()))
				log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
			default:
				log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
				log.Error(echo.Red(fmt.Sprintf("%v", errStr)))
				log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
			}
			log.Errorf("Update failed, data rolled back...")
		}
	}()

	dm.tables = make([]interface{}, 0)
	dm.patch.Init()
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Migrator().AutoMigrate(dm.tables...)
		if err != nil {
			return err
		}
		dm.patch.Patch(tx)
		return tx.Create(&models.Migration{Version: dm.Version}).Error
	})
}

// 执行一段MYSQL脚本，如果执行失败则会返回error
func (dm *DataMigrator) ExecSql(db *gorm.DB, script string) error {
	if err := db.Exec(script).Error; err != nil {
		log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
		log.Error(echo.Yellow(script))
		log.Error(echo.Green("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓"))
		log.Error(echo.Red(err.Error()))
		log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
		if !strings.Contains(err.Error(), "Query was empty") {
			return err
		}
	}
	return nil
}

// 执行一段MYSQL脚本，如果执行失败，则会抛出panic 中止此次补丁更新事物
func (dm *DataMigrator) ExecScript(db *gorm.DB, script string) error {
	if err := db.Exec(script).Error; err != nil {
		panic(PatchException{script: script, excption: err})
	}
	return nil
}

func (dm *DataMigrator) ExecSqlFile(db *gorm.DB, filePath string) error {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("数据库基础数据初始化脚本读取失败！原因:", err.Error())
		return err
	}
	sqlList := strings.Split(string(contents), "\n")
	for i := 0; i < len(sqlList)-1; i++ {
		sql := strings.TrimSpace(sqlList[i])
		if len(sql) == 0 || strings.HasPrefix(sqlList[i], "--") {
			continue
		}
		if err = db.Exec(sql).Error; err != nil {
			log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
			log.Error(echo.Yellow(sql))
			log.Error(echo.Green("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓"))
			log.Error(echo.Red(err.Error()))
			log.Error(echo.Green("=================================================== PATCH-ERROR ==================================================="))
			if !strings.Contains(err.Error(), "Query was empty") {
				return err
			}
		}
	}
	return nil
}
