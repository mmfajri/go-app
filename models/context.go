package models

import (
	"fmt"
	// "log"
	"os"
	// "time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func DBConnection() (*gorm.DB, error) {

	USER := os.Getenv("USER")
	PASS := os.Getenv("PASS")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("DBNAME")
	SCHEMA := os.Getenv("SCHEMA")
	fmt.Println("DB Config:", USER, PASS, HOST, PORT, DBNAME, "Schema:", SCHEMA)

	// Define DSN with search_path to use the correct schema
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
		HOST, USER, PASS, DBNAME, PORT, SCHEMA,
	)

	// Open DB connection with GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Set schema explicitly for future queries
	if err := db.Exec("SET search_path TO " + SCHEMA).Error; err != nil {
		return nil, err
	}

	return db, nil
}
			
