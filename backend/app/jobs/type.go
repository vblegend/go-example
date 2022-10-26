package jobs

import "github.com/robfig/cron/v3"

type Job interface {
	Run()
	addJob(*cron.Cron) (int, error)
}

type JobsExec interface {
	Exec(arg string) error
}

func CallExec(e JobsExec, arg string) error {
	return e.Exec(arg)
}
