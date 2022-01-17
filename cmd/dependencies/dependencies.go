package dependencies

import (
	"sync"

	"github.com/aboglioli/configd/infrastructure"
)

var once sync.Once
var deps *Dependencies

type Dependencies struct {
	EventBus                *infrastructure.InMemEventBus
	SchemaRepository        *infrastructure.InMemSchemaRepository
	ConfigRepository        *infrastructure.InMemConfigRepository
	AuthorizationRepository *infrastructure.InMemAuthorizationRepository
}

func Get() *Dependencies {
	once.Do(func() {
		deps = &Dependencies{
			EventBus:                infrastructure.NewInMemEventBus(),
			SchemaRepository:        infrastructure.NewInMemSchemaRepository(),
			ConfigRepository:        infrastructure.NewInMemConfigRepository(),
			AuthorizationRepository: infrastructure.NewInMemAuthorizationRepository(),
		}
	})

	return deps
}
