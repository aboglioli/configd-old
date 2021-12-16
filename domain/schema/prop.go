package schema

import (
	"errors"
)

type PropType string

const (
	STRING PropType = "string"
	INT    PropType = "integer"
	FLOAT  PropType = "float"
	BOOL   PropType = "bool"
)

func (t PropType) IsValid() bool {
	if t == STRING || t == INT || t == FLOAT || t == BOOL {
		return true
	}
	return false
}

type Interval struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type PropConfig struct {
	Type     PropType      `json:"type"`
	Required bool          `json:"required"`
	Values   []interface{} `json:"values"`
	Regex    string        `json:"regex"`
	Interval *Interval     `json:"interval"`
}

type Prop struct {
	Name   string
	Config *PropConfig
	Props  []*Prop
}

func NewProp(name string, config *PropConfig, props ...*Prop) (*Prop, error) {
	if len(name) == 0 {
		return nil, errors.New("empty name in prop")
	}

	if config.Type == "" {
		return nil, errors.New("missing type for prop")
	}

	if !config.Type.IsValid() {
		return nil, errors.New("unknown type")
	}

	if config.Interval != nil {
		if config.Type != INT && config.Type != FLOAT {
			return nil, errors.New("interval must be used with a numeric type")
		}

		if config.Interval.Min > config.Interval.Max {
			return nil, errors.New("invalid interval")
		}
	}

	for _, value := range config.Values {
		switch config.Type {
		case STRING:
			if _, ok := value.(string); !ok {
				return nil, errors.New("string values expected")
			}
		case INT:
			if _, ok := value.(int); !ok {
				return nil, errors.New("integer values expected")
			}
		case FLOAT:
			if _, ok := value.(float64); !ok {
				return nil, errors.New("float values expected")
			}
		case BOOL:
			if _, ok := value.(bool); !ok {
				return nil, errors.New("bool values expected")
			}
		}
	}

	prop := &Prop{
		Name:   name,
		Config: config,
		Props:  props,
	}

	return prop, nil
}
