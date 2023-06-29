package logger

import (
	"os"

	"gin-layout/config"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() (err error) {
	cfg := config.Conf().Logger
	writerSyncer := logWriter(cfg.FileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := logEncoder()
	level := new(zapcore.Level)
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return err
	}
	var core zapcore.Core
	if config.Conf().Site.Mode != gin.ReleaseMode {
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

	l := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(l)
	zap.L().Info("logger init success")
	return
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
