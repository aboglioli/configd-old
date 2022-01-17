package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type HashedPassword struct {
	hashedPassword string
}

func NewHashedPassword(hashedPassword string) (HashedPassword, error) {
	if len(hashedPassword) < 30 {
		return HashedPassword{}, fmt.Errorf("password hash %s is invalid", hashedPassword)
	}

	return HashedPassword{
		hashedPassword: hashedPassword,
	}, nil
}

func (p HashedPassword) Value() string {
	return p.hashedPassword
}

func (p HashedPassword) Equals(o HashedPassword) bool {
	return p.hashedPassword == o.hashedPassword
}

func (p HashedPassword) Validate(pwd Password) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p.hashedPassword), []byte(pwd.password)); err != nil {
		return false
	}

	return true
}
