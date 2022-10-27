package jobs

import (
	"backend/core/log"
	"reflect"
	"time"
)

var timeFormat = "2006-01-02 15:04:05"

type JobCore struct {
	InvokeTarget   string
	Name           string
	JobId          int
	EntryId        int
	CronExpression string
	Args           string
}

func (e *JobCore) Run() {
	startTime := time.Now()
	typed := GetClassTypeFromClassName(e.InvokeTarget)
	if typed == nil {
		log.Warn("Invalid job object name：" + e.InvokeTarget)
		return
	}
	obj := reflect.New(typed).Interface()
	if obj == nil {
		log.Warn("unknown error, object type did not generate any instances：" + e.InvokeTarget)
		return
	}
	job := obj.(JobsExec)
	err := job.Exec(e.Args)
	if err != nil {
		log.Errorf("mission failed!\n%v", err)
	}
	endTime := time.Now()
	latencyTime := endTime.Sub(startTime)
	log.Infof("job %s exec success , spend :%v", e.Name, latencyTime)
	return
}
