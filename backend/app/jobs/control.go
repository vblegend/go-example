package jobs

import (
	"backend/app/jobs/models"
	"backend/app/jobs/socket"
	"backend/core/log"
	"backend/core/sdk"
	"backend/core/utils"
	"backend/core/ws"
	"sync"

	"github.com/robfig/cron/v3"
)

var tasks = make(map[int]*JobCore)
var mux sync.RWMutex
var crontab *cron.Cron
var WebSocketChannel *socket.JobSocketChannel

// 初始化
func Setup() {
	log.Info("JobCore Starting...")
	db := sdk.Runtime.GetDb("default")
	crontab = sdk.Runtime.GetCrontab("*")
	if crontab != nil {
		crontab.Stop().Done()
		crontab = nil
	}
	crontab = utils.NewCronEngine()
	sdk.Runtime.SetCrontab("*", crontab)
	sysJob := models.SysJob{}
	jobList := make([]models.SysJob, 0)
	err := db.Table(sysJob.TableName()).Find(&jobList).Error
	if err != nil {
		log.Errorf("JobCore init error\n%v", err)
	}

	for i := 0; i < len(jobList); i++ {
		if jobList[i].Enabled {
			err := StartJob(jobList[i])
			if err != nil {
				log.Errorf("job %s fails to be started, because %v", jobList[i].JobName, err)
			}
		}
	}
	// 初始化 Job websocket

	WebSocketChannel = &socket.JobSocketChannel{}
	ws.Default.RegisterChannel(&ws.WSChannel{}, ws.Auth_PostAndSendNeedJoin)
	// 其中任务
	crontab.Start()
	log.Info("JobCore start success.")
}

func ConfigJob(model models.SysJob) error {
	StopJob(model.JobId)
	if model.Enabled {
		return StartJob(model)
	}
	return nil
}

func StartJob(model models.SysJob) error {
	mux.Lock()
	defer mux.Unlock()
	task := tasks[model.JobId]
	if task != nil {
		return nil
	}
	j := &JobCore{}
	j.InvokeTarget = model.InvokeTarget
	j.CronExpression = model.CronExpression
	j.JobId = model.JobId
	j.Name = model.JobName
	j.Args = model.Args
	id, err := crontab.AddJob(j.CronExpression, j)
	if err != nil {
		return err
	}
	tasks[model.JobId] = j
	j.EntryId = int(id)
	return nil
}

func StopJob(jobId int) error {
	mux.Lock()
	defer mux.Unlock()
	task := tasks[jobId]
	if task == nil {
		return nil
	}
	crontab.Remove(cron.EntryID(task.EntryId))
	delete(tasks, jobId)
	return nil
}
