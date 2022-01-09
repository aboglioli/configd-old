package infrastructure

import (
	"sync"

	"github.com/aboglioli/configd/common/models"
	"github.com/aboglioli/configd/domain/schema"
)

var _ schema.SchemaRepository = (*inMemSchemaRepository)(nil)

type inMemSchemaRepository struct {
	mux     sync.Mutex
	schemas map[string]*schema.Schema
}

func NewInMemSchemaRepository() *inMemSchemaRepository {
	return &inMemSchemaRepository{
		schemas: make(map[string]*schema.Schema),
	}
}

func (r *inMemSchemaRepository) FindById(slug *models.Slug) (*schema.Schema, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if s, ok := r.schemas[slug.Value()]; ok {
		return s, nil
	}

	return nil, schema.ErrNotFound
}

func (r *inMemSchemaRepository) Save(schema *schema.Schema) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.schemas[schema.Slug().Value()] = schema

	return nil
}

func (r *inMemSchemaRepository) Delete(slug *models.Slug) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.schemas, slug.Value())

	return nil
}
