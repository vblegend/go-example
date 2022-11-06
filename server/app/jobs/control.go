package jobs

import (
	"server/app/jobs/models"
	"server/app/jobs/socket"
	"server/common/config"
	"server/sugar/log"
	"server/sugar/state"
	"server/sugar/utils"
	"server/sugar/ws"
	"sync"

	"github.com/robfig/cron/v3"
)

var tasks = make(map[int]*JobCore)
var mux sync.RWMutex
var crontab *cron.Cron

// Setup 初始化
func Setup() {
	log.Info("JobCore Starting...")
	db := state.Default.GetDB(config.DefaultDB)
	crontab = state.Default.SetCrontab("*", utils.NewCronEngine())
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
	ws.Default.RegisterChannel("jobs", &socket.JobSocketChannel{}, ws.AuthAnonymous)
	// 其中任务
	crontab.Start()
	log.Info("JobCore start success.")
}

// ConfigJob ConfigJob
func ConfigJob(model models.SysJob) error {
	StopJob(model.JobID)
	if model.Enabled {
		return StartJob(model)
	}
	return nil
}

// StartJob StartJob
func StartJob(model models.SysJob) error {
	mux.Lock()
	defer mux.Unlock()
	task := tasks[model.JobID]
	if task != nil {
		return nil
	}
	j := &JobCore{}
	j.InvokeTarget = model.InvokeTarget
	j.CronExpression = model.CronExpression
	j.JobId = model.JobID
	j.Name = model.JobName
	j.Args = model.Args
	id, err := crontab.AddJob(j.CronExpression, j)
	if err != nil {
		return err
	}
	tasks[model.JobID] = j
	j.EntryId = int(id)
	return nil
}

// StopJob  StopJob
func StopJob(jobID int) error {
	mux.Lock()
	defer mux.Unlock()
	task := tasks[jobID]
	if task == nil {
		return nil
	}
	crontab.Remove(cron.EntryID(task.EntryId))
	delete(tasks, jobID)
	return nil
}
