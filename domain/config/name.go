package config

import (
	"errors"

	"github.com/gosimple/slug"
)

type Name struct {
	name string
}

func NewName(n string) (*Name, error) {
	n = slug.Make(n)

	if len(n) < 4 {
		return nil, errors.New("config name too short")
	}

	return &Name{
		name: n,
	}, nil
}
