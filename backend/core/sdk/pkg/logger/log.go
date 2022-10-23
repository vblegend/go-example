package logger

import (
	"io"
	"os"

	"backend/core/debug/writer"
	"backend/core/logger"
	"backend/core/sdk/pkg"

	log "backend/core/logger"
)

// SetupLogger 日志 cap 单位为kb
func SetupLogger(opts ...Option) logger.Logger {
	op := setDefault()
	for _, o := range opts {
		o(&op)
	}
	if !pkg.PathExist(op.path) {
		err := pkg.PathCreate(op.path)
		if err != nil {
			log.Fatalf("create dir error: %s", err.Error())
		}
	}
	var err error
	var output io.Writer
	switch op.stdout {
	case "file":
		output, err = writer.NewFileWriter(
			writer.WithPath(op.path),
			writer.WithCap(op.cap<<10),
		)
		if err != nil {
			log.Fatal("logger setup error: %s", err.Error())
		}
	default:
		output = os.Stdout
	}
	var level logger.Level
	level, err = logger.GetLevel(op.level)
	if err != nil {
		log.Fatalf("get logger level error, %s", err.Error())
	}
	lvl, err := logger.GetLevel(os.Getenv("LOG_LEVEL"))
	if err == nil {
		level = lvl
	}
	log.DefaultLogger = logger.NewLogger(logger.WithLevel(level), logger.WithOutput(output), logger.WithLocation(op.location))
	return log.DefaultLogger
}
