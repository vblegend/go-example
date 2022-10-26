package service

import (
	"backend/core/log"
	"fmt"

	"gorm.io/gorm"
)

type Service struct {
	Orm       *gorm.DB
	Msg       string
	MsgID     string
	Log       *log.Helper
	Error     error
	RequestId string
}

func (db *Service) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}
