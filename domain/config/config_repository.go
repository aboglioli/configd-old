package config

import (
	"errors"

	"github.com/aboglioli/configd/pkg/models"
)

var (
	ErrNotFound = errors.New("config not found")
)

type ConfigRepository interface {
	FindById(slug models.Id) (*Config, error)
	FindBySchemaId(schemaId models.Id) ([]*Config, error)
	Save(config *Config) error
	Delete(slug models.Id) error
}
