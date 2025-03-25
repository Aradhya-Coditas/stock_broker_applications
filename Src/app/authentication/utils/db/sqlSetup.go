package db

import (
	"authentication/models"
	"authentication/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	config, err := utils.LoadConfig()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBName)

	fmt.Println("dsn", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

type User struct {
}
