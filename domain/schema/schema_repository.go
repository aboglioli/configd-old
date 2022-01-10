package schema

import (
	"errors"

	"github.com/aboglioli/configd/pkg/models"
)

var (
	ErrNotFound = errors.New("schema not found")
)

type SchemaRepository interface {
	FindById(slug *models.Slug) (*Schema, error)
	Save(schema *Schema) error
	Delete(slug *models.Slug) error
}
