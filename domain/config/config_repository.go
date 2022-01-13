package config

import (
	"context"
	"errors"

	"github.com/aboglioli/configd/pkg/models"
)

var (
	ErrNotFound = errors.New("config not found")
)

type ConfigRepository interface {
	FindById(ctx context.Context, id models.Id) (*Config, error)
	FindBySchemaId(ctx context.Context, schemaId models.Id) ([]*Config, error)
	Save(ctx context.Context, config *Config) error
	Delete(ctx context.Context, id models.Id) error
}
