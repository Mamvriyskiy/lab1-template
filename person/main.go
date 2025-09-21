package main

import (
	"github.com/spf13/viper"
	logger "github.com/Mamvriyskiy/lab1-template/person/logger"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		logger.Fatal("Ошибка чтения файла конфигурации: ", err)
		return
	}

	logger.Info("Файл конфигурации успешно прочитан")
}
