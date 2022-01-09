package infrastructure

import (
	"sync"

	"github.com/aboglioli/configd/common/models"
	"github.com/aboglioli/configd/domain/config"
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

func (r *inMemConfigRepository) FindById(slug *models.Slug) (*config.Config, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if s, ok := r.configs[slug.Value()]; ok {
		return s, nil
	}

	return nil, config.ErrNotFound
}

func (r *inMemConfigRepository) FindBySchemaId(schemaSlug *models.Slug) ([]*config.Config, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	found := make([]*config.Config, 0)

	for _, c := range r.configs {
		if c.SchemaSlug().Equals(schemaSlug) {
			found = append(found, c)
		}
	}

	return found, nil
}

func (r *inMemConfigRepository) Save(config *config.Config) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.configs[config.Slug().Value()] = config

	return nil
}

func (r *inMemConfigRepository) Delete(slug *models.Slug) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.configs, slug.Value())

	return nil
}
