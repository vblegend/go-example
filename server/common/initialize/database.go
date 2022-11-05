package initialize

import (
	"fmt"

	"server/common/config"
	"server/sugar/echo"
	"server/sugar/log"
	"server/sugar/state"

	"server/common/config/types"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var opens = map[string]func(string) gorm.Dialector{
	"sqlite3": sqlite.Open,
	"mysql":   mysql.Open,
}

func CleanDBConnect(key state.DataBaseKey) {
	db := state.Default.GetDB(key)
	if db != nil {
		d, e := db.DB()
		if e == nil {
			d.Close()
		}
	}
}

// Setup 配置数据库
func InitSQLDB() {
	configmap := config.Database
	for key, cfg := range configmap {
		dbkey := state.DataBaseKey(key)
		CleanDBConnect(dbkey)
		log.Info(echo.Green(fmt.Sprintf("Connecting to database %s ...", key)))
		db, err := NewDBConnection(cfg)
		if err != nil {
			log.Error(echo.Red(fmt.Sprintf("Database %s connect fail...", key)))
			continue
		}
		state.Default.SetDB(dbkey, db)
		log.Info(echo.Green(fmt.Sprintf("Database %s connect sucess...", key)))
	}
}

func NewDBConnection(c *types.DatabaseConfigure) (*gorm.DB, error) {
	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: log.NewGORMLogger(log.GetLogger()),
	}
	db, err := gorm.Open(opens[c.Driver](c.Source), &config)
	if err != nil {
		return nil, err
	}
	if c.Driver == "mysql" {
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	return db, err

}
