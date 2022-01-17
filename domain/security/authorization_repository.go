package security

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("authorization not found")
)

type AuthorizationRepository interface {
	FindByApiKey(ctx context.Context, hashedApiKey HashedApiKey) (*Authorization, error)
	Save(ctx context.Context, authorization *Authorization) error
	Delete(ctx context.Context, hashedApiKey HashedApiKey) error
}
