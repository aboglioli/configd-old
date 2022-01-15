package application

import (
	"context"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/pkg/models"
)

type GetConfigCommand struct {
	Id string `json:"id"`
}

type GetConfigResponse struct {
	Id          string            `json:"id"`
	SchemaId    string            `json:"schema_id"`
	Name        string            `json:"name"`
	Config      config.ConfigData `json:"config"`
	ValidSchema bool              `json:"valid_schema"`
}

type GetConfig struct {
	schemaRepo schema.SchemaRepository
	configRepo config.ConfigRepository
}

func NewGetConfig(
	schemaRepo schema.SchemaRepository,
	configRepo config.ConfigRepository,
) *GetConfig {
	return &GetConfig{
		schemaRepo: schemaRepo,
		configRepo: configRepo,
	}
}

func (uc *GetConfig) Exec(
	ctx context.Context,
	cmd *GetConfigCommand,
) (*GetConfigResponse, error) {
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

	validSchema := true
	if err := s.Validate(c.Config()); err != nil {
		validSchema = false
	}

	return &GetConfigResponse{
		Id:          c.Base().Id().Value(),
		SchemaId:    c.SchemaId().Value(),
		Name:        c.Name().Value(),
		Config:      c.Config(),
		ValidSchema: validSchema,
	}, nil
}
