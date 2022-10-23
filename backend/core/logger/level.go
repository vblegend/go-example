package logger

import (
	"fmt"
)

type Level int8

const (
	// TraceLevel 级别。 指定比调试更细粒度的信息事件。
	TraceLevel Level = iota - 2
	// DebugLevel 级别。 通常只在调试时启用。 非常详细的日志记录。
	DebugLevel
	// InfoLevel 是默认的日志记录优先级。
	// 关于应用程序内部发生的事情的一般操作条目。
	InfoLevel
	// 警告级别。 值得关注的非关键条目。
	WarnLevel
	// ErrorLevel 级别。 日志。 用于绝对应该注意的错误。
	ErrorLevel
	// 致命级别。 记录然后调用`logger.Exit(1)`。 最高级别的严重性。
	FatalLevel
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	}
	return ""
}

// LevelForGorm 转换成gorm日志级别
func (l Level) LevelForGorm() int {
	switch l {
	case FatalLevel, ErrorLevel:
		return 2
	case WarnLevel:
		return 3
	case InfoLevel, DebugLevel, TraceLevel:
		return 4
	default:
		return 1
	}
}

// Enabled returns true if the given level is at or above this level.
func (l Level) Enabled(lvl Level) bool {
	return lvl >= l
}

// GetLevel converts a level string into a logger Level value.
// returns an error if the input string does not match known values.
func GetLevel(levelStr string) (Level, error) {
	switch levelStr {
	case TraceLevel.String():
		return TraceLevel, nil
	case DebugLevel.String():
		return DebugLevel, nil
	case InfoLevel.String():
		return InfoLevel, nil
	case WarnLevel.String():
		return WarnLevel, nil
	case ErrorLevel.String():
		return ErrorLevel, nil
	case FatalLevel.String():
		return FatalLevel, nil
	}
	return InfoLevel, fmt.Errorf("Unknown Level String: '%s', defaulting to InfoLevel", levelStr)
}

func Info(args ...interface{}) {
	DefaultLogger.Log(InfoLevel, args...)
}

func Infof(template string, args ...interface{}) {
	DefaultLogger.Logf(InfoLevel, template, args...)
}

func Trace(args ...interface{}) {
	DefaultLogger.Log(TraceLevel, args...)
}

func Tracef(template string, args ...interface{}) {
	DefaultLogger.Logf(TraceLevel, template, args...)
}

func Debug(args ...interface{}) {
	DefaultLogger.Log(DebugLevel, args...)
}

func Debugf(template string, args ...interface{}) {
	DefaultLogger.Logf(DebugLevel, template, args...)
}

func Warn(args ...interface{}) {
	DefaultLogger.Log(WarnLevel, args...)
}

func Warnf(template string, args ...interface{}) {
	DefaultLogger.Logf(WarnLevel, template, args...)
}

func Error(args ...interface{}) {
	DefaultLogger.Log(ErrorLevel, args...)
}

func Errorf(template string, args ...interface{}) {
	DefaultLogger.Logf(ErrorLevel, template, args...)
}

func Fatal(args ...interface{}) {
	DefaultLogger.Log(FatalLevel, args...)
	// os.Exit(1)
}

func Fatalf(template string, args ...interface{}) {
	DefaultLogger.Logf(FatalLevel, template, args...)
	// os.Exit(1)
}

// Returns true if the given level is at or lower the current logger level
func V(lvl Level, log Logger) bool {
	l := DefaultLogger
	if log != nil {
		l = log
	}
	return l.Options().Level <= lvl
}
