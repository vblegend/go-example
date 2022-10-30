package initialize

import (
	"fmt"
	"time"

	"backend/common/config"
	"backend/core/echo"
	"backend/core/log"
	"backend/core/sdk"

	"backend/common/config/types"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var opens = map[string]func(string) gorm.Dialector{
	"sqlite3": sqlite.Open,
	"mysql":   mysql.Open,
}

func CleanDBConnect(key string) {
	db := sdk.Runtime.GetDb(key)
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
		CleanDBConnect(key)
		log.Info(echo.Green(fmt.Sprintf("Connecting to database %s ...", key)))
		db, err := NewDBConnection(cfg)
		if err != nil {
			log.Error(echo.Red(fmt.Sprintf("Database %s connect fail...", key)))
			continue
		}

		sdk.Runtime.SetDb(key, db)
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
	var register *dbresolver.DBResolver
	if register == nil {
		register = dbresolver.Register(dbresolver.Config{})
	}
	if c.ConnMaxIdleTime > 0 {
		register = register.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Second)
	}
	if c.ConnMaxLifeTime > 0 {
		register = register.SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTime) * time.Second)
	}
	if c.MaxOpenConns > 0 {
		register = register.SetMaxOpenConns(c.MaxOpenConns)
	}
	if c.MaxIdleConns > 0 {
		register = register.SetMaxIdleConns(c.MaxIdleConns)
	}
	if register != nil {
		err = db.Use(register)
	}
	return db, err

}