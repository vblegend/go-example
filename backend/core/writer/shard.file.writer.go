package writer

import (
	"backend/core/futils"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// shardFileWriter 文件写入结构体
type shardFileWriter struct {
	file     *os.File
	num      uint
	opts     Options
	fileName string
}

// NewShardFileWriter 实例化FileWriter, 支持大文件分割
func NewShardFileWriter(opts ...Option) (io.Writer, error) {
	options := setDefault()
	for _, o := range opts {
		o(&options)
	}
	p := &shardFileWriter{
		opts:     options,
		fileName: filepath.Join(options.path, fmt.Sprintf("%s.%s", options.fileName, options.suffix)),
	}
	err := p.open()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *shardFileWriter) open() error {
	if p.file != nil {
		p.file.Close()
	}
	info, err := os.Stat(p.fileName)
	if err == nil && info.Size() > int64(p.opts.cap) {
		nextFileName, err := p.getNextFileName()
		if err == nil {
			os.Rename(p.fileName, nextFileName)
		}
	}
	p.file, err = os.OpenFile(p.fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (p *shardFileWriter) rmEchoColor(data []byte) []byte {
	re := regexp.MustCompile(`\x1b\[.[\s\S]*?\x1b\[0m`) //
	if re.Match(data) {
		val := re.ReplaceAllStringFunc(string(data), func(element string) string {
			index := strings.Index(element, "m")
			v := element[index+1 : len(element)-4]
			return v
		})
		return []byte(val)
	}
	return data
}

func (p *shardFileWriter) checkFile() {
	info, _ := p.file.Stat()
	if p.opts.cap > 0 && uint(info.Size()) > p.opts.cap {
		p.open()
	}
}

// Write 写入方法
func (p *shardFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	data2 := p.rmEchoColor(data)
	n, err = p.file.Write(data2)
	if err != nil {
		log.Printf("write file failed, %s\n", err.Error())
		return n, err
	}
	p.checkFile()
	return len(data), nil
}

func (p *shardFileWriter) getNextFileName() (string, error) {
	var i int = 0
	for {
		fileName := filepath.Join(p.opts.path, fmt.Sprintf("%s.%d.%s", p.opts.fileName, i, p.opts.suffix))
		if !futils.FileExist(fileName) {
			return fileName, nil
		}
		if i > 1024 {
			return "", nil
		}
		i++
	}
}
