package database

import (
	"fmt"
	"time"

	"backend/core/console"
	"backend/core/log"
	"backend/core/sdk"
	"backend/core/sdk/config"
	"backend/core/sdk/pkg"

	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
)

const (
	// 进程性能数据
	SnapDB_Performance = "snaps.performance"
)

const (
	Default = "Default"
	Standby = "Standby"
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
func InitDatabase() {
	configmap := config.DatabaseConfig
	for key, cfg := range configmap {
		CleanDBConnect(key)
		db, err := NewDBConnection(cfg)
		if err != nil {
			log.Error(console.Red(fmt.Sprintf("Database %s connect fail...", key)))
			return
		}
		sdk.Runtime.SetDb(Default, db)
		log.Info(console.Green(fmt.Sprintf("Database %s connect sucess...", key)))
	}
}

func Development() {
	db := sdk.Runtime.GetDb(Default)
	if db != nil {
		visible := 0
		if config.ApplicationConfig.Mode == pkg.Production {
			visible = 1
		} // 开发模式下显示菜单配置页面
		db.Exec("UPDATE sys_menu SET visible = ? WHERE menu_id = 51", visible)
	}
}

// 初始化redis连接
func InitRedisDB() {
	client := sdk.Runtime.GetRedisClient()
	if client != nil {
		client.Close()
	}
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisConfig.Host, config.RedisConfig.Port),
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Error(console.Red("Redis connect fail..."))
		return
	}
	sdk.Runtime.SetRedisClient(client)
	log.Info(console.Green("Redis connect sucess..."))
}

func NewDBConnection(c *config.Database) (*gorm.DB, error) {
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
