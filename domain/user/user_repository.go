package user

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("user not found")
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username Username) (*User, error)
	Save(ctx context.Context, user *User) error
	Delete(ctx context.Context, username Username) error
}
