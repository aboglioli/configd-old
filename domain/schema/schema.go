package schema

import (
	"errors"
	"fmt"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/props"
	"github.com/aboglioli/configd/pkg/models"
)

type Schema struct {
	agg *models.AggregateRoot

	name  Name
	props map[string]*props.Prop
}

func BuildSchema(slug models.Id, name Name, ps ...*props.Prop) (*Schema, error) {
	if len(ps) == 0 {
		return nil, errors.New("schema does not have props")
	}

	psMap := make(map[string]*props.Prop)
	for _, p := range ps {
		psMap[p.Name()] = p
	}

	agg, err := models.NewAggregateRoot(slug)
	if err != nil {
		return nil, err
	}

	return &Schema{
		agg:   agg,
		name:  name,
		props: psMap,
	}, nil
}

func NewSchema(name Name, ps ...*props.Prop) (*Schema, error) {
	slug, err := models.NewSlug(name.Value())
	if err != nil {
		return nil, err
	}

	return BuildSchema(slug, name, ps...)
}

func (s *Schema) Base() models.PublicAggregateRoot {
	return s.agg
}

func (s *Schema) Name() Name {
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

func (s *Schema) ToMap() map[string]interface{} {
	return propsToMap(s.props)
}

func propsToMap(ps map[string]*props.Prop) map[string]interface{} {
	m := make(map[string]interface{})

	for k, p := range ps {
		var s map[string]interface{}

		if p.Type() == props.OBJECT {
			s = propsToMap(p.Props())
		} else {
			// Basic prop types
			var interval map[string]interface{}
			if p.Interval() != nil {
				interval = map[string]interface{}{
					"min": p.Interval().Min(),
					"max": p.Interval().Max(),
				}
			}

			s = map[string]interface{}{
				SCHEMA_KEY: map[string]interface{}{
					"type":     p.Type(),
					"default":  p.Default(),
					"required": p.IsRequired(),
					"enum":     p.Enum(),
					"regex":    p.Regex(),
					"interval": interval,
				},
			}
		}

		if p.IsArray() {
			m[k] = []interface{}{s}
		} else {
			m[k] = s
		}
	}

	return m
}
