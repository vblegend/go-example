package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Between 返回一个介于 min - max 之间的整型数值
func Between(min int, max int) int {
	intRange := max - min + 1
	return min + rand.Intn(intRange)
}

// Bytes 使用随机数填充 bytes数组
func Bytes(b []byte) {
	rand.Read(b)
}

// Betweenf 返回一个介于 min - max 之间的浮点数数值
func Betweenf(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
