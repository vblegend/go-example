package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func NewFileLogger(file string) (*log.Logger, error) {
	paths, _ := filepath.Split(file)
	err := MkDirIfNotExist(paths)
	if err != nil {
		return nil, err
	}
	handle, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Fail to OpenFile :%s", err.Error()))
	}
	var logger = log.New(handle, "", log.LstdFlags)
	return logger, nil
}
