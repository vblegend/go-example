package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"server/sugar/echo"
	"strings"
	"time"
)

type defaultLogger struct {
	opts Options
}

// Init (opts...) should only overwrite provided options
func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}
	return nil
}

func (l *defaultLogger) String() string {
	return "default"
}

// logCallerfilePath returns a package/file:line description of the caller,
// preserving only the leaf directory name and file name.
func logCallerfilePath(loggingFilePath string) string {
	// To make sure we trim the path correctly on Windows too, we
	// counter-intuitively need to use '/' and *not* os.PathSeparator here,
	// because the path given originates from Go stdlib, specifically
	// runtime.Caller() which (as of Mar/17) returns forward slashes even on
	// Windows.
	//
	// See https://github.com/golang/go/issues/3335
	// and https://github.com/golang/go/issues/18151
	//
	// for discussion on the issue on Go side.
	idx := strings.LastIndexByte(loggingFilePath, '/')
	if idx == -1 {
		return loggingFilePath
	}
	idx = strings.LastIndexByte(loggingFilePath[:idx], '/')
	if idx == -1 {
		return loggingFilePath
	}
	return loggingFilePath[idx+1:]
}

func (l *defaultLogger) Log(level Level, v ...interface{}) {
	l.logf(level, "", v...)
}

func (l *defaultLogger) Logf(level Level, format string, v ...interface{}) {
	l.logf(level, format, v...)
}

func (l *defaultLogger) logf(level Level, format string, v ...interface{}) {
	// TODO decide does we need to write message if log level not used?
	if !l.opts.Level.Enabled(level) {
		return
	}
	var message string
	if format == "" {
		message = fmt.Sprint(v...)
	} else {
		message = fmt.Sprintf(format, v...)
	}
	if level == PrintLevel {
		_, err := l.opts.Out.Write(append([]byte(message), '\n'))
		if err != nil {
			log.Printf("log [Logf] write error: %s \n", err.Error())
		}
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	colord := echo.Green
	if WarnLevel.Enabled(level) {
		colord = echo.Yellow
	}
	if ErrorLevel.Enabled(level) {
		colord = echo.Red
	}
	if l.opts.Location {
		atFile := ""
		if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
			atFile = fmt.Sprintf("\t@%s:%d", logCallerfilePath(file), line)
		}
		message = message + atFile
	}
	logStr := fmt.Sprintf("%s [%s] %v\n", now, colord(strings.ToUpper(level.String())), message)
	_, err := l.opts.Out.Write([]byte(logStr))
	if err != nil {
		log.Printf("log [Logf] write error: %s \n", err.Error())
	}
}

func (l *defaultLogger) Options() Options {
	// not guard against options Context values
	opts := l.opts
	return opts
}

// NewLogger builds a new logger based on options
func NewLogger(opts ...Option) Logger {
	// Default options
	options := Options{
		Level:           InfoLevel,
		Out:             os.Stderr,
		CallerSkipCount: 3,
	}
	l := &defaultLogger{opts: options}
	if err := l.Init(opts...); err != nil {
		l.Log(FatalLevel, err)
	}
	return l
}
