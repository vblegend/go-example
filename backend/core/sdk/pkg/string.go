package pkg

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

// Top 命令 内存大小展开转换为字节
func ExpendTopMemorySize(str string) (uint64, error) {
	if str == "" {
		return 0, nil
	}
	before, _, found := strings.Cut(str, "g")
	if found {
		value, err := strconv.ParseFloat(before, 64)
		return uint64(value * 1024 * 1024 * 1024), err
	}
	before, _, found = strings.Cut(str, "m")
	if found {
		value, err := strconv.ParseFloat(before, 64)
		return uint64(value * 1024 * 1024), err
	}
	value, err := strconv.ParseUint(str, 0, 64)
	if err == nil {
		value = value * 1024
	}
	return value, err
}
