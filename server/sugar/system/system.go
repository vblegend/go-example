package system

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitQuitSignal() {
	quit := make(chan os.Signal, 5)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
}
