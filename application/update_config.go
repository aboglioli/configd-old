package application

import (
	"context"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/pkg/models"
)

type UpdateConfigCommand struct {
	Id     string  `json:"id"`
	Name   *string `json:"name"`
	Config *config.ConfigData
}

type UpdateConfigResponse struct {
	SchemaId    string            `json:"schema_id"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Config      config.ConfigData `json:"config"`
	ValidSchema bool              `json:"valid_schema"`
}

type UpdateConfig struct {
	schemaRepo schema.SchemaRepository
	configRepo config.ConfigRepository
}

func NewUpdateConfig(
	schemaRepo schema.SchemaRepository,
	configRepo config.ConfigRepository,
) *UpdateConfig {
	return &UpdateConfig{
		configRepo: configRepo,
		schemaRepo: schemaRepo,
	}
}

func (uc *UpdateConfig) Exec(
	ctx context.Context,
	cmd *UpdateConfigCommand,
) (*UpdateConfigResponse, error) {
	id, err := models.BuildId(cmd.Id)
	if err != nil {
		return nil, err
	}

	c, err := uc.configRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	s, err := uc.schemaRepo.FindById(ctx, c.SchemaId())
	if err != nil {
		return nil, err
	}

	// Update parameteres
	if cmd.Name != nil {
		name, err := config.NewName(*cmd.Name)
		if err != nil {
			return nil, err
		}

		c.ChangeName(name)
	}

	if cmd.Config != nil {
		c.ChangeConfig(*cmd.Config)
	}

	if err := uc.configRepo.Save(ctx, c); err != nil {
		return nil, err
	}

	validSchema := true
	if err := s.Validate(c.Config()); err != nil {
		validSchema = false
	}

	return &UpdateConfigResponse{
		Id:          c.Base().Id().Value(),
		SchemaId:    c.SchemaId().Value(),
		Name:        c.Name().Value(),
		Config:      c.Config(),
		ValidSchema: validSchema,
	}, nil
}
