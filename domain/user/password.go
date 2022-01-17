package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	PASSWORD_HASH_COST = bcrypt.DefaultCost
)

type Password struct {
	password string
}

func NewPassword(password string) (Password, error) {
	if len(password) < 8 {
		return Password{}, fmt.Errorf("password %s is too weak", password)
	}

	return Password{
		password: password,
	}, nil
}

func (p Password) Value() string {
	return p.password
}

func (p Password) Equals(o Password) bool {
	return p.password == o.password
}

func (p Password) Hash() (HashedPassword, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p.password), PASSWORD_HASH_COST)
	if err != nil {
		return HashedPassword{}, err
	}

	return NewHashedPassword(string(b))
}
