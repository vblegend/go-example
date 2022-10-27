package locks

import (
	"errors"
	"os"
	"syscall"
)

func NewFileLocker(filepath string) (*FileLocker, error) {
	flock, err := os.Open(filepath)
	if err != nil {
		flock, err = os.Create(filepath)
	}
	fl := FileLocker{
		flock: flock,
	}
	return &fl, err
}

type FileLocker struct {
	flock *os.File
}

func (f *FileLocker) Lock() error {
	err := syscall.Flock(int(f.flock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return errors.New("Process is only allowed to run once")
	}
	return nil
}

func (f *FileLocker) UnLock() {
	if f.flock != nil {
		syscall.Flock(int(f.flock.Fd()), syscall.LOCK_UN)
		f.flock = nil
	}
}

func (f *FileLocker) IsLocked() bool {
	err := syscall.Flock(int(f.flock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return true
	}
	syscall.Flock(int(f.flock.Fd()), syscall.LOCK_UN)
	return false
}
