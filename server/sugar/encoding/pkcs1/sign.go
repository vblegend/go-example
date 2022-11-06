package pkcs1

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
)

// RSASSA-PKCS1-v1_5
type dataSign struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// IDataSignature 数据签名器
type IDataSignature interface {
	// Sign 用私钥对数据签名
	Sign(data []byte) (string, error)
}

// IDataValidator 数据签名验证器
type IDataValidator interface {
	// Verify 用公钥验签
	Verify(data []byte, sign string) (err error)
}

// NewSignature 用私钥创建一个签名验证器
func NewSignature(privateKeyFile string, passwd string) (IDataSignature, error) {
	sign := dataSign{}
	err := sign.loadPrivateKey(privateKeyFile, passwd)
	if err != nil {
		return nil, err
	}
	return &sign, nil
}

// NewValidator 用公钥创建一个签名验证器
func NewValidator(publicKeyFile string) (IDataValidator, error) {
	sign := dataSign{}
	err := sign.loadPublicKey(publicKeyFile)
	if err != nil {
		return nil, err
	}
	return &sign, nil
}

// NewKey 生成密钥对
// ssh-keygen -t rsa -f private.pem -m pem
// openssl rsa -in private.pem -pubout -out public.pem
// ====================================================
// openssl genrsa -aes128 -passout pass:"123456" -out private.pem 2048
// openssl rsa -in private.pem  -out public.pem -pubout -outform PEM  -passin pass:123456
func NewKey() {
	// privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	panic(err)
	// }
	// // The public key is a part of the *rsa.PrivateKey struct
	// publicKey := privateKey.PublicKey
}

func (ds *dataSign) loadPrivateKey(file string, passwd string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(data)
	der := block.Bytes
	if x509.IsEncryptedPEMBlock(block) {
		data, err := x509.DecryptPEMBlock(block, []byte(passwd))
		if err != nil {
			return err
		}
		der = data
	}
	prk, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		return err
	}
	ds.privateKey = prk
	return nil
}

func (ds *dataSign) loadPublicKey(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(data)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	ds.publicKey = pubInterface.(*rsa.PublicKey)
	return nil
}

func (ds *dataSign) Sign(data []byte) (string, error) {
	// 对请求参数按照字母顺序进行排序并组合
	h := sha256.New()
	h.Write(data)
	Sha256Code := h.Sum(nil)
	// 使用rsa算法进行签名
	// 第一个参数是一个随机数参数器，确保每次相同输入产生不同的签名
	// 第二个参数是密钥
	// 第三个参数是我们上面使用的hash函数
	// 第四个参数是被hash函数处理过的原始输入
	signatureAfter, err := rsa.SignPKCS1v15(rand.Reader, ds.privateKey, crypto.SHA256, Sha256Code)
	if err != nil {
		return "", err
	}
	// 返回base64编码的字符串
	return base64.StdEncoding.EncodeToString(signatureAfter), nil
}

func (ds *dataSign) Verify(data []byte, sign string) (err error) {

	// 和签名步骤相同，将排序后的signature进行hash操作
	h := sha256.New()
	h.Write(data)
	Sha256Code := h.Sum(nil)
	// 对签名进行base64解码
	decodeSignature, err := base64.StdEncoding.DecodeString(sign)
	// 使用rsa验签函数
	// 第一个参数是公钥
	// 第二个参数是hash函数
	// 第三个参数是被hash函数处理过的原始输入
	// 第四个参数是被处理过的签名
	err = rsa.VerifyPKCS1v15(ds.publicKey, crypto.SHA256, Sha256Code, decodeSignature)
	if err != nil { // 验证失败
		return err
	}
	return nil // 验证成功
}
