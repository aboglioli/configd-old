package user

import (
	"fmt"
)

type Username struct {
	username string
}

func NewUsername(username string) (Username, error) {
	if len(username) < 4 {
		return Username{}, fmt.Errorf("username %s too short", username)
	}

	return Username{
		username: username,
	}, nil
}

func (u Username) Value() string {
	return u.username
}

func (u Username) Equals(o Username) bool {
	return u.username == o.username
}
