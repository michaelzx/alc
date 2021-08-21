package alc_gorm

import (
	"context"
	"fmt"
	"github.com/michaelzx/alc/alc_color"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

const (
	traceStr     = "\n%s \t \n" + alc_color.Reset + alc_color.LightYellow + "[%.3fms] " + alc_color.LightBlue + "[rows:%v]" + alc_color.Reset + " %s"
	traceWarnStr = alc_color.Yellow + "%s\n" + alc_color.Reset + alc_color.LightRed + "[%.3fms] " + alc_color.Yellow + "[rows:%v]" + alc_color.LightPurple + " %s" + alc_color.Reset
	traceErrStr  = alc_color.LightRed + "%s" + alc_color.Reset + "\n%s \t \n" + alc_color.Reset + alc_color.Yellow + "[%.3fms] " + alc_color.LightBlue + "[rows:%v]" + alc_color.Reset + " %s"
)

type Logger struct {
	zapLogger     *zap.Logger
	SlowThreshold time.Duration
}

func NewLogger(zapLogger *zap.Logger) *Logger {
	return &Logger{
		zapLogger:     zapLogger,
		SlowThreshold: 200 * time.Millisecond,
	}
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	panic("implement me")
}

func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.zapLogger.Sugar().Infof(msg, data...)
}

func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.zapLogger.Sugar().Warnf(msg, data...)
}

func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.zapLogger.Sugar().Errorf(msg, data...)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil:
		sql, rows := fc()
		if rows == -1 {
			l.zapLogger.Sugar().Errorf(traceErrStr, err, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.zapLogger.Sugar().Debugf(traceErrStr, err, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.zapLogger.Sugar().Warnf(traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.zapLogger.Sugar().Warnf(traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	default:
		sql, rows := fc()
		if rows == -1 {
			l.zapLogger.Sugar().Debugf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.zapLogger.Sugar().Debugf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
