package logger

import (
	"github.com/leiqD/go-socket5/infra/conf"
	"go.uber.org/zap"
	"runtime"
	"strings"
	"sync"
)

type LoggerInterface interface {
	Debugw(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Printf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(template string, args ...interface{})
	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalw(template string, args ...interface{})
	Stop()
}

type LoggerConfig interface {
	Log() *conf.LogInfo
}

var LoggerIns *Zap
var LoggerInsOnce sync.Once

func NewLogger(cfg LoggerConfig) *Zap {
	log := cfg.Log()
	param := NewParam(log.Level, log.Path, log.MaxSize, log.MaxBackupNum, log.MackupDuration)
	LoggerInsOnce.Do(func() {
		LoggerIns = NewLoggerZap(param)
	})
	return LoggerIns
}

func Printf(template string, args ...interface{}) {
	LoggerIns.Infof(template, args...)
}

func Info(args ...interface{}) {
	LoggerIns.Info(args...)
}

func Infof(template string, args ...interface{}) {
	LoggerIns.Infof(template, args...)
}

func Infow(template string, args ...interface{}) {
	LoggerIns.Infow(template, args...)
}

func Warn(args ...interface{}) {
	LoggerIns.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	LoggerIns.Warnf(template, args...)
}

func Error(args ...interface{}) {
	LoggerIns.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	LoggerIns.Errorf(template, args...)
}

func Errorw(template string, args ...interface{}) {
	LoggerIns.Errorw(template, args...)
}

func Panic(args ...interface{}) {
	LoggerIns.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	LoggerIns.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	LoggerIns.Fatal(args...)
}

func Fatalw(template string, args ...interface{}) {
	LoggerIns.Fatalw(template, args...)
}

func Debugw(template string, args ...interface{}) {
	LoggerIns.Debugw(template, args...)
}

func Debug(args ...interface{}) {
	LoggerIns.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	if LoggerIns.loggerLevel == zap.DebugLevel {
		pc, _, _, _ := runtime.Caller(1)
		callName := runtime.FuncForPC(pc).Name()
		if arr := strings.Split(callName, "."); len(arr) > 0 {
			temp := arr[len(arr)-1] + " " + template
			LoggerIns.Debugf(temp, args...)
			return
		}
	}
	LoggerIns.Debugf(template, args...)
}

func Stop() {
	LoggerIns.Stop()
}
