package application

import (
	"context"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/domain/security"
	"github.com/aboglioli/configd/pkg/models"
)

type GetConfigCommand struct {
	Id     string `json:"id"`
	ApiKey string `json:"api_key"`
}

type GetConfigResponse struct {
	Id          string            `json:"id"`
	SchemaId    string            `json:"schema_id"`
	Name        string            `json:"name"`
	Config      config.ConfigData `json:"config"`
	ValidSchema bool              `json:"valid_schema"`
	ConfigSum   string            `json:"config_sum"`
}

type GetConfig struct {
	schemaRepo        schema.SchemaRepository
	configRepo        config.ConfigRepository
	authorizationRepo security.AuthorizationRepository
}

func NewGetConfig(
	schemaRepo schema.SchemaRepository,
	configRepo config.ConfigRepository,
	authorizationRepo security.AuthorizationRepository,
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

	// Check API Key
	apiKey, err := security.NewApiKey(cmd.ApiKey)
	if err != nil {
		return nil, err
	}

	hashedApiKey, err := apiKey.Hash()
	if err != nil {
		return nil, err
	}

	auth, err := uc.authorizationRepo.FindByApiKey(ctx, hashedApiKey)
	if err != nil {
		return nil, ErrUnauthorized
	}

	if !auth.ResourceId().Equals(id) {
		return nil, ErrUnauthorized
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
		ConfigSum:   c.Config().Hash(),
	}, nil
}
