package logger

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
