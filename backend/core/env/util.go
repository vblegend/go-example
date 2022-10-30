package env

import "backend/core/model"

// get Current Time
func Time() model.DateTime {
	return model.NowTime()
}
