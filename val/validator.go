package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z0-9\\s]+$`).MatchString
)

func Validatestring(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d character", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := Validatestring(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only letters, digits, or underscore")
	}
	return nil
}

func ValidatePassword(value string) error {
	return Validatestring(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := Validatestring(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil

}

func ValidateFullname(value string) error {
	if err := Validatestring(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullname(value) {
		return fmt.Errorf("must contain only letters, digits, or underscore")
	}
	return nil
}
