package random

import (
	"encoding/hex"
	"math/rand"
	"time"
	"unsafe"

	"golang.org/x/crypto/scrypt"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIDBits = 6
	// All 1-bits as many as letterIdBits
	letterIDMask = 1<<letterIDBits - 1
	letterIDMax  = 63 / letterIDBits
)

// String 随机生成指定长度的随机字符串
func String(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIDMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIDMax
		}
		if idx := int(cache & letterIDMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIDBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// var (
// 	chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
// )

// // 随机生成指定长度的随机字符串
// func String(length int) string {
// 	clen := len(chars)
// 	maxrb := 255 - (256 % clen)
// 	b := make([]byte, length)
// 	r := make([]byte, length+(length/4)) // storage for random bytes.
// 	i := 0
// 	for {
// 		if _, err := rand.Read(r); err != nil {
// 			panic("Error reading random bytes: " + err.Error())
// 		}
// 		for _, rb := range r {
// 			c := int(rb)
// 			if c > maxrb {
// 				continue // Skip this number to avoid modulo bias.
// 			}
// 			b[i] = chars[c%clen]
// 			i++
// 			if i == length {
// 				return string(b)
// 			}
// 		}
// 	}
// }

// SetPassword 根据明文密码和加盐值生成密码
func SetPassword(password string, salt string) (verify string, err error) {
	var rb []byte
	rb, err = scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		return
	}
	verify = hex.EncodeToString(rb)
	return
}
