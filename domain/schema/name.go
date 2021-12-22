package schema

import (
	"errors"

	"github.com/gosimple/slug"
)

type Name struct {
	name string
}

func NewName(n string) (*Name, error) {
	n = slug.Make(n)

	if len(n) == 0 {
		return nil, errors.New("empty schema name")
	}

	return &Name{
		name: n,
	}, nil
}
