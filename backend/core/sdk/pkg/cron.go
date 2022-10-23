package pkg

import "github.com/robfig/cron/v3"

var parser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)

func CheckCron(exp string) error {
	_, err := parser.Parse(exp)
	return err
}
