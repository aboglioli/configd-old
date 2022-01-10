package infrastructure

import (
	"sync"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/pkg/models"
)

var _ config.ConfigRepository = (*inMemConfigRepository)(nil)

type inMemConfigRepository struct {
	mux     sync.Mutex
	configs map[string]*config.Config
}

func NewInMemConfigRepository() *inMemConfigRepository {
	return &inMemConfigRepository{
		configs: make(map[string]*config.Config),
	}
}

func (r *inMemConfigRepository) FindById(slug models.Id) (*config.Config, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if s, ok := r.configs[slug.Value()]; ok {
		return s, nil
	}

	return nil, config.ErrNotFound
}

func (r *inMemConfigRepository) FindBySchemaId(schemaId models.Id) ([]*config.Config, error) {
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

func (r *inMemConfigRepository) Save(config *config.Config) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.configs[config.Base().Id().Value()] = config

	return nil
}

func (r *inMemConfigRepository) Delete(slug models.Id) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.configs, slug.Value())

	return nil
}
