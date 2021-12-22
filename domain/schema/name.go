package schema

import (
	"errors"

	"github.com/gosimple/slug"
)

type Name struct {
	name string
}

func NewName(n string) (*Name, error) {
	if len(n) < 4 {
		return nil, errors.New("schema name too short")
	}

	return &Name{
		name: slug.Make(n),
	}, nil
}
