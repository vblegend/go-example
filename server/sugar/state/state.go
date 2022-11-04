package state

import (
	"sync"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type state struct {
	dbs   sync.Map
	crons sync.Map
}

func (s *state) GetDB(key DataBaseKey) *gorm.DB {
	if value, ok := s.dbs.Load(key); ok {
		return value.(*gorm.DB)
	}
	return nil
}

func (s *state) SetDB(key DataBaseKey, value *gorm.DB) {
	if value == nil {
		s.dbs.Delete(key)
	} else {
		s.dbs.Store(key, value)
	}
}

func (s *state) GetCrontab(key string) *cron.Cron {
	if value, ok := s.crons.Load(key); ok {
		return value.(*cron.Cron)
	}
	return nil
}

func (s *state) SetCrontab(key string, value *cron.Cron) *cron.Cron {
	if value, ok := s.crons.Load(key); ok {
		cron := value.(*cron.Cron)
		cron.Stop().Done()
	}
	if value == nil {
		s.crons.Delete(key)
	} else {
		s.crons.Store(key, value)
	}
	return value
}
