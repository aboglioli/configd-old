package schema

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aboglioli/configd/common/models"
	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/props"
)

type Schema struct {
	slug  *models.Slug
	name  *Name
	props map[string]*props.Prop
}

func BuildSchema(slug *models.Slug, name *Name, ps ...*props.Prop) (*Schema, error) {
	if len(ps) == 0 {
		return nil, errors.New("schema does not have props")
	}

	psMap := make(map[string]*props.Prop)
	for _, p := range ps {
		psMap[p.Name()] = p
	}

	return &Schema{
		slug:  slug,
		name:  name,
		props: psMap,
	}, nil
}

func NewSchema(name *Name, ps ...*props.Prop) (*Schema, error) {
	slug, err := models.NewSlug(name.Value())
	if err != nil {
		return nil, err
	}

	return BuildSchema(slug, name, ps...)
}

func (s *Schema) Slug() *models.Slug {
	return s.slug
}

func (s *Schema) Name() *Name {
	return s.name
}

func (s *Schema) Props() map[string]*props.Prop {
	return s.props
}

func (s *Schema) Validate(c config.ConfigData) error {
	for k, p := range s.props {
		entry, ok := c[k]
		if !ok {
			return fmt.Errorf("prop %s not found in config", k)
		}

		if err := p.Validate(entry); err != nil {
			return fmt.Errorf("path %s: %s", k, err.Error())
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
