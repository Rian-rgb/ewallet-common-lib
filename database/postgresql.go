package database

import (
	"fmt"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type PostgresConfig struct {
	Host         string
	User         string
	Password     string
	DBName       string
	Port         string
	MaxIdleConns int
	MaxOpenConns int
}

func NewPostgresClient(cfg PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database %s: %v", cfg.DBName, err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)
	}

	logger.Info("Successfully connected to database: %s", cfg.DBName)
	return db, nil
}
