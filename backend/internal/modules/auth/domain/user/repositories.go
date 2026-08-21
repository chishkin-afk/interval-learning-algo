package user

import (
	"context"
	"errors"
	"uuid"
)

var (
	ErrAlreadyExists = errors.New("user already exists")
	ErrNotFound      = errors.New("user not found")
)

type UpdateFunc func(*User) error

type UserPersistenceRepository interface {
	Save(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email Email) (*User, error)
	Update(ctx context.Context, updFunc UpdateFunc, id uuid.UUID) (*User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
