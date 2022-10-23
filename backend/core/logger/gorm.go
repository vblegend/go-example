package logger

import (
	"backend/core/sdk/console"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)



type gormLogger struct {
	logger.Config
}

func (l *gormLogger) getLogger(ctx context.Context) Logger {
	requestId := ctx.Value("X-Request-Id")
	if requestId != nil {
		return DefaultLogger.Fields(map[string]interface{}{
			"x-request-id": requestId,
		})
	}
	return DefaultLogger
}

// LogMode log mode
func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (l gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		//l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		log := l.getLogger(ctx)
		// log.Logf(InfoLevel, l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		log.Logf(TraceLevel, "┏[%s]%s\t", console.Green("GORM"), utils.FileWithLineNum())
		log.Logf(TraceLevel, "┗[%s]%s\t"+msg, console.Green("GORM"), fmt.Sprintf(msg))

	}
}

// Warn print warn messages
func (l gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		//l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		log := l.getLogger(ctx)
		// log.Logf(WarnLevel, l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		log.Logf(TraceLevel, "┏[%s]%s\t", console.Yellow("GORM"), utils.FileWithLineNum())
		log.Logf(TraceLevel, "┗[%s]%s\t"+msg, console.Yellow("GORM"), fmt.Sprintf(msg))
	}
}

// Error print error messages
func (l gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		//l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		log := l.getLogger(ctx)
		// log.Logf(ErrorLevel, l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		log.Logf(TraceLevel, "┏[%s]%s\t", console.Red("GORM"), utils.FileWithLineNum())
		log.Logf(TraceLevel, "┗[%s]%s\t"+msg, console.Red("GORM"), fmt.Sprintf(msg))
	}
}

func (l gormLogger) getUseTimeColor(elapsed time.Duration) func(string) string {
	if elapsed > time.Millisecond*500 { // 很慢
		return console.Red
	} else if elapsed > time.Millisecond*100 { //慢
		return console.Yellow
	} else if elapsed > time.Millisecond*10 { // 一般
		return console.Cyan
	} else {
		return console.Green // 快
	}
}

// Trace print sql message
func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > logger.Silent {
		log := l.getLogger(ctx)
		elapsed := time.Since(begin)
		sql, rows := fc()
		timeColorFunc := l.getUseTimeColor(elapsed)
		sqlColorFunc := console.Green
		if err != nil {
			sqlColorFunc = console.Red
		}
		rowOut := console.Yellow(fmt.Sprintf("[rows:%d]", rows))
		useTimeOut := timeColorFunc(fmt.Sprintf("[%.3fms]", float64(elapsed.Nanoseconds())/1e6))
		sqlout := sqlColorFunc(sql)

		//
		log.Logf(TraceLevel, "┏%s\t%s", useTimeOut, utils.FileWithLineNum())
		log.Logf(TraceLevel, "┗%s\t%s", rowOut, sqlout)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Log(ErrorLevel, err)
		}
	}
}

type traceRecorder struct {
	logger.Interface
	BeginAt      time.Time
	SQL          string
	RowsAffected int64
	Err          error
}

func (l traceRecorder) New() *traceRecorder {
	return &traceRecorder{Interface: l.Interface, BeginAt: time.Now()}
}

func (l *traceRecorder) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	l.BeginAt = begin
	l.SQL, l.RowsAffected = fc()
	l.Err = err
}

func New(config logger.Config) logger.Interface {
	return &gormLogger{
		Config: config,
	}
}
