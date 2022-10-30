package assembly

import (
	"backend/core/locks"
	"backend/core/log"
	"errors"
	"fmt"
)

var locker locks.IRunOfOnecLocker

func init() {
	var err error
	lockFile := fmt.Sprintf("/tmp/%s.lock", AppFileName)
	pidFile := fmt.Sprintf("/tmp/%s.pid", AppFileName)
	locker, err = locks.NewRunOfOnecLocker(lockFile, pidFile)
	if err != nil {
		log.Errorf("init run of onec fail %v", err)
	}
}
func RunOfOnec() error {
	if locker == nil {
		return errors.New("")
	}
	return locker.RunOfOnec()
}

func IsRuning(pidOut *int) bool {
	if locker == nil {
		return false
	}
	return locker.IsRuning(pidOut)
}
