package repository

import (
	"fmt"

	logger "github.com/Mamvriyskiy/lab1-template/person/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	// Импорт драйвера PostgreSQL для его регистрации.
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logger.Fatal("Error connect DB:", zap.Error(err))
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Fatal("Error connect DB:", zap.Error(err))
		return nil, err
	}

	return db, nil
}
