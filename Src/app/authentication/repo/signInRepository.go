package repo

import (
	"authentication/models"
	"log"

	"gorm.io/gorm"
)

type SignInRepository struct {
	DB *gorm.DB
}

func (r *SignInRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
    err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("User not found for email: %s, error: %v", email, err)
		return nil, err
	}
	return &user, nil
}
