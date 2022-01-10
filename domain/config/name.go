package config

import (
	"errors"
)

type Name struct {
	name string
}

func NewName(n string) (Name, error) {
	if len(n) < 4 {
		return Name{}, errors.New("config name too short")
	}

	return Name{
		name: n,
	}, nil
}

func (n Name) Value() string {
	return n.name
}

func (n Name) Equals(o Name) bool {
	return n.name == o.name
}
