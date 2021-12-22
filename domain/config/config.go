package config

import (
	"errors"

	"github.com/aboglioli/configd/domain/schema"
)

type Config struct {
	schemaName *schema.Name
	name       *Name
	config     map[string]interface{}
}

func NewConfig(
	schemaName *schema.Name,
	name *Name,
	config map[string]interface{},
) (*Config, error) {
	if len(config) == 0 {
		return nil, errors.New("empty configuration")
	}

	return &Config{
		schemaName: schemaName,
		name: name,
		config: config,
	}, nil
}

func (c *Config) SchemaName() *schema.Name {
	return c.schemaName
}

func (c *Config) Name() *Name {
	return c.name
}

func (c *Config) Config() map[string]interface{} {
	return c.config
}
