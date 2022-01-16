package schema

import (
	"errors"
	"fmt"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/props"
	"github.com/aboglioli/configd/pkg/events"
	"github.com/aboglioli/configd/pkg/models"
)

type Schema struct {
	agg *models.AggregateRoot

	name  Name
	props map[string]*props.Prop
}

func BuildSchema(id models.Id, name Name, ps ...*props.Prop) (*Schema, error) {
	if len(ps) == 0 {
		return nil, errors.New("schema does not have props")
	}

	psMap := make(map[string]*props.Prop)
	for _, p := range ps {
		psMap[p.Name()] = p
	}

	agg, err := models.NewAggregateRoot(id)
	if err != nil {
		return nil, err
	}

	return &Schema{
		agg:   agg,
		name:  name,
		props: psMap,
	}, nil
}

func NewSchema(id models.Id, name Name, ps ...*props.Prop) (*Schema, error) {
	s, err := BuildSchema(id, name, ps...)
	if err != nil {
		return nil, err
	}

	event, err := events.NewEvent(
		s.agg.Id().Value(),
		CreatedTopic,
		Created{
			Id:    s.agg.Id().Value(),
			Name:  s.name.Value(),
			Props: s.ToMap(),
		},
	)
	if err != nil {
		return nil, err
	}

	s.agg.RecordEvent(event)

	return s, nil
}

func (s *Schema) Base() models.ReadOnlyAggregateRoot {
	return s.agg
}

func (s *Schema) Name() Name {
	return s.name
}

func (s *Schema) ChangeName(name Name) error {
	s.name = name
	s.agg.Update()

	event, err := events.NewEvent(
		s.agg.Id().Value(),
		NameChangedTopic,
		NameChanged{
			Id:   s.agg.Id().Value(),
			Name: s.name.Value(),
		},
	)
	if err != nil {
		return err
	}

	s.agg.RecordEvent(event)

	return nil
}

func (s *Schema) Props() map[string]*props.Prop {
	return s.props
}

func (s *Schema) ChangeProps(ps ...*props.Prop) error {
	psMap := make(map[string]*props.Prop)
	for _, p := range ps {
		psMap[p.Name()] = p
	}

	s.props = psMap
	s.agg.Update()

	event, err := events.NewEvent(
		s.agg.Id().Value(),
		PropsChangedTopic,
		PropsChanged{
			Id:    s.agg.Id().Value(),
			Props: s.ToMap(),
		},
	)
	if err != nil {
		return err
	}

	s.agg.RecordEvent(event)

	return nil
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
