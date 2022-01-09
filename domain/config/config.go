package config

import (
	"errors"

	"github.com/aboglioli/configd/common/models"
)

type ConfigData map[string]interface{}

type Config struct {
	schemaSlug *models.Slug
	slug       *models.Slug
	name       *Name
	config     ConfigData
}

func BuildConfig(
	schemaSlug *models.Slug,
	slug *models.Slug,
	name *Name,
	config ConfigData,
) (*Config, error) {
	if len(config) == 0 {
		return nil, errors.New("empty configuration")
	}

	return &Config{
		schemaSlug: schemaSlug,
		slug:       slug,
		name:       name,
		config:     config,
	}, nil

}

func NewConfig(
	schemaSlug *models.Slug,
	name *Name,
	config ConfigData,
) (*Config, error) {
	slug, err := models.NewSlug(name.Value())
	if err != nil {
		return nil, err
	}

	return BuildConfig(schemaSlug, slug, name, config)
}

func (c *Config) SchemaSlug() *models.Slug {
	return c.schemaSlug
}

func (c *Config) Slug() *models.Slug {
	return c.slug
}

func (c *Config) Name() *Name {
	return c.name
}

func (c *Config) Config() ConfigData {
	return c.config
}
