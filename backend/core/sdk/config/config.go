package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"backend/core/config"
	"backend/core/config/source/file"
	"backend/core/sdk/console"
	"backend/core/sdk/pkg"

	"github.com/ghodss/yaml"
)

var (
	_cfg *Settings
)

// Settings 兼容原先的配置结构
type Settings struct {
	Settings  Config `yaml:"settings"`
	callbacks []func()
}

func (e *Settings) runCallback() {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

func (e *Settings) OnChange() {
	e.init()
	log.Println(console.Green("[INFO] Application Configure File Changed ..."))
}

func (e *Settings) Init() {
	e.init()
}

func (e *Settings) init() {
	e.Settings.Logger.Setup()
	e.runCallback()
}

// Config 配置集合
type Config struct {
	Application *Application `yaml:"application"`
	Ssl         *Ssl         `yaml:"ssl"`
	Logger      *Logger      `yaml:"logger"`
	Jwt         *Jwt         `yaml:"jwt"`
	Database    *Database    `yaml:"database"`
	Mysql       *Mysql       `yaml:"mysql"`
	Redis       *Redis       `yaml:"redis"`
	Gen         *Gen         `yaml:"gen"`
	Queue       *Queue       `yaml:"queue"`
	DataMock    *DataMock    `yaml:"dataMock"`
}

var ConfigYamlPath string

// Setup 载入配置文件
func Setup(configYml string, fs ...func()) {
	ConfigYamlPath = configYml
	fo := file.NewSource(file.WithPath(configYml))
	_cfg = &Settings{
		Settings: Config{
			Application: ApplicationConfig,
			Ssl:         SslConfig,
			Logger:      LoggerConfig,
			Jwt:         JwtConfig,
			Database:    DatabaseConfig,
			Mysql:       MysqlConfig,
			Redis:       RedisConfig,
			Gen:         GenConfig,
			Queue:       QueueConfig,
			DataMock:    DataMockConfig,
		},
		callbacks: fs,
	}
	var err error
	config.DefaultConfig, err = config.NewConfig(
		config.WithSource(fo),
		config.WithEntity(_cfg),
	)
	// 运行模式， 开发环境 或 生产环境
	// 如果命令行指定了环境变量 将覆盖配置文件中的选项
	RUN_MODE := os.Getenv("RUN_MODE")
	if len(RUN_MODE) > 0 {
		if RUN_MODE == pkg.Develop {
			_cfg.Settings.Application.Mode = pkg.Develop
		}
		if RUN_MODE == pkg.Production {
			_cfg.Settings.Application.Mode = pkg.Production
		}
	}
	if err != nil {
		log.Fatal(fmt.Sprintf("New config object fail: %s", err.Error()))
	}
	_cfg.Init()
}

func Save() {
	bytes, err := yaml.Marshal(_cfg)
	if err == nil {
		ioutil.WriteFile(ConfigYamlPath, bytes, os.ModePerm)
	}
}
