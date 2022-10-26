package config

import (
	"backend/core/sdk/pkg/logger"
)

type Logger struct {
	Path       string // 日志路径
	Level      string // 日志等级  trace/debug/info/warn/error/fatal
	Enabled    bool   // 是否输出到文件 false 仅打印至console
	Cap        uint
	Location   bool // 是否显示代码位置
	FileName   string
	FileSuffix string
}

// Setup 设置logger
func (e Logger) Setup() {
	logger.SetupLogger(
		logger.WithPath(e.Path),
		logger.WithLevel(e.Level),
		logger.WithEnabled(e.Enabled),
		logger.WithCap(e.Cap),
		logger.WithLocation(e.Location),
		logger.WithFileName(e.FileName),
		logger.WithFileSuffix(e.FileSuffix),
	)
}

var LoggerConfig = new(Logger)
