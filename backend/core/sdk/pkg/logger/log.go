package logger

import (
	"io"
	"os"
	"path/filepath"

	"backend/core/log"
	"backend/core/sdk/pkg"
	"backend/core/writer"
)

// SetupLogger 日志 cap 单位为kb
func SetupLogger(opts ...Option) log.Logger {
	op := setDefault()
	for _, o := range opts {
		o(&op)
	}
	logPath, err := filepath.Abs(op.path)
	if err != nil {
		log.Fatalf("dir error: %s", err.Error())
	}
	if !pkg.PathExist(logPath) {
		err := pkg.PathCreate(logPath)
		if err != nil {
			log.Fatalf("create dir error: %s", err.Error())
		}
	}

	var output io.Writer = os.Stdout
	if op.enabled {
		fileout, err := writer.NewFileWriter(
			writer.WithPath(logPath),
			writer.WithFileName(op.fileName),
			writer.WithSuffix(op.fileSuffix),
			writer.WithCap(op.cap),
		)
		if err != nil {
			log.Fatal("logger setup error: %s", err.Error())
		} else {
			output = io.MultiWriter(fileout, os.Stdout)
		}
	}

	var level log.Level
	level, err = log.GetLevel(op.level)
	if err != nil {
		log.Fatalf("get logger level error, %s", err.Error())
	}
	lvl, err := log.GetLevel(os.Getenv("LOG_LEVEL"))
	if err == nil {
		level = lvl
	}
	log.SetLogger(log.NewLogger(log.WithLevel(level), log.WithOutput(output), log.WithLocation(op.location)))
	return log.GetLogger()
}
