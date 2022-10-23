package database

import (
	"fmt"
	"time"

	log "backend/core/logger"
	"backend/core/sdk"
	"backend/core/sdk/config"
	toolsConfig "backend/core/sdk/config"
	"backend/core/sdk/console"
	"backend/core/sdk/pkg"
	toolsDB "backend/core/tools/database"

	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
)

const (
	// 进程性能数据
	SnapDB_Performance = "snaps.performance"
)

const (
	SQLite = "*"
	MySQL  = "mysql"
)

var opens = map[string]func(string) gorm.Dialector{
	"sqlite3": sqlite.Open,
	"mysql":   mysql.Open,
}

func CleanDBConnect(key string) {
	db := sdk.Runtime.GetDbByKey(key)
	if db != nil {
		d, e := db.DB()
		if e == nil {
			d.Close()
		}
	}
}

// Setup 配置数据库
func InitSQLiteDB() {
	CleanDBConnect("*")
	db, err := NewDBConnection(toolsConfig.DatabaseConfig)
	if err != nil {
		log.Error(console.Red("SQLite connect fail..."))
		return
	}
	sdk.Runtime.SetDb(SQLite, db)
	log.Info(console.Green("SQLite connect sucess..."))
}

func Development() {
	db := sdk.Runtime.GetDbByKey(SQLite)
	if db != nil {
		visible := 0
		if config.ApplicationConfig.Mode == pkg.Production {
			visible = 1
		} // 开发模式下显示菜单配置页面
		db.Exec("UPDATE sys_menu SET visible = ? WHERE menu_id = 51", visible)
	}
}

// 初始化mysql连接
func InitMysqlDB() {
	fmtstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/siteweb?charset=utf8&parseTime=True&loc=Local&timeout=1000ms", config.MysqlConfig.User, config.MysqlConfig.Password, config.MysqlConfig.Host, config.MysqlConfig.Port)
	mysqlcfg := config.Database{Driver: "mysql", Source: fmtstr}
	db, err := NewDBConnection(&mysqlcfg)
	if err != nil {
		log.Error(console.Red("Mysql connect fail..."))
		return
	}
	CleanDBConnect(MySQL)
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	sdk.Runtime.SetDb(MySQL, db)
	log.Info(console.Green("Mysql connect sucess..."))
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

func NewDBConnection(c *toolsConfig.Database) (*gorm.DB, error) {
	registers := make([]toolsDB.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := toolsDB.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	return resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: log.New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.LevelForGorm()),
			},
		),
	}, opens[c.Driver])
}
