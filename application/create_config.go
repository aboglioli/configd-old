package application

import (
	"context"
	"fmt"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/pkg/models"
)

type CreateConfigCommand struct {
	Id       *string           `json:"id"`
	SchemaId string            `json:"schema_id"`
	Name     string            `json:"name"`
	Config   config.ConfigData `json:"config"`
}

type CreateConfigResponse struct {
	SchemaId    string            `json:"schema_id"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Config      config.ConfigData `json:"config"`
	ValidSchema bool              `json:"valid_schema"`
}

type CreateConfig struct {
	schemaRepo schema.SchemaRepository
	configRepo config.ConfigRepository
}

func NewCreateConfig(
	schemaRepo schema.SchemaRepository,
	configRepo config.ConfigRepository,
) *CreateConfig {
	return &CreateConfig{
		configRepo: configRepo,
		schemaRepo: schemaRepo,
	}
}

func (uc *CreateConfig) Exec(
	ctx context.Context,
	cmd *CreateConfigCommand,
) (*CreateConfigResponse, error) {
	// Check schema existence
	schemaId, err := models.BuildId(cmd.SchemaId)
	if err != nil {
		return nil, err
	}

	s, err := uc.schemaRepo.FindById(ctx, schemaId)
	if err != nil {
		return nil, err
	}

	// Name
	name, err := config.NewName(cmd.Name)
	if err != nil {
		return nil, err
	}

	// Id
	var id models.Id
	if cmd.Id != nil {
		id, err = models.NewSlug(*cmd.Id)
	} else {
		id, err = models.NewSlug(name.Value())
	}
	if err != nil {
		return nil, err
	}

	// Check unique id
	if _, err := uc.configRepo.FindById(ctx, id); err != config.ErrNotFound {
		return nil, fmt.Errorf("config with id %s already exists", id.Value())
	}

	// Create new config
	c, err := config.NewConfig(id, schemaId, name, cmd.Config)
	if err != nil {
		return nil, err
	}

	if err := uc.configRepo.Save(ctx, c); err != nil {
		return nil, err
	}

	validSchema := true
	if err := s.Validate(c.Config()); err != nil {
		validSchema = false
	}

	return &CreateConfigResponse{
		Id:          c.Base().Id().Value(),
		SchemaId:    c.SchemaId().Value(),
		Name:        c.Name().Value(),
		Config:      c.Config(),
		ValidSchema: validSchema,
	}, nil
}
