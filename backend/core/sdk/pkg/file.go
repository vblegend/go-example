package pkg

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func FileExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

const (
	Sha256Key = "liu.yandong.hanks"
)

// 计算文件的sha256特征码  extendKey为可选密钥 不需要则填nil
func FileSha256(file string, extendKey []byte) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return BinarySha256(content, extendKey)
}

// 计算已打开的文件sha256特征码  extendKey为可选密钥 不需要则填nil
func IOFileSha256(file fs.File, extendKey []byte) (string, error) {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return BinarySha256(content, extendKey)
}

// 计算数据的sha256特征码  extendKey为可选密钥 不需要则填nil
func BinarySha256(content []byte, extendKey []byte) (string, error) {
	h := sha256.New()
	h.Write([]byte(Sha256Key))
	if extendKey != nil {
		h.Write(extendKey)
	}
	h.Write(content)
	sha := h.Sum(nil)
	return hex.EncodeToString(sha), nil
}

func FilesCheck(files []string) error {
	for _, file := range files {
		if !FileExist(file) {
			return errors.New(file)
		}
	}
	return nil
}

func AbsPathCheck(paths []string) error {
	for _, path := range paths {
		str := path[:1]
		if str != "/" {
			return errors.New(path)
		}
	}
	return nil
}

func CheckEmptyInArray(values []string) bool {
	for _, value := range values {
		if value == "" {
			return true
		}
	}
	return false
}

// 校验目录列表中目录是否存在  若不存在则返回异常
func DirsCheck(dirs []string) error {
	for _, dir := range dirs {
		if !PathExist(dir) {
			return errors.New(dir)
		}
	}
	return nil
}

func FileCreate(content bytes.Buffer, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		log.Println(err)
	}
	file.Close()
}

type ReplaceHelper struct {
	Root    string //路径
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWrok() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		log.Println("DIR:", path)
		return nil
	}
	//文件类型需要进行过滤
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	log.Printf("h.OldText: %s \n", h.OldText)
	log.Printf("h.NewText: %s \n", h.NewText)

	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}

func FileMonitoringById(ctx context.Context, filePth string, id string, group string, hookfn func(context.Context, string, string, string)) {

	// initMaxLine := 100;
	f, err := os.Open(filePth)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	f.Seek(0, 2)
	for {
		select {
		case <-ctx.Done(): //取出值即说明是结束信号
			// fmt.Println("收到信号，父context的协程退出,time=", time.Now().Unix())
			return
		default:
			line, err := rd.ReadString('\n')
			// 如果是文件末尾不返回
			if len(line) > 0 {
				hookfn(ctx, id, group, line)
			}
			if err == io.EOF {
				time.Sleep(500 * time.Millisecond)
				continue
			} else if err != nil {
				log.Println(err)
				break
			}
		}
		// if ctx.Err() != nil {
		// 	break
		// }
	}
}

// 获取文件大小
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

//获取当前路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
