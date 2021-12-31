package config

import (
	"errors"

	"github.com/aboglioli/configd/common/model"
)

type ConfigData map[string]interface{}

type Config struct {
	schemaSlug *model.Slug
	slug       *model.Slug
	name       *Name
	config     ConfigData
}

func BuildConfig(
	schemaSlug *model.Slug,
	slug *model.Slug,
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
	schemaSlug *model.Slug,
	name *Name,
	config ConfigData,
) (*Config, error) {
	slug, err := model.NewSlug(name.Value())
	if err != nil {
		return nil, err
	}

	return BuildConfig(schemaSlug, slug, name, config)
}

func (c *Config) SchemaSlug() *model.Slug {
	return c.schemaSlug
}

func (c *Config) Slug() *model.Slug {
	return c.slug
}

func (c *Config) Name() *Name {
	return c.name
}

func (c *Config) Config() ConfigData {
	return c.config
}
