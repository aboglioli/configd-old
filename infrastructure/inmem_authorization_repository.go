package infrastructure

import (
	"context"
	"sync"

	"github.com/aboglioli/configd/domain/security"
)

type InMemAuthorizationRepository struct {
	mux            sync.Mutex
	authorizations map[string]*security.Authorization
}

func NewInMemAuthorizationRepository() *InMemAuthorizationRepository {
	return &InMemAuthorizationRepository{
		authorizations: make(map[string]*security.Authorization),
	}
}

func (r *InMemAuthorizationRepository) FindByApiKey(
	ctx context.Context,
	hashedApiKey security.HashedApiKey,
) (*security.Authorization, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if s, ok := r.authorizations[hashedApiKey.Value()]; ok {
		return s, nil
	}

	return nil, security.ErrNotFound
}

func (r *InMemAuthorizationRepository) Save(ctx context.Context, authorization *security.Authorization) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.authorizations[authorization.HashedApiKey().Value()] = authorization

	return nil
}

func (r *InMemAuthorizationRepository) Delete(ctx context.Context, hashedApiKey security.HashedApiKey) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.authorizations, hashedApiKey.Value())

	return nil
}
