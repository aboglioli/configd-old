package application

import (
	"context"
	"fmt"

	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/pkg/models"
)

type CreateSchemaCommand struct {
	Id     *string                `json:"id"`
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
	// Name
	name, err := schema.NewName(cmd.Name)
	if err != nil {
		return nil, err
	}

	// Use id from command or generate a new one from name
	var id models.Id
	if cmd.Id != nil {
		id, err = models.NewSlug(*cmd.Id)
	} else {
		id, err = models.NewSlug(name.Value())
	}
	if err != nil {
		return nil, err
	}

	if _, err := uc.schemaRepo.FindById(ctx, id); err != schema.ErrNotFound {
		return nil, fmt.Errorf("schema with id %s already exists", id.Value())
	}

	// Parse props
	props, err := schema.PropsFromMap(cmd.Schema)
	if err != nil {
		return nil, err
	}

	s, err := schema.NewSchema(id, name, props...)
	if err != nil {
		return nil, err
	}

	if err := uc.schemaRepo.Save(ctx, s); err != nil {
		return nil, err
	}

	return &CreateSchemaResponse{
		Id:     s.Base().Id().Value(),
		Name:   s.Name().Value(),
		Schema: s.ToMap(),
	}, nil
}
