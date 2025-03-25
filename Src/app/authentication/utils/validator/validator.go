package validator

import (
	"authentication/models"
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidUsername    = errors.New("username must be 3-20 characters and alphanumeric")
	ErrInvalidPhoneNumber = errors.New("phone number must be 10 digits")
	ErrInvalidPanCard     = errors.New("PAN card must be 10 characters (e.g., ABCDE1234F)")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidPassword    = errors.New("password must be at least 8 characters")
)

func ValidateUser(User *models.User) error {
	if err := ValidateUsername(User.UserName); err != nil {
		return err
	}
	if err := ValidatePassword(User.Password); err != nil {
		return err
	}
	if err := ValidateEmail(User.Email); err != nil {
		return err
	}
	if err := ValidatePhoneNumber(User.PhoneNumber); err != nil {
		return err
	}
	if err := ValidatePanCard(User.PanCard); err != nil {
		return err
	}
	return nil
}

func ValidateUsername(username string) error {
	if len(username) < 3 || len(username) > 20 {
		return ErrInvalidUsername
	}
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, username); !matched {
		return ErrInvalidUsername
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrInvalidPassword
	}
	return nil
}

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email); !matched {
		return ErrInvalidEmail
	}
	return nil
}

func ValidatePhoneNumber(phone string) error {
	phone = strings.TrimSpace(phone)
	if matched, _ := regexp.MatchString(`^[0-9]{10}$`, phone); !matched {
		return ErrInvalidPhoneNumber
	}
	return nil
}

func ValidatePanCard(pan string) error {
	pan = strings.TrimSpace(pan)
	if matched, _ := regexp.MatchString(`^[A-Z]{5}[0-9]{4}[A-Z]$`, pan); !matched {
		return ErrInvalidPanCard
	}
	return nil
}
