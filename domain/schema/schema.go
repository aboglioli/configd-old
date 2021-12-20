package schema

import (
	"errors"

	"github.com/aboglioli/configd/domain/props"
)

type Schema struct {
	name  string
	props map[string]props.Prop
}

func NewSchema(name string, ps ...props.Prop) (*Schema, error) {
	if len(ps) == 0 {
		return nil, errors.New("schema does not have props")
	}

	psMap := make(map[string]props.Prop)
	for _, p := range ps {
		psMap[p.Name()] = p
	}

	return &Schema{
		name:  name,
		props: psMap,
	}, nil
}

func (s *Schema) Name() string {
	return s.name
}

func (s *Schema) Props() map[string]props.Prop {
	return s.props
}
