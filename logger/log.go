package logger

import (
	"os"
	"path"

	"github.com/northes/gin-layout/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.Logger

func Init() (err error) {
	writer, err := logWriter(config.Log.Path)
	if err != nil {
		return err
	}
	core := zapcore.NewCore(logEncoder(), writer, toLevel(config.Log.Level))
	L = zap.New(core, zap.AddCaller())
	//L = logger.Sugar()
	return nil
}

func logEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 将时间格式化为人类可读
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 将Level格式化为大写(可选CapitalColorLevelEncoder为Level添加颜色)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func logWriter(logPath string) (zapcore.WriteSyncer, error) {
	if err := os.MkdirAll(path.Dir(logPath), os.ModePerm); err != nil {
		return nil, err
	}
	file, err := os.Create(logPath)
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(file), nil
}

func toLevel(level string) zapcore.Level {
	switch level {
	case zapcore.DebugLevel.String():
		return zapcore.DebugLevel
	case zapcore.InfoLevel.String():
		return zapcore.InfoLevel
	case zapcore.WarnLevel.String():
		return zapcore.WarnLevel
	case zapcore.ErrorLevel.String():
		return zapcore.ErrorLevel
	case zapcore.DPanicLevel.String():
		return zapcore.DPanicLevel
	case zapcore.FatalLevel.String():
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
