package jobs

import (
	"reflect"

	"github.com/robfig/cron/v3"
)

type JobTyped struct {
	Name      string       `json:"name"`
	ClassName string       `json:"className"`
	ClassType reflect.Type `json:"-"`
}

type Job interface {
	Run()
	addJob(*cron.Cron) (int, error)
}

type JobsExec interface {
	Exec(arg string) error
}

var customJobTypedList map[string]*JobTyped = make(map[string]*JobTyped)

func RegisterClass(name string, obejct JobsExec) {
	classType := reflect.TypeOf(obejct)
	className := classType.Name()
	customJobTypedList[className] = &JobTyped{
		Name:      name,
		ClassName: className,
		ClassType: classType,
	}
}

func GetClassTypeFromClassName(className string) reflect.Type {
	typed := customJobTypedList[className]
	if typed != nil {
		return typed.ClassType
	}
	return nil
}
