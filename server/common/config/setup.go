package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"server/sugar/env"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// YamlFilePath Yaml配置文件路径
var YamlFilePath string

// Setup 载入配置文件
func Setup(configYml string, fs ...func()) {
	YamlFilePath = configYml
	full, err := filepath.Abs(configYml)
	if err != nil {
		panic(errors.New(""))
	}
	dir := filepath.Dir(full)
	name := filepath.Base(full)
	vtoml := viper.New()
	vtoml.SetConfigName(name)
	vtoml.SetConfigType("yaml")
	vtoml.AddConfigPath(dir)
	// vtoml.AutomaticEnv() // 支持一下 环境变量被
	// // vtoml.SetEnvPrefix("$")
	if err := vtoml.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := vtoml.Unmarshal(_config); err != nil {
		panic(err)
	}
	env.SetMode(env.ApplicationRunMode(_config.Settings.Application.Mode))
	for _, fo := range fs {
		fo()
	}
}

// Save 保存配置文件
func Save() {
	bytes, err := yaml.Marshal(_config)
	if err == nil {
		ioutil.WriteFile(YamlFilePath, bytes, os.ModePerm)
	}
}
