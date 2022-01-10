package application

import (
	"context"

	"github.com/aboglioli/configd/domain/schema"
)

type CreateSchemaCommand struct {
	Name   string                 `json:"name"`
	Schema map[string]interface{} `json:"schema"`
}

type CreateSchemaResponse struct {
	Id     string                 `json:"id"`
	Name   string                 `json:"name"`
	Schema map[string]interface{} `json:"schema"`
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
	props, err := schema.PropsFromMap(cmd.Schema)
	if err != nil {
		return nil, err
	}

	n, err := schema.NewName(cmd.Name)
	if err != nil {
		return nil, err
	}

	s, err := schema.NewSchema(n, props...)
	if err != nil {
		return nil, err
	}

	if err := uc.schemaRepo.Save(s); err != nil {
		return nil, err
	}

	return &CreateSchemaResponse{
		Id:     s.Base().Id().Value(),
		Name:   s.Name().Value(),
		Schema: s.ToMap(),
	}, nil
}
