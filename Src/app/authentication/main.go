package main

import (
	"authentication/handlers"
	"authentication/models"
	"authentication/repo"
	"authentication/router"
	"authentication/service"
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
	
	dbInstance.AutoMigrate(&models.User{})

	repoInstance := &repo.SignUpRepository{DB: dbInstance}
	serviceInstance := &service.SignUpService{Repo: repoInstance}
	handlerInstance := &handlers.SignUpHandler{Service: serviceInstance}

	routerInstance := router.SetupRouter(handlerInstance)
	log.Printf("Server starting on port %s", "8080")
	log.Fatal(routerInstance.Run(":" + "8080"))
}
