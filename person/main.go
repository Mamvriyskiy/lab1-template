package main

import (
	"go.uber.org/zap"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
	logger "github.com/Mamvriyskiy/lab1-template/person/logger"
	server "github.com/Mamvriyskiy/lab1-template/server"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		logger.Fatal("Ошибка чтения файла конфигурации: ", zap.Error(err))
		return
	}

	logger.Info("Файл конфигурации успешно прочитан")

	if err := godotenv.Load(); err != nil {
		logger.Log("Error", "Load", "Load env file:", err, "")
		fmt.Println(err)
		return
	}
	logger.Log("Info", "", "Load env", nil)

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		fmt.Println(err)
		logger.Log("Error", "initCongig", "Error config DB:", err, "")
		return
	}

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		logger.Log("Error", "Run", "Error occurred while running http server:", err, "")
		return
	}
}
