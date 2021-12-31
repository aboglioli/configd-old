package schema

import (
	"errors"
)

type Name struct {
	name string
}

func NewName(n string) (*Name, error) {
	if len(n) == 0 {
		return nil, errors.New("empty schema name")
	}

	return &Name{
		name: n,
	}, nil
}

func (n *Name) Value() string {
	return n.name
}

func (n *Name) Equals(o *Name) bool {
	return n.name == o.name
}
