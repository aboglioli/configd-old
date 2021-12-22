package schema

import (
	"errors"

	"github.com/aboglioli/configd/common/model"
	"github.com/aboglioli/configd/domain/props"
)

type Schema struct {
	id    *model.Id
	name  *Name
	props map[string]props.Prop
}

func NewSchema(name *Name, ps ...props.Prop) (*Schema, error) {
	if len(ps) == 0 {
		return nil, errors.New("schema does not have props")
	}

	psMap := make(map[string]props.Prop)
	for _, p := range ps {
		psMap[p.Name()] = p
	}

	id, err := model.GenerateId()
	if err != nil {
		return nil, err
	}

	return &Schema{
		id:    id,
		name:  name,
		props: psMap,
	}, nil
}

func (s *Schema) Id() *model.Id {
	return s.id
}

func (s *Schema) Name() *Name {
	return s.name
}

func (s *Schema) Props() map[string]props.Prop {
	return s.props
}

func (s *Schema) Validate(c map[string]interface{}) error {
	return nil
}
