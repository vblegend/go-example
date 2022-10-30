package initialize

import (
	"backend/common/config"
	"backend/core/futils"
	"backend/core/log"
	"backend/core/writer"
	"io"
	"os"
	"path/filepath"
)

func InitLogger() {
	logPath, err := filepath.Abs(config.Logger.Path)
	if err != nil {
		log.Fatalf("dir error: %s", err.Error())
	}
	if !futils.PathExist(logPath) {
		err := futils.PathCreate(logPath)
		if err != nil {
			log.Fatalf("create dir error: %s", err.Error())
		}
	}

	var output io.Writer = os.Stdout
	if config.Logger.Enabled {
		fileout, err := writer.NewFileWriter(
			writer.WithPath(logPath),
			writer.WithFileName(config.Logger.FileName),
			writer.WithSuffix(config.Logger.FileSuffix),
			writer.WithCap(config.Logger.Cap),
		)
		if err != nil {
			log.Fatal("logger setup error: %s", err.Error())
		} else {
			output = io.MultiWriter(fileout, os.Stdout)
		}
	}

	var level log.Level
	level, err = log.GetLevel(config.Logger.Level)
	if err != nil {
		log.Fatalf("get logger level error, %s", err.Error())
	}
	lvl, err := log.GetLevel(os.Getenv("LOG_LEVEL"))
	if err == nil {
		level = lvl
	}
	log.SetLogger(log.NewLogger(log.WithLevel(level), log.WithOutput(output), log.WithLocation(config.Logger.Location)))
}
