package Logger

import (
	"log"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		log.Fatal("Ошибка инициализации логгера: ", err)
	}
}

func Sync() {
	Logger.Sync()
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
