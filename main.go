package main

import (
	"authentication/utils/db"
	"log"
	"stock_broker_applications/src/models"
)

func main() {
	dbInstance, err := db.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	dbInstance.AutoMigrate(&models.User{})
}
