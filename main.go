package main

import (
	"fmt"
	"go-app/models"
	"go-app/routes"
	"log"
)

func Main() {
	fmt.Print("PROGRAM RUN")

	db,err := models.DBConnection()
	if err != nil {
		log.Print("System Error", err)
	}
	routes.SetupRoutes(db)
}
