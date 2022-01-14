package dependencies

import (
	"sync"

	"github.com/aboglioli/configd/infrastructure"
)

var once sync.Once
var deps *Dependencies

type Dependencies struct {
	SchemaRepository *infrastructure.InMemSchemaRepository
	ConfigRepository *infrastructure.InMemConfigRepository
}

func Get() *Dependencies {
	once.Do(func() {
		deps = &Dependencies{
			SchemaRepository: infrastructure.NewInMemSchemaRepository(),
			ConfigRepository: infrastructure.NewInMemConfigRepository(),
		}
	})

	return deps
}
