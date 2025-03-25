package service

import (
	"authentication/models"
	"authentication/repo"
	"authentication/utils/validator"
)

type SignUpService struct {
	Repo *repo.SignUpRepository
}

func (s *SignUpService) SignUp(User *models.User) (string, error) {

	if err := validator.ValidateUser(User); err != nil {
		return "", err
	}

	err := s.Repo.CreateUser(User)
	if err != nil {
		return "", err
	}
	return "register successfully", nil
}
