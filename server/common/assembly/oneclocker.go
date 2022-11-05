package assembly

import (
	"errors"
	"fmt"
	"server/sugar/locks"
	"server/sugar/log"
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

// RunOfOnec 仅运行一次，重复运行将返回异常
func RunOfOnec() error {
	if locker == nil {
		return errors.New("")
	}
	return locker.RunOfOnec()
}

// IsRuning 检查是否有程序已经调用RunOfOnec
func IsRuning(pidOut *int) bool {
	if locker == nil {
		return false
	}
	return locker.IsRuning(pidOut)
}
