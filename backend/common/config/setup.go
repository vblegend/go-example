package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var ConfigYamlPath string

// Setup 载入配置文件
func Setup(configYml string, fs ...func()) {
	ConfigYamlPath = configYml
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
	if err := vtoml.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := vtoml.Unmarshal(_config); err != nil {
		panic(err)
	}
	for _, fo := range fs {
		fo()
	}
}

func Save() {
	bytes, err := yaml.Marshal(_config)
	if err == nil {
		ioutil.WriteFile(ConfigYamlPath, bytes, os.ModePerm)
	}
}
