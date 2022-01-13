package application

import (
	"context"

	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/pkg/models"
)

type DeleteSchemaCommand struct {
	Id string `json:"id"`
}

type DeleteSchemaResponse struct {
	Success bool `json:"success"`
}

type DeleteSchema struct {
	schemaRepo schema.SchemaRepository
}

func NewDeleteSchema(
	schemaRepo schema.SchemaRepository,
) *DeleteSchema {
	return &DeleteSchema{
		schemaRepo: schemaRepo,
	}
}

func (uc *DeleteSchema) Exec(
	ctx context.Context,
	cmd *DeleteSchemaCommand,
) (*DeleteSchemaResponse, error) {
	id, err := models.BuildId(cmd.Id)
	if err != nil {
		return nil, err
	}

	s, err := uc.schemaRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Delete
	if err := uc.schemaRepo.Delete(ctx, s.Base().Id()); err != nil {
		return nil, err
	}

	return &DeleteSchemaResponse{
		Success: true,
	}, nil
}
