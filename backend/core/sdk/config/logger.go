package config

import (
	"backend/core/sdk/pkg/logger"
)

type Logger struct {
	Path     string // 日志路径
	Level    string // 日志等级  trace/debug/info/warn/error/fatal
	Stdout   string // 输出到哪  file/留空
	Cap      uint
	Location bool // 是否显示代码位置
}

// Setup 设置logger
func (e Logger) Setup() {
	logger.SetupLogger(
		logger.WithPath(e.Path),
		logger.WithLevel(e.Level),
		logger.WithStdout(e.Stdout),
		logger.WithCap(e.Cap),
		logger.WithLocation(e.Location),
	)
}

var LoggerConfig = new(Logger)
