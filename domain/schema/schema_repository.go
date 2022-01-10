package schema

import (
	"errors"

	"github.com/aboglioli/configd/pkg/models"
)

var (
	ErrNotFound = errors.New("schema not found")
)

type SchemaRepository interface {
	FindById(slug models.Id) (*Schema, error)
	Save(schema *Schema) error
	Delete(slug models.Id) error
}
