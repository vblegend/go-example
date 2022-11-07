package service

import (
	"fmt"

	"gorm.io/gorm"
)

type IService interface {
	SetOrm(orm *gorm.DB)
	SetTraceID(traceID string)
	AddError(err error) error
}

type Service struct {
	Orm     *gorm.DB
	Msg     string
	MsgID   string
	Error   error
	TraceID string
}

func (db *Service) SetOrm(orm *gorm.DB) {
	db.Orm = orm
}

func (db *Service) SetTraceID(traceID string) {
	db.TraceID = traceID
}

func (db *Service) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}
