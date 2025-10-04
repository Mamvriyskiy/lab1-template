package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(development bool) error {
	var config zap.Config
	
	if development {
		// Красивый вывод для разработки
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.TimeKey = "time"
		config.EncoderConfig.LevelKey = "level"
		config.EncoderConfig.NameKey = "logger"
		config.EncoderConfig.CallerKey = "caller"
		config.EncoderConfig.MessageKey = "msg"
		config.EncoderConfig.StacktraceKey = "stacktrace"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	} else {
		// Структурированный вывод для продакшена
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "time"
		config.EncoderConfig.LevelKey = "level"
		config.EncoderConfig.MessageKey = "msg"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}
	
	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}
	
	return nil
}

func Fatal(msg string, args ...zap.Field) {
	Logger.Fatal(msg, args...)
}

func Error(msg string, args ...zap.Field) {
	Logger.Error(msg, args...)
}

func Debug(msg string, args ...zap.Field) {
	Logger.Debug(msg, args...)
}

func Info(msg string, args ...zap.Field) {
	Logger.Info(msg, args...)
}
