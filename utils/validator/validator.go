package validator

import (
	"net/mail"
	"regexp"
)

func ValidateEmail(mailAddress string) bool {
	_, err := mail.ParseAddress(mailAddress)
	return err == nil
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	//8 characters, 1 uppercase, 1 lowercase, 1 number
	passwordPattern := `^[A-Za-z\d]{8,}$`
	_, err := regexp.MatchString(passwordPattern, password)
	return err == nil
}
