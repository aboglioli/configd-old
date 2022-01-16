package config

import (
	"errors"

	"github.com/aboglioli/configd/pkg/events"
	"github.com/aboglioli/configd/pkg/models"
)

type Config struct {
	agg *models.AggregateRoot

	schemaId models.Id
	name     Name
	config   ConfigData
}

func BuildConfig(
	id models.Id,
	schemaId models.Id,
	name Name,
	config ConfigData,
) (*Config, error) {
	if len(config) == 0 {
		return nil, errors.New("empty configuration")
	}

	agg, err := models.NewAggregateRoot(id)
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
	id models.Id,
	schemaId models.Id,
	name Name,
	config ConfigData,
) (*Config, error) {
	c, err := BuildConfig(id, schemaId, name, config)
	if err != nil {
		return nil, err
	}

	event, err := events.NewEvent(
		c.agg.Id().Value(),
		CreatedTopic,
		Created{
			Id:        c.agg.Id().Value(),
			SchemaId:  c.schemaId.Value(),
			Name:      c.name.Value(),
			Config:    c.config,
			ConfigSum: c.config.Hash(),
		},
	)
	if err != nil {
		return nil, err
	}

	c.agg.RecordEvent(event)

	return c, nil
}

func (c *Config) Base() models.ReadOnlyAggregateRoot {
	return c.agg
}

func (c *Config) SchemaId() models.Id {
	return c.schemaId
}

func (c *Config) Name() Name {
	return c.name
}

func (c *Config) ChangeName(name Name) error {
	c.name = name
	c.agg.Update()

	event, err := events.NewEvent(
		c.agg.Id().Value(),
		NameChangedTopic,
		NameChanged{
			Id:   c.agg.Id().Value(),
			Name: c.name.Value(),
		},
	)
	if err != nil {
		return err
	}

	c.agg.RecordEvent(event)

	return nil
}

func (c *Config) Config() ConfigData {
	return c.config
}

func (c *Config) ChangeConfig(config ConfigData) error {
	c.config = config
	c.agg.Update()

	event, err := events.NewEvent(
		c.agg.Id().Value(),
		ConfigChangedTopic,
		ConfigChanged{
			Id:        c.agg.Id().Value(),
			Config:    c.config,
			ConfigSum: c.config.Hash(),
		},
	)
	if err != nil {
		return err
	}

	c.agg.RecordEvent(event)

	return nil
}
