package schema

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aboglioli/configd/common/model"
	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/props"
)

type Schema struct {
	slug  *model.Slug
	name  *Name
	props map[string]props.Prop
}

func BuildSchema(slug *model.Slug, name *Name, ps ...props.Prop) (*Schema, error) {
	if len(ps) == 0 {
		return nil, errors.New("schema does not have props")
	}

	psMap := make(map[string]props.Prop)
	for _, p := range ps {
		psMap[p.Name()] = p
	}

	return &Schema{
		slug:  slug,
		name:  name,
		props: psMap,
	}, nil
}

func NewSchema(name *Name, ps ...props.Prop) (*Schema, error) {
	slug, err := model.NewSlug(name.Value())
	if err != nil {
		return nil, err
	}

	return BuildSchema(slug, name, ps...)
}

func (s *Schema) Slug() *model.Slug {
	return s.slug
}

func (s *Schema) Name() *Name {
	return s.name
}

func (s *Schema) Props() map[string]props.Prop {
	return s.props
}

func (s *Schema) Validate(c config.ConfigData) error {
	for k, v := range c {
		prop, ok := s.props[k]
		if !ok {
			return fmt.Errorf("prop %s not found in schema", k)
		}

		switch v.(type) {
		case string:
			if prop.Type() != props.STRING {
				return fmt.Errorf("%v is not a string", v)
			}
		case int:
			if prop.Type() != props.INT {
				return fmt.Errorf("%v is not an integer", v)
			}
		case int32:
			if prop.Type() != props.INT {
				return fmt.Errorf("%v is not an integer", v)
			}
		case int64:
			if prop.Type() != props.INT {
				return fmt.Errorf("%v is not an integer", v)
			}
		case float32:
			if prop.Type() != props.FLOAT {
				return fmt.Errorf("%v is not a float", v)
			}
		case float64:
			if prop.Type() != props.FLOAT {
				return fmt.Errorf("%v is not a float", v)
			}
		case bool:
			if prop.Type() != props.BOOL {
				return fmt.Errorf("%v is not a boolean", v)
			}
		case interface{}:
			if prop.Type() != props.OBJECT {
				return fmt.Errorf("%v is not an object", v)
			}
		}
	}

	return nil
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	d := map[string]interface{}{
		"slug":  s.slug.Value(),
		"name":  s.name.Value(),
		"props": s.props,
	}

	return json.Marshal(&d)
}
