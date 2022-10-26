package jobs

import (
	models2 "backend/app/jobs/models"
	"backend/common/database"
	"backend/core/log"
	"backend/core/sdk"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/robfig/cron/v3"

	"backend/core/sdk/pkg/cronjob"
)

var timeFormat = "2006-01-02 15:04:05"
var customJobTypedList map[string]reflect.Type = make(map[string]reflect.Type)
var lock sync.Mutex
var crontab *cron.Cron

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
	typed := customJobTypedList[e.InvokeTarget]
	obj := reflect.New(typed).Interface()

	if obj == nil {
		log.Warn("ExecJob Run job nil")
		return
	}
	err := CallExec(obj.(JobsExec), e.Args)
	if err != nil {
		// 如果失败暂停一段时间重试
		log.Errorf("mission failed!\n%v", err)
	}
	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)
	//TODO: 待完善部分
	log.Infof("JobCore %s exec success , spend :%v", e.Name, latencyTime)
	return
}

func (h *JobCore) addJob(c *cron.Cron) (int, error) {
	id, err := c.AddJob(h.CronExpression, h)
	if err != nil {
		log.Errorf("JobCore AddJob error\n%v", err)
		return 0, err
	}
	EntryId := int(id)
	return EntryId, nil
}

func RegisterClass(obejct JobsExec) {
	typed := reflect.TypeOf(obejct)
	name := typed.Name()
	customJobTypedList[name] = typed
}

// 初始化
func Setup() {
	log.Info("JobCore Starting...")
	db := sdk.Runtime.GetDb(database.Default)
	sdk.Runtime.SetCrontab("*", cronjob.NewWithSeconds())
	crontab = sdk.Runtime.GetCrontab("*")
	sysJob := models2.SysJob{}
	jobList := make([]models2.SysJob, 0)
	err := sysJob.GetList(db, &jobList)
	if err != nil {
		log.Errorf("JobCore init error\n%v", err)
	}
	if len(jobList) == 0 {
		log.Info("JobCore total:0")
	}

	_, err = sysJob.RemoveAllEntryID(db)
	if err != nil {
		log.Errorf("JobCore remove entry_id error\n%v", err)
	}

	for i := 0; i < len(jobList); i++ {
		j := &JobCore{}
		j.InvokeTarget = jobList[i].InvokeTarget
		j.CronExpression = jobList[i].CronExpression
		j.JobId = jobList[i].JobId
		j.Name = jobList[i].JobName
		j.Args = jobList[i].Args
		sysJob.EntryId, err = AddJob(j)
		err = sysJob.Update(db, jobList[i].JobId)
	}

	// 其中任务
	crontab.Start()
	log.Info("JobCore start success.")
	// // 关闭任务
	// defer crontab.Stop()
	// select {}
}

// 添加任务 AddJob(invokeTarget string, jobId int, jobName string, cronExpression string)
func AddJob(job Job) (int, error) {
	if job == nil {
		fmt.Println("unknown")
		return 0, nil
	}
	return job.addJob(crontab)
}

// 移除任务
func Remove(entryID int) chan bool {
	ch := make(chan bool)
	go func() {
		crontab.Remove(cron.EntryID(entryID))
		log.Info("JobCore Remove success ,info entryID :", entryID)
		ch <- true
	}()
	return ch
}

// 任务停止
// func Stop() chan bool {
// 	ch := make(chan bool)
// 	go func() {
// 		global.GADMCron.Stop()
// 		ch <- true
// 	}()
// 	return ch
// }
