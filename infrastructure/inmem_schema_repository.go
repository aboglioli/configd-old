package infrastructure

import (
	"context"
	"sync"

	"github.com/aboglioli/configd/domain/schema"
	"github.com/aboglioli/configd/pkg/models"
)

var _ schema.SchemaRepository = (*InMemSchemaRepository)(nil)

type InMemSchemaRepository struct {
	mux     sync.Mutex
	schemas map[string]*schema.Schema
}

func NewInMemSchemaRepository() *InMemSchemaRepository {
	return &InMemSchemaRepository{
		schemas: make(map[string]*schema.Schema),
	}
}

func (r *InMemSchemaRepository) FindById(ctx context.Context, id models.Id) (*schema.Schema, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if s, ok := r.schemas[id.Value()]; ok {
		return s, nil
	}

	return nil, schema.ErrNotFound
}

func (r *InMemSchemaRepository) Save(ctx context.Context, schema *schema.Schema) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.schemas[schema.Base().Id().Value()] = schema

	return nil
}

func (r *InMemSchemaRepository) Delete(ctx context.Context, id models.Id) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.schemas, id.Value())

	return nil
}
