package schema

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aboglioli/configd/domain/props"
	"github.com/mitchellh/mapstructure"
)

const (
	SCHEMA_KEY = "$schema"
)

type SchemaInterval struct {
	Min float64 `mapstructure:"min"`
	Max float64 `mapstructure:"max"`
}

type schema struct {
	Type     string          `mapstructure:"type"`
	Default  interface{}     `mapstructure:"default"`
	Required bool            `mapstructure:"required"`
	Enum     []interface{}   `mapstructure:"enum"`
	Regex    string          `mapstructure:"regex"`
	Interval *SchemaInterval `mapstructure:"interval"`
}

func FromJson(name string, data string) (*Schema, error) {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return nil, err
	}

	schema, err := FromMap(name, m)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

func FromMap(name string, data map[string]interface{}) (*Schema, error) {
	props, err := parseProps(data)
	if err != nil {
		return nil, err
	}

	n, err := NewName(name)
	if err != nil {
		return nil, err
	}

	s, err := NewSchema(n, props...)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func parseProps(m map[string]interface{}) ([]props.Prop, error) {
	ps := make([]props.Prop, 0)

	for k, v := range m {
		// It's an schema
		if v, ok := v.(map[string]interface{}); ok {
			// Schema
			if v, ok := v[SCHEMA_KEY]; ok {
				var s schema
				if err := mapstructure.Decode(v, &s); err != nil {
					return nil, err
				}

				prop, err := parseSchemaProp(k, &s)
				if err != nil {
					return nil, err
				}

				ps = append(ps, prop)
				continue
			}

			// Object
			subProps, err := parseProps(v)
			if err != nil {
				return nil, err
			}

			prop, err := props.NewObject(
				k,
				props.WithProps(subProps...),
			)
			if err != nil {
				return nil, err
			}

			ps = append(ps, prop)
		}
	}

	return ps, nil
}

func parseSchemaProp(propName string, schema *schema) (props.Prop, error) {
	// Parse options
	opts := make([]props.Option, 0)
	opts = append(opts, props.WithRequired(schema.Required))

	if schema.Default != nil {
		opts = append(opts, props.WithDefault(schema.Default))
	}

	if schema.Enum != nil {
		opts = append(opts, props.WithEnum(schema.Enum...))
	}

	if schema.Regex != "" {
		opts = append(opts, props.WithRegex(schema.Regex))
	}

	if schema.Interval != nil {
		opts = append(opts, props.WithInterval(schema.Interval.Min, schema.Interval.Max))
	}

	switch strings.ToLower(schema.Type) {
	case string(props.STRING):
		return props.NewString(propName, opts...)
	case props.INT:
		return props.NewInteger(propName, opts...)
	case props.FLOAT:
		return props.NewFloat(propName, opts...)
	case props.BOOL:
		return props.NewBool(propName, opts...)
	}

	return nil, fmt.Errorf("invalid type %s", schema.Type)
}
