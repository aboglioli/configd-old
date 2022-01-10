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
