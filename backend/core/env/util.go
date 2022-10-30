package env

import "backend/core/model"

func SetDefaultDateTimeFormat(format string) {
	model.DefaultDateTimeFormat = format
}

func GetDefaultDateTimeFormat() string {
	return model.DefaultDateTimeFormat
}

// get Current Time
func Time() model.DateTime {
	return model.NowTime()
}
