package main

import (
	"fmt"
	"go-app/models"
	"go-app/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Print("PROGRAM RUN")
	//Load .env file First
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db,err := models.DBConnection()
	if err != nil {
		log.Print("System Error", err)
	}
	routes.SetupRoutes(db)
}
