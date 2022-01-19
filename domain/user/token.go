package user

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

const (
	JWT_SECRET = "my-secret"
)

type TokenData = jwt.MapClaims

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
	if data == nil {
		return Token{}, fmt.Errorf("nil token data")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	tokenStr, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return Token{}, err
	}

	return NewToken(tokenStr)
}

func (t Token) ParseData() (TokenData, error) {
	token, err := jwt.Parse(t.token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(TokenData)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token %s is invalid", t.token)
	}

	return claims, nil
}

func (t Token) Value() string {
	return t.token
}
