package futils

import (
	"errors"
	"os"
)

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// 如果目录不存在则创建目录， 如果存在则赋予0777权限
func MkDirIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return os.Chmod(path, 0777)
	}
	return os.MkdirAll(path, os.ModeDir|os.ModePerm)
}

// PathExist 判断目录是否存在
func PathExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 校验目录列表中目录是否存在  若不存在则返回异常
func PathExists(dirs []string) error {
	for _, dir := range dirs {
		if !PathExist(dir) {
			return errors.New(dir)
		}
	}
	return nil
}

func AbsPathCheck(paths ...string) error {
	for _, path := range paths {
		str := path[:1]
		if str != "/" {
			return errors.New(path)
		}
	}
	return nil
}
