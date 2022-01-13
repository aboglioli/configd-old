package application

import (
	"context"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/pkg/models"
)

type DeleteConfigCommand struct {
	Id string `json:"id"`
}

type DeleteConfigResponse struct {
	Success bool `json:"success"`
}

type DeleteConfig struct {
	configRepo config.ConfigRepository
}

func NewDeleteConfig(
	configRepo config.ConfigRepository,
) *DeleteConfig {
	return &DeleteConfig{
		configRepo: configRepo,
	}
}

func (uc *DeleteConfig) Exec(
	ctx context.Context,
	cmd *DeleteConfigCommand,
) (*DeleteConfigResponse, error) {
	id, err := models.BuildId(cmd.Id)
	if err != nil {
		return nil, err
	}

	c, err := uc.configRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := uc.configRepo.Delete(ctx, c.Base().Id()); err != nil {
		return nil, err
	}

	return &DeleteConfigResponse{
		Success: true,
	}, nil
}
