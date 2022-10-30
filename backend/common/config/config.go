package config

import (
	"backend/common/config/types"
)

// Config 配置集合
type Config struct {
	Application *types.ApplicationConfigure         `yaml:"application"`
	Logger      *types.LoggerConfigure              `yaml:"logger"`
	Jwt         *types.JwtConfigure                 `yaml:"jwt"`
	Database    map[string]*types.DatabaseConfigure `yaml:"database"`
	Redis       *types.RedisConfigure               `yaml:"redis"`
	Queue       *types.QueueConfigure               `yaml:"queue"`
}

// Settings 兼容原先的配置结构
type Settings struct {
	Settings Config `yaml:"settings"`
}

var Application = new(types.ApplicationConfigure)
var Redis = new(types.RedisConfigure)
var Database = map[string]*types.DatabaseConfigure{}
var Jwt = new(types.JwtConfigure)
var Logger = new(types.LoggerConfigure)
var Queue = new(types.QueueConfigure)

var _config = &Settings{
	Settings: Config{
		Application: Application,
		Logger:      Logger,
		Jwt:         Jwt,
		Database:    Database,
		Redis:       Redis,
		Queue:       Queue,
	},
}
