package user

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

// InvalidEmailError is a structure,
// that can be used to check validation of user email
//
// Its fully implements standart error interface
type InvalidEmailError struct {
	original error
	msg string
}

func (ie *InvalidEmailError) Error() string {
	return fmt.Sprintf("%s: %s", ie.original.Error(), ie.msg)
}

func (ie *InvalidEmailError) Unwrap() error {
	return ie.original
}

// Email is a value-object of User,
// responses to validate & normalize user email
type Email string

func (e Email) String() string {
	return string(e)
}

// Norm trim spaces of email and provides its to lower case.
// returns new value-object Email
func (e Email) Norm() Email {
	trimmed := strings.TrimSpace(e.String())
	return Email(strings.ToLower(trimmed))
}

func (e Email) Validate() error {
	normalized := e.Norm()	
	if _, err := mail.ParseAddress(normalized.String()); err != nil {
		return &InvalidEmailError{
			original: ErrInvalidEmail,
			msg: err.Error(),
		}
	}

	return nil
}
