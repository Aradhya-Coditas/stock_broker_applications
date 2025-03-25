package constants

const (
	InvalidPassword                   = "invalid password"
	ErrNoCapitalStart                 = "username must start with a capital letter"
	ErrInvalidUsernameLength          = "username must be between 3 and 20 characters"
	ErrContainsSpaces                 = "username cannot contain spaces"
	ErrInvalidCharacters              = "username can only contain letters and numbers"
	ErrEmptyUsername                  = "username cannot be empty"
	ErrEmptyPhoneNumber               = "phone number cannot be empty"
	ErrInvalidPhoneNumberLength       = "phone number must be between 10 and 15 digits (including country code)"
	ErrNoCountryCode                  = "phone number must start with a country code (e.g., +1, +91)"
	ErrInvalidCharactersofphonenumber = "phone number can only contain digits, +, and optional spaces or hyphens"
	ErrTooManySeparators              = "phone number cannot have more than 2 separators (spaces or hyphens)"
	ErrEmptyEmail                     = "email cannot be empty"
	ErrInvalidEmail                   = "invalid email format"
	ErrNoAtSymbol                     = "email must contain exactly one @ symbols"
	ErrMultipleAtSymbols              = "email cannot contain multiple @ symbols"
	ErrInvalidPassword                = "password must be at least 8 characters"
	ErrInvalidPanCard                 = "PAN card must be 10 characters (e.g., ABCDE1234F)"
)
