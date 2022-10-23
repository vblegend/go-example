package common

import (
	"database/sql/driver"
	"fmt"
	"time"
)

var TimeFormat = "2006-01-02 15:04:05"

type DateTime struct {
	time.Time
}

// 返回一个当前时间对象，该对象在Json中默认格式化为标准的时间格式字符串
func NowTime() DateTime {
	return DateTime{time.Now()}
}

// 返回一个格式化后的当前时间字符串
func NowString() string {
	return time.Now().Format(TimeFormat)
}

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = DateTime{Time: time.Time{}}
		return
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), loc)
	*t = DateTime{Time: now}
	return
}

// MarshalJSON on JSONTime format Time field with Y-m-d H:i:s
func (t DateTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
