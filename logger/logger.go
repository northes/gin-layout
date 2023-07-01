package logger

import (
	"os"

	"gin-layout/config"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	LogInfo interface {
		Debug(args ...interface{})
		Info(args ...interface{})
		Warn(args ...interface{})
		Error(args ...interface{})
		Panic(args ...interface{})
		Fatal(args ...interface{})
	}

	LogFormat interface {
		Debugf(template string, args ...interface{})
		Infof(template string, args ...interface{})
		Warnf(template string, args ...interface{})
		Errorf(template string, args ...interface{})
		Panicf(template string, args ...interface{})
		Fatalf(template string, args ...interface{})
	}

	LogInfoFormat interface {
		LogInfo
		LogFormat
	}
)

type Logger struct {
	zap *zap.SugaredLogger
}

func NewLogger(cfg *config.AppConf) (LogInfoFormat, error) {
	writerSyncer := logWriter(cfg.FileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := logEncoder()
	level := new(zapcore.Level)
	err := level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return nil, err
	}
	var core zapcore.Core
	if cfg.Site.Mode != gin.ReleaseMode {
		// 开发模式同时输出终端和文件
		console := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writerSyncer, level),
			zapcore.NewCore(console, zapcore.Lock(os.Stdout), level),
		)
	} else {
		// 发布模式输出文件
		core = zapcore.NewCore(encoder, writerSyncer, level)
	}

	l := zap.New(core, zap.AddCaller()).Sugar()
	// zap.ReplaceGlobals(l)
	// zap.L().Info("logger init success")

	return &Logger{zap: l}, nil
}

func logEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func logWriter(name string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   name,
		MaxSize:    maxAge,
		MaxBackups: maxBackup,
		MaxAge:     maxSize,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func (l *Logger) Debug(args ...interface{}) {
	l.zap.Debug(args)
}

func (l *Logger) Info(args ...interface{}) {
	l.zap.Info(args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.zap.Warn(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.zap.Error(args)
}

func (l *Logger) Panic(args ...interface{}) {
	l.zap.Panic(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.zap.Fatal(args)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.zap.Debugf(template, args)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.zap.Infof(template, args)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.zap.Warnf(template, args)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.zap.Errorf(template, args)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.zap.Panicf(template, args)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.zap.Fatalf(template, args)
}
