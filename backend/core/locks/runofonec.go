package locks

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type IRunOfOnecLocker interface {
	RunOfOnec() error
	IsRuning(pidOut *int) bool
}

func NewRunOfOnecLocker(lockFile string, pidFile string) (IRunOfOnecLocker, error) {
	locker, err := NewFileLocker(lockFile)
	if err != nil {
		return nil, err
	}

	return &runOfOnecLocker{
		locker:  locker,
		pidFile: pidFile,
	}, nil
}

type runOfOnecLocker struct {
	locker  *FileLocker
	pidFile string
}

func (l *runOfOnecLocker) RunOfOnec() error {
	err := l.locker.Lock()
	if err == nil {
		os.Remove(l.pidFile)
		file, err := os.OpenFile(l.pidFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		file.WriteString(fmt.Sprint(os.Getpid()))
		file.Close()
	}
	return err
}

func (l *runOfOnecLocker) IsRuning(outPid *int) bool {
	locked := l.locker.IsLocked()
	if locked {
		data, err := ioutil.ReadFile(l.pidFile)
		if err == nil {
			pid, err := strconv.Atoi(string(data))
			if err == nil {
				*outPid = pid
			}
		}
		return true
	}
	return false
}
