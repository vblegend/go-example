package log

import (
	"backend/core/echo"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type gormLogger struct {
	// logger.Config
	logger Logger
}

func (l gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info print info
func (l gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Logf(InfoLevel, "┏[%s]%s\t", echo.Green("GORM"), utils.FileWithLineNum())
	l.logger.Logf(InfoLevel, "┗[%s]%s\t"+msg, echo.Green("GORM"), fmt.Sprintf(msg))
}

// Warn print warn messages
func (l gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Logf(WarnLevel, "┏[%s]%s\t", echo.Yellow("GORM"), utils.FileWithLineNum())
	l.logger.Logf(WarnLevel, "┗[%s]%s\t"+msg, echo.Yellow("GORM"), fmt.Sprintf(msg))
}

// Error print error messages
func (l gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Logf(ErrorLevel, "┏[%s]%s\t", echo.Red("GORM"), utils.FileWithLineNum())
	l.logger.Logf(ErrorLevel, "┗[%s]%s\t"+msg, echo.Red("GORM"), fmt.Sprintf(msg))
}

func (l gormLogger) getUseTimeColor(elapsed time.Duration) func(string) string {
	if elapsed > time.Millisecond*500 { // 很慢
		return echo.Red
	} else if elapsed > time.Millisecond*100 { //慢
		return echo.Yellow
	} else if elapsed > time.Millisecond*10 { // 一般
		return echo.Cyan
	} else {
		return echo.Green // 快
	}
}

// Trace print sql message
func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	// log := l.getLogger(ctx)
	elapsed := time.Since(begin)
	sql, rows := fc()
	timeColorFunc := l.getUseTimeColor(elapsed)
	sqlColorFunc := echo.Green
	if err != nil {
		sqlColorFunc = echo.Red
	}
	rowOut := echo.Yellow(fmt.Sprintf("[rows:%d]", rows))
	useTimeOut := timeColorFunc(fmt.Sprintf("[%.3fms]", float64(elapsed.Nanoseconds())/1e6))
	sqlout := sqlColorFunc(sql)
	//
	l.logger.Logf(TraceLevel, "┏%s\t%s", useTimeOut, utils.FileWithLineNum())
	l.logger.Logf(TraceLevel, "┗%s\t%s", rowOut, sqlout)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		l.logger.Log(ErrorLevel, err)
	}
}

func NewGORMLogger(logger Logger) logger.Interface {
	return &gormLogger{
		logger: logger,
	}
}
