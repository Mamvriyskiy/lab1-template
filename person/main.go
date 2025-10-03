package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	handler "github.com/Mamvriyskiy/lab1-template/person/handler"
	logger "github.com/Mamvriyskiy/lab1-template/person/logger"
	repo "github.com/Mamvriyskiy/lab1-template/person/repository"
	server "github.com/Mamvriyskiy/lab1-template/person/server"
	service "github.com/Mamvriyskiy/lab1-template/person/services"
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

	// if err := godotenv.Load(); err != nil {
	// 	logger.Fatal("Error load env file:", zap.Error(err))
	// 	return
	// }
	// logger.Info("Load env")

	db, err := repo.NewPostgresDB(&repo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logger.Fatal("Error connect db:", zap.Error(err))
		return
	}

	repos := repo.NewRepository(db)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRouters()); err != nil {
		logger.Fatal("Error occurred while running http server:", zap.Error(err))
		return
	}
}
