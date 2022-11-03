package env

import (
	"os"
	"path/filepath"
)

type Environment struct {
	AssemblyDir  string
	AssemblyFile string
	Mode         ApplicationRunMode
}

var (
	// 程序集所在目录
	AssemblyDir string
	// 程序集全文件名
	AssemblyFile string
)

func init() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file, _ := filepath.Abs(os.Args[0])
	if ex, err := os.Executable(); err == nil {
		file, _ = filepath.Abs(ex)
		dir = filepath.Dir(file)
	}
	AssemblyDir = dir
	AssemblyFile = file
}
