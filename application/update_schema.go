package application

import (
	"context"

	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/pkg/models"
)

type UpdateSchemaCommand struct {
	Id     string                  `json:"id"`
	Name   *string                 `json:"name"`
	Schema *map[string]interface{} `json:"schema"`
}

type UpdateSchemaResponse struct {
	Id     string                 `json:"id"`
	Name   string                 `json:"name"`
	Schema map[string]interface{} `json:"schema"`
}

type UpdateSchema struct {
	schemaRepo schema.SchemaRepository
}

func NewUpdateSchema(
	schemaRepo schema.SchemaRepository,
) *UpdateSchema {
	return &UpdateSchema{
		schemaRepo: schemaRepo,
	}
}

func (uc *UpdateSchema) Exec(
	ctx context.Context,
	cmd *UpdateSchemaCommand,
) (*UpdateSchemaResponse, error) {
	id, err := models.BuildId(cmd.Id)
	if err != nil {
		return nil, err
	}

	s, err := uc.schemaRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Name
	if cmd.Name != nil {
		name, err := schema.NewName(*cmd.Name)
		if err != nil {
			return nil, err
		}

		s.ChangeName(name)
	}

	if cmd.Schema != nil {
		props, err := schema.PropsFromMap(*cmd.Schema)
		if err != nil {
			return nil, err
		}

		s.ChangeProps(props...)
	}

	if err := uc.schemaRepo.Save(ctx, s); err != nil {
		return nil, err
	}

	return &UpdateSchemaResponse{
		Id:     s.Base().Id().Value(),
		Name:   s.Name().Value(),
		Schema: s.ToMap(),
	}, nil
}
