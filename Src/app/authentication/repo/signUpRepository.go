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
