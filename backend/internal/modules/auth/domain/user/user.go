package user

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidUsername = errors.New("invalid username")
)

// User is a root aggragate of auth module.
type User struct {
	id uuid.UUID
	username string
	email Email
	passwordHash PasswordHash
	createdAt time.Time
	updatedAt time.Time
}

func New(
	username string,
	email Email,
	password string,
) (*User, error) {
	username = strings.TrimSpace(username)
	if err := validateUsername(username); err != nil {
		return nil, fmt.Errorf("%w: %s", err, username)
	}

	email = email.Norm()
	if err := email.Validate(); err != nil {
		return nil, err
	}

	passwordHash, err := NewPasswordHash(password)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return &User{
		id: uuid.New(),
		username: username,
		email: email,
		passwordHash: passwordHash,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func From(
	id uuid.UUID,
	username string,
	email Email,
	passwordHash PasswordHash,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		id: id,
		username: username,
		email: email,
		passwordHash: passwordHash,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) PasswordHash() PasswordHash {
	return u.passwordHash
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func validateUsername(username string) error {
	n := len([]rune(username))
	if n < 2 || n > 128 {
		return ErrInvalidUsername
	}

	return nil
}