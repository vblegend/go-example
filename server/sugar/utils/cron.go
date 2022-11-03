package utils

import (
	"github.com/robfig/cron/v3"
)

var parser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)

// NewWithSeconds newWithSeconds returns a Cron with the seconds field enabled.
func NewCronEngine() *cron.Cron {
	return cron.New(cron.WithParser(parser), cron.WithChain())
}

func CheckCronExpression(exp string) error {
	_, err := parser.Parse(exp)
	return err
}
