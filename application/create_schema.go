package application

import (
	"context"

	"github.com/aboglioli/configd/domain/schema"
)

type CreateSchemaCommand struct {
	Name   string
	Schema map[string]interface{}
}

type CreateSchemaResponse struct {
	Slug string
	Name string
}

type CreateSchema struct {
	schemaRepo schema.SchemaRepository
}

func NewCreateSchema(
	schemaRepo schema.SchemaRepository,
) *CreateSchema {
	return &CreateSchema{
		schemaRepo: schemaRepo,
	}
}

func (uc *CreateSchema) Exec(
	ctx context.Context,
	cmd *CreateSchemaCommand,
) (*CreateSchemaResponse, error) {
	s, err := schema.FromMap(cmd.Name, cmd.Schema)
	if err != nil {
		return nil, err
	}

	if err := uc.schemaRepo.Save(s); err != nil {
		return nil, err
	}

	return &CreateSchemaResponse{
		Slug: s.Slug().Value(),
		Name: s.Name().Value(),
	}, nil
}
