package schema

import (
	"context"
	"errors"

	"github.com/aboglioli/configd/pkg/models"
)

var (
	ErrNotFound = errors.New("schema not found")
)

type SchemaRepository interface {
	FindById(ctx context.Context, slug models.Id) (*Schema, error)
	Save(ctx context.Context, schema *Schema) error
	Delete(ctx context.Context, slug models.Id) error
}
