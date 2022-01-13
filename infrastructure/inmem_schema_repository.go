package infrastructure

import (
	"context"
	"sync"

	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/pkg/models"
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

func (r *inMemSchemaRepository) FindById(ctx context.Context, slug models.Id) (*schema.Schema, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if s, ok := r.schemas[slug.Value()]; ok {
		return s, nil
	}

	return nil, schema.ErrNotFound
}

func (r *inMemSchemaRepository) Save(ctx context.Context, schema *schema.Schema) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.schemas[schema.Base().Id().Value()] = schema

	return nil
}

func (r *inMemSchemaRepository) Delete(ctx context.Context, slug models.Id) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.schemas, slug.Value())

	return nil
}
