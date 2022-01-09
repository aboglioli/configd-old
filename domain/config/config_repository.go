package config

import (
	"errors"

	"github.com/aboglioli/configd/common/models"
)

var (
	ErrNotFound = errors.New("config not found")
)

type ConfigRepository interface {
	FindById(slug *models.Slug) (*Config, error)
	FindBySchemaId(schemaSlug *models.Slug) ([]*Config, error)
	Save(config *Config) error
	Delete(slug *models.Slug) error
}
