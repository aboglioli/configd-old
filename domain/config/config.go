package config

import (
	"errors"

	"github.com/aboglioli/configd/pkg/models"
)

type ConfigData map[string]interface{}

type Config struct {
	agg *models.AggregateRoot

	schemaId models.Id
	name     Name
	config   ConfigData
}

func BuildConfig(
	slug models.Id,
	schemaId models.Id,
	name Name,
	config ConfigData,
) (*Config, error) {
	if len(config) == 0 {
		return nil, errors.New("empty configuration")
	}

	agg, err := models.NewAggregateRoot(slug)
	if err != nil {
		return nil, err
	}

	return &Config{
		agg:      agg,
		schemaId: schemaId,
		name:     name,
		config:   config,
	}, nil

}

func NewConfig(
	schemaId models.Id,
	name Name,
	config ConfigData,
) (*Config, error) {
	slug, err := models.NewSlug(name.Value())
	if err != nil {
		return nil, err
	}

	return BuildConfig(schemaId, slug, name, config)
}

func (c *Config) Base() *models.AggregateRoot {
	return c.agg
}

func (c *Config) SchemaId() models.Id {
	return c.schemaId
}

func (c *Config) Name() Name {
	return c.name
}

func (c *Config) Config() ConfigData {
	return c.config
}
