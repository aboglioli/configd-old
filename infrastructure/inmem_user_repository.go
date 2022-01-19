package infrastructure

import (
	"context"
	"sync"

	"github.com/aboglioli/configd/domain/user"
)

var _ user.UserRepository = (*InMemUserRepository)(nil)

type InMemUserRepository struct {
	mux   sync.Mutex
	users map[string]*user.User
}

func NewInMemUserRepository() *InMemUserRepository {
	return &InMemUserRepository{
		users: make(map[string]*user.User),
	}
}

func (r *InMemUserRepository) FindByUsername(ctx context.Context, username user.Username) (*user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if u, ok := r.users[username.Value()]; ok {
		return u, nil
	}

	return nil, user.ErrNotFound
}

func (r *InMemUserRepository) Save(ctx context.Context, user *user.User) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.users[user.Username().Value()] = user

	return nil
}

func (r *InMemUserRepository) Delete(ctx context.Context, username user.Username) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.users, username.Value())

	return nil
}
