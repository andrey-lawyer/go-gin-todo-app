package errors

import "unicode"

type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string {
	return e.Msg
}

func ValidatePassword(password string) error {
	if len(password) < 4 {
		return &ValidationError{Msg: "password must be at least 4 characters long"}
	}
	hasUpper := false
	for _, c := range password {
		if unicode.IsUpper(c) {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return &ValidationError{Msg: "the password must contain at least one capital letter"}
	}
	return nil
}
