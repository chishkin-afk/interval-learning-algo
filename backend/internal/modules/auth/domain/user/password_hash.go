package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
)

// InvalidPasswordError is a error structure,
// that can be used to check validations of password.
//
// It's implements standart error interface
type InvalidPasswordError struct {
	original error
	msg string
}

func (ip *InvalidPasswordError) Error() string {
	return fmt.Sprintf("%s: %s", 
		ip.original.Error(), 
		ip.msg)
}

func (ip *InvalidPasswordError) Unwrap() error {
	return ip.original
}

// PasswordHash is a value-object of User, 
// responses to hashing & comparing user password.
type PasswordHash string

func (ph PasswordHash) String() string {
	return string(ph)
}

// Compare checks hash and raw password to equal.
// If hash == password returns true, otherwise false
func (ph PasswordHash) Compare(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(ph.String()), []byte(password)) == nil
}

func NewPasswordHash(password string) (PasswordHash, error) {
	n := len([]rune(password))
	if n < 3 || n > 32 {
		return "", &InvalidPasswordError{
			original: ErrInvalidPassword,
			msg: "len of password must be more than 2 and less than 32",
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", &InvalidPasswordError{
			original: ErrInvalidPassword,
			msg: "failed to generate hash",
		}
	}

	return PasswordHash(hash), nil
}