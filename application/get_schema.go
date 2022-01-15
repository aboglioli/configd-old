package application

import (
	"context"

	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/pkg/models"
)

type GetSchemaCommand struct {
	Id string `json:"id"`
}

type GetSchemaResponse struct {
	Id     string                 `json:"id"`
	Name   string                 `json:"name"`
	Schema map[string]interface{} `json:"schema"`
}

type GetSchema struct {
	schemaRepo schema.SchemaRepository
}

func NewGetSchema(
	schemaRepo schema.SchemaRepository,
) *GetSchema {
	return &GetSchema{
		schemaRepo: schemaRepo,
	}
}

func (uc *GetSchema) Exec(
	ctx context.Context,
	cmd *GetSchemaCommand,
) (*GetSchemaResponse, error) {
	id, err := models.BuildId(cmd.Id)
	if err != nil {
		return nil, err
	}

	s, err := uc.schemaRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetSchemaResponse{
		Id:     s.Base().Id().Value(),
		Name:   s.Name().Value(),
		Schema: s.ToMap(),
	}, nil
}
