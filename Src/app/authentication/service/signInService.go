package service

import "authentication/repo"

type SignInService struct {
	Repo *repo.SignInRepository
}

func (s *SignInService) SignIn(email, password string) (string, error) {
	user, err := s.Repo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "invalid password", nil
	}
	return "logged in successfully", nil
}
