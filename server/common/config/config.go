package config

import (
	"server/common/config/types"
	"server/sugar/state"
)

// Config 配置集合
type Config struct {
	Application *types.ApplicationConfigure         `yaml:"Application"`
	Logger      *types.LoggerConfigure              `yaml:"Logger"`
	Jwt         *types.JwtConfigure                 `yaml:"Jwt"`
	Database    map[string]*types.DatabaseConfigure `yaml:"Database"`
	Redis       *types.RedisConfigure               `yaml:"Redis"`
	Web         *types.WebConfigure                 `yaml:"Web"`
}

// Settings 兼容原先的配置结构
type Settings struct {
	Settings Config `yaml:"settings"`
}

const (
	// DefaultDB 默认数据库配置名
	DefaultDB = state.DataBaseKey("default")
	// StandbyDB 备用数据库配置名
	StandbyDB = state.DataBaseKey("standby")
)

// Application 应用配置
var Application = new(types.ApplicationConfigure)

// Redis Redis连接配置
var Redis = new(types.RedisConfigure)

// Database 多数据库配置
var Database = map[string]*types.DatabaseConfigure{}

// Jwt JWT配置
var Jwt = new(types.JwtConfigure)

// Logger 日志配置
var Logger = new(types.LoggerConfigure)

// Web web服务配置
var Web = new(types.WebConfigure)

var _config = &Settings{
	Settings: Config{
		Application: Application,
		Logger:      Logger,
		Jwt:         Jwt,
		Database:    Database,
		Redis:       Redis,
		Web:         Web,
	},
}
