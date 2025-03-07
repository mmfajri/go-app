package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnection() (*gorm.DB, error) {
	USER := "postgres"
	PASS := "root" 
	HOST := "localhost"
	PORT := "1111"
	DBNAME := "go-app-db"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
				SlowThreshold: time.Second,
				LogLevel: logger.Info,
				Colorful: true,
			},
	)

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	return gorm.Open(postgres.Open(url), &gorm.Config{Logger: newLogger})
}
			
