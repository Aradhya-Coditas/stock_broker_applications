package repo

import (
	"authentication/models"

	"gorm.io/gorm"
)

type SignUpRepository struct {
	DB *gorm.DB
}

func (r *SignUpRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *SignUpRepository) FindUserByUsername(Username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("UserName = ?", Username).First(&user).Error
	return &user, err
}
