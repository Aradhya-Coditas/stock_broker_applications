package main

import (
	"authentication/models"
	"authentication/router"
	"authentication/utils"
	"authentication/utils/db"
	"log"
)

func main() {
	_, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbInstance, err := db.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := dbInstance.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	routerInstance := router.SetupRouter(dbInstance)

	log.Printf("Server starting on port %s", "8080")
	if err := routerInstance.Run(":" + "8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
