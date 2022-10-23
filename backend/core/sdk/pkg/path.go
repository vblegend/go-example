package pkg

import (
	"os"
	"path/filepath"
)

var (
	assemblyDir  string
	assemblyFile string
)

func AssemblyDir() string {
	return assemblyDir
}
func AssemblyFile() string {
	return assemblyFile
}
func init() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	assemblyFile, _ = filepath.Abs(os.Args[0])
	assemblyDir = dir
	if ex, err := os.Executable(); err == nil {
		assemblyFile, _ = filepath.Abs(ex)
		assemblyDir = filepath.Dir(assemblyFile)
	}
}
