package main

import (
	"authentication/models"
	"authentication/utils/db"
	"log"
)

func main() {
	dbInstance, err := db.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	dbInstance.AutoMigrate(&models.User{})
}
