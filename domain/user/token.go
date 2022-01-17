package user

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

const (
	JWT_SECRET = "my-secret"
)

type TokenData map[string]interface{}

func (td TokenData) Valid() error {
	return nil
}

type Token struct {
	token string
}

func NewToken(token string) (Token, error) {
	if len(token) < 10 {
		return Token{}, fmt.Errorf("token %s invalid", token)
	}

	return Token{
		token: token,
	}, nil
}

func GenerateToken(data TokenData) (Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	tokenStr, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return Token{}, err
	}

	return NewToken(tokenStr)
}
