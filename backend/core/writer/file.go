package writer

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// timeFormat 时间格式
// 用于文件名称格式
const timeFormat = "2006-01-02"

// FileWriter 文件写入结构体
type FileWriter struct {
	file *os.File
	num  uint
	opts Options
}

// NewFileWriter 实例化FileWriter, 支持大文件分割
func NewFileWriter(opts ...Option) (*FileWriter, error) {
	p := &FileWriter{
		opts: setDefault(),
	}
	for _, o := range opts {
		o(&p.opts)
	}
	err := p.open()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *FileWriter) open() error {
	var filename string
	var err error
	if p.file != nil {
		p.file.Close()
	}

	for {
		filename = p.getFilename()
		info, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				break
			}
			return err
		}
		if p.opts.cap == 0 {
			break
		}
		if info.Size() < int64(p.opts.cap) {
			break
		}
		p.num++
	}
	p.file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		return err
	}
	return nil

}

func (p *FileWriter) rmEchoColor(data []byte) []byte {
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

func (p *FileWriter) checkFile() {
	info, _ := p.file.Stat()
	if p.opts.cap > 0 && uint(info.Size()) > p.opts.cap {
		// 生成新文件
		if uint(info.Size()) > p.opts.cap {
			p.num++
		} else {
			p.num = 0
		}
		filename := p.getFilename()
		_ = p.file.Close()
		p.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	}
}

// Write 写入方法
func (p *FileWriter) Write(data []byte) (n int, err error) {
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

func (p *FileWriter) getFilename() string {
	if p.opts.cap == 0 {
		return filepath.Join(p.opts.path, fmt.Sprintf("%s.%s", p.opts.fileName, p.opts.suffix))
	}
	return filepath.Join(p.opts.path, fmt.Sprintf("%s.%d.%s", p.opts.fileName, p.num, p.opts.suffix))
}
