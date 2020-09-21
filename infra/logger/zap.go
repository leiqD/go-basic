package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	StructMessageFilePath string = "filePath"
	StructMessageError    string = "errMsg"
	StructMessageDataNum  string = "dataNum"
)

type Zap struct {
	logger      *zap.Logger
	sugar       *zap.SugaredLogger
	loggerLevel zapcore.Level
}

func NewLoggerZap(param *Param) *Zap {
	logfilePath := param.Path

	wh := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logfilePath,
		MaxSize:    param.MaxSize, // megabytes
		MaxBackups: param.MaxBackupNum,
		MaxAge:     param.MackupDuration, // days
		LocalTime:  true,
		Compress:   true,
	})

	getConfigLevel := func() zapcore.Level {
		configLevel := param.Level
		switch configLevel {
		case "INFO":
			return zap.InfoLevel
		case "DEBUG":
			return zap.DebugLevel
		default:
			return zap.InfoLevel
		}
	}

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	//jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		//zapcore.NewCore(jsonEncoder, w, zap.NewAtomicLevelAt(getConfigLevel())),
		zapcore.NewCore(consoleEncoder, wh, zap.NewAtomicLevelAt(getConfigLevel())),
	)

	//core := zapcore.NewCore(
	//	zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
	//	w,
	//	zap.NewAtomicLevelAt(getConfigLevel()),
	//)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))

	zap := &Zap{
		logger:      logger,
		sugar:       logger.Sugar(),
		loggerLevel: getConfigLevel(),
	}

	zap.Infow("logger init zap finish",
		StructMessageFilePath, logfilePath,
	)

	return zap
}

func (lz *Zap) Printf(template string, args ...interface{}) {
	lz.sugar.Infof(template, args...)
}

func (lz *Zap) Info(args ...interface{}) {
	lz.sugar.Info(args...)
}

func (lz *Zap) Infof(template string, args ...interface{}) {
	lz.sugar.Infof(template, args...)
}

func (lz *Zap) Infow(template string, args ...interface{}) {
	lz.sugar.Infow(template, args...)
}

func (lz *Zap) Warn(args ...interface{}) {
	lz.sugar.Warn(args...)
}

func (lz *Zap) Warnf(template string, args ...interface{}) {
	lz.sugar.Warnf(template, args...)
}

func (lz *Zap) Error(args ...interface{}) {
	lz.sugar.Error(args...)
}

func (lz *Zap) Errorf(template string, args ...interface{}) {
	lz.sugar.Errorf(template, args...)
}

func (lz *Zap) Errorw(template string, args ...interface{}) {
	lz.sugar.Errorw(template, args...)
}

func (lz *Zap) Panic(args ...interface{}) {
	lz.sugar.Panic(args...)
}

func (lz *Zap) Panicf(template string, args ...interface{}) {
	lz.sugar.Panicf(template, args...)
}

func (lz *Zap) Fatal(args ...interface{}) {
	lz.sugar.Fatal(args...)
}

func (lz *Zap) Fatalw(template string, args ...interface{}) {
	lz.sugar.Fatalw(template, args...)
}

func (lz *Zap) Debugw(template string, args ...interface{}) {
	lz.sugar.Debugw(template, args...)
}

func (lz *Zap) Debug(args ...interface{}) {
	lz.sugar.Debug(args...)
}

func (lz *Zap) Debugf(template string, args ...interface{}) {
	lz.sugar.Debugf(template, args...)
}

func (lz *Zap) Stop() {
	lz.logger.Sync()
	lz.sugar.Sync()
}
