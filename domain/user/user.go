package user

import (
	"errors"
	"time"
)

var (
	ErrInvalidLogin = errors.New("invalid username or password")
)

type Access string

const (
	READ_ONLY_ACCESS Access = "read_only"
	FULL_ACCESS      Access = "full_access"
)

type User struct {
	username       Username
	hashedPassword HashedPassword
	access         Access
}

func BuildUser(
	username Username,
	hashedPassword HashedPassword,
	access Access,
) (*User, error) {
	return &User{
		username:       username,
		hashedPassword: hashedPassword,
		access:         access,
	}, nil
}

func NewUser(
	username Username,
	password Password,
	access Access,
) (*User, error) {
	hashedPassword, err := password.Hash()
	if err != nil {
		return nil, err
	}

	return BuildUser(username, hashedPassword, access)
}

func (u *User) Username() Username {
	return u.username
}

func (u *User) HashedPassword() HashedPassword {
	return u.hashedPassword
}

func (u *User) Access() Access {
	return u.access
}

func (u *User) Login(username Username, password Password) (Token, error) {
	if u.username.Equals(username) && u.hashedPassword.Validate(password) {
		return GenerateToken(TokenData{
			"username":  u.username.Value(),
			"timestamp": time.Now(),
		})
	}

	return Token{}, ErrInvalidLogin
}
