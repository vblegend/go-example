package log

import (
	"os"
)

var (
	loger Logger
)

func init() {
	lvl, err := GetLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		lvl = InfoLevel
	}
	loger = NewLogger(WithLevel(lvl))
}

func SetLogger(log Logger) {
	loger = log
}

func GetLogger() Logger {
	return loger
}

// 无条件必定输出 且不遵守日志格式
func Print(args ...interface{}) {
	loger.Log(PrintLevel, args...)
}

func Println(args ...interface{}) {
	loger.Log(PrintLevel, append(args, '\n')...)

}

// 无条件必定输出 且不遵守日志格式
func Printf(template string, args ...interface{}) {
	loger.Logf(PrintLevel, template, args...)
}

func Info(args ...interface{}) {
	loger.Log(InfoLevel, args...)
}

func Infof(template string, args ...interface{}) {
	loger.Logf(InfoLevel, template, args...)
}

func Trace(args ...interface{}) {
	loger.Log(TraceLevel, args...)
}

func Tracef(template string, args ...interface{}) {
	loger.Logf(TraceLevel, template, args...)
}

func Debug(args ...interface{}) {
	loger.Log(DebugLevel, args...)
}

func Debugf(template string, args ...interface{}) {
	loger.Logf(DebugLevel, template, args...)
}

func Warn(args ...interface{}) {
	loger.Log(WarnLevel, args...)
}

func Warnf(template string, args ...interface{}) {
	loger.Logf(WarnLevel, template, args...)
}

func Error(args ...interface{}) {
	loger.Log(ErrorLevel, args...)
}

func Errorf(template string, args ...interface{}) {
	loger.Logf(ErrorLevel, template, args...)
}

func Fatal(args ...interface{}) {
	loger.Log(FatalLevel, args...)
	// os.Exit(1)
}

func Fatalf(template string, args ...interface{}) {
	loger.Logf(FatalLevel, template, args...)
	// os.Exit(1)
}
