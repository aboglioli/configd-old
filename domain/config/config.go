package config

import (
	"errors"
)

type Config struct {
	schemaName *Name
	name       *Name
	config     map[string]interface{}
}

func NewConfig(
	schemaName *Name,
	name *Name,
	config map[string]interface{},
) (*Config, error) {
	if len(config) == 0 {
		return nil, errors.New("empty configuration")
	}

	return &Config{
		schemaName: schemaName,
		name:       name,
		config:     config,
	}, nil
}

func (c *Config) SchemaName() *Name {
	return c.schemaName
}

func (c *Config) Name() *Name {
	return c.name
}

func (c *Config) Config() map[string]interface{} {
	return c.config
}
