package validator

import (
	"authentication/constants"
	"authentication/models"
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// var (
// 	ErrInvalidUsernameLength          = errors.New("username must be between 3 and 20 characters")
// 	//ErrNoCapitalStart                 = errors.New("username must start with a capital letter")
// 	ErrContainsSpaces                 = errors.New("username cannot contain spaces")
// 	ErrInvalidCharacters              = errors.New("username can only contain letters and numbers")
// 	ErrEmptyUsername                  = errors.New("username cannot be empty")
// 	ErrEmptyPhoneNumber               = errors.New("phone number cannot be empty")
// 	ErrInvalidPhoneNumberLength       = errors.New("phone number must be between 10 and 15 digits (including country code)")
// 	ErrNoCountryCode                  = errors.New("phone number must start with a country code (e.g., +1, +91)")
// 	ErrInvalidCharactersofphonenumber = errors.New("phone number can only contain digits, +, and optional spaces or hyphens")
// 	ErrTooManySeparators              = errors.New("phone number cannot have more than 2 separators (spaces or hyphens)")
// 	ErrEmptyEmail                     = errors.New("email cannot be empty")
// 	ErrInvalidEmail                   = errors.New("invalid email format")
// 	ErrNoAtSymbol                     = errors.New("email must contain exactly one @ symbol")
// 	ErrMultipleAtSymbols              = errors.New("email cannot contain multiple @ symbols")
// 	ErrInvalidPassword                = errors.New("password must be at least 8 characters")
// 	ErrInvalidPanCard                 = errors.New("PAN card must be 10 characters (e.g., ABCDE1234F)")
// )

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

	if username == "" {
		return errors.New(constants.ErrEmptyUsername)
	}

	if len(username) < 3 || len(username) > 20 {
		return errors.New(constants.ErrInvalidUsernameLength)
	}

	if !unicode.IsUpper(rune(username[0])) {
		return errors.New(constants.ErrNoCapitalStart)
	}

	if regexp.MustCompile(`\s`).MatchString(username) {
		return errors.New(constants.ErrContainsSpaces)
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, username); !matched {
		return errors.New(constants.ErrInvalidCharacters)
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New(constants.ErrInvalidPassword)
	}
	return nil
}

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.New(constants.ErrEmptyEmail)
	}

	if strings.Contains(email, " ") {
		return errors.New(constants.ErrContainsSpaces)
	}

	atCount := strings.Count(email, "@")
	if atCount == 0 {
		return errors.New(constants.ErrNoAtSymbol)
	}
	if atCount > 1 {
		return errors.New(constants.ErrMultipleAtSymbols)
	}

	email = strings.TrimSpace(email)
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email); !matched {
		return errors.New(constants.ErrInvalidEmail)
	}
	return nil
}

func ValidatePhoneNumber(phone string) error {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return errors.New(constants.ErrEmptyPhoneNumber)
	}

	if !strings.HasPrefix(phone, "+") {
		return errors.New(constants.ErrNoCountryCode)
	}

	digitCount := 0
	separatorCount := 0
	for _, r := range phone {
		if unicode.IsDigit(r) {
			digitCount++
		} else if r == ' ' || r == '-' {
			separatorCount++
		}
	}

	if digitCount < 10 || digitCount > 15 {
		return errors.New(constants.ErrInvalidPhoneNumberLength)
	}

	if separatorCount > 2 {
		return errors.New(constants.ErrTooManySeparators)
	}

	if matched, _ := regexp.MatchString(`^\+[0-9]+([ -]?[0-9]+)*$`, phone); !matched {
		return errors.New(constants.ErrInvalidCharactersofphonenumber)
	}

	return nil
}

func ValidatePanCard(pan string) error {
	pan = strings.TrimSpace(pan)
	if matched, _ := regexp.MatchString(`^[A-Z]{5}[0-9]{4}[A-Z]$`, pan); !matched {
		return errors.New(constants.ErrInvalidPanCard)
	}
	return nil
}
