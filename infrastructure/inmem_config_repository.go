package infrastructure

import (
	"context"
	"sync"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/pkg/models"
)

var _ config.ConfigRepository = (*InMemConfigRepository)(nil)

type InMemConfigRepository struct {
	mux     sync.Mutex
	configs map[string]*config.Config
}

func NewInMemConfigRepository() *InMemConfigRepository {
	return &InMemConfigRepository{
		configs: make(map[string]*config.Config),
	}
}

func (r *InMemConfigRepository) FindById(ctx context.Context, id models.Id) (*config.Config, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if s, ok := r.configs[id.Value()]; ok {
		return s, nil
	}

	return nil, config.ErrNotFound
}

func (r *InMemConfigRepository) FindBySchemaId(ctx context.Context, schemaId models.Id) ([]*config.Config, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	found := make([]*config.Config, 0)

	for _, c := range r.configs {
		if c.SchemaId().Equals(schemaId) {
			found = append(found, c)
		}
	}

	return found, nil
}

func (r *InMemConfigRepository) Save(ctx context.Context, config *config.Config) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.configs[config.Base().Id().Value()] = config

	return nil
}

func (r *InMemConfigRepository) Delete(ctx context.Context, id models.Id) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.configs, id.Value())

	return nil
}
