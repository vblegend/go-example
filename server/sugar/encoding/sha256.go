package encoding

import (
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"io/ioutil"
	"os"
)

// 计算文件的sha256特征码  extendKey为可选密钥 不需要则填nil
func Sha256Path(file string, extendKey []byte) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return Sha256Binary(content, extendKey)
}

// 计算已打开的文件sha256特征码  extendKey为可选密钥 不需要则填nil
func Sha256File(file fs.File, extendKey []byte) (string, error) {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return Sha256Binary(content, extendKey)
}

func Sha256String(content string, extendKey []byte) (string, error) {
	return Sha256Binary([]byte(content), extendKey)
}

// 计算数据的sha256特征码  extendKey为可选密钥 不需要则填nil
func Sha256Binary(content []byte, extendKey []byte) (string, error) {
	h := sha256.New()
	if extendKey != nil {
		h.Write(extendKey)
	}
	h.Write(content)
	sha := h.Sum(nil)
	return hex.EncodeToString(sha), nil
}
