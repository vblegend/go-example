package futils

import (
	"errors"
	"os"
)

func FileExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

func MFileExist(files []string) error {
	for _, file := range files {
		if !FileExist(file) {
			return errors.New(file)
		}
	}
	return nil
}
