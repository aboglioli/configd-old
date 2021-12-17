package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/aboglioli/configd/domain/props"
	"github.com/aboglioli/configd/domain/schema"
	"github.com/mitchellh/mapstructure"
)

const (
	SCHEMA_KEY = "$schema"
)

type SchemaInterval struct {
	Min float64 `mapstructure:"min"`
	Max float64 `mapstructure:"max"`
}

type Schema struct {
	Type     string          `mapstructure:"type"`
	Default  interface{}     `mapstructure:"default"`
	Required bool            `mapstructure:"required"`
	Values   []interface{}   `mapstructure:"values"`
	Regex    string          `mapstructure:"regex"`
	Interval *SchemaInterval `mapstructure:"interval"`
}

func FromJson(data string) (schema.Schema, error) {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return nil, err
	}

	props, err := parseProps(m)
	if err != nil {
		return nil, err
	}

	s, err := schema.NewSchema(props...)
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
				var s Schema
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

func parseSchemaProp(propName string, schema *Schema) (props.Prop, error) {
	// Parse options
	opts := make([]props.Option, 0)
	opts = append(opts, props.WithRequired(schema.Required))

	if schema.Default != nil {
		opts = append(opts, props.WithDefault(schema.Default))
	}

	if schema.Values != nil {
		opts = append(opts, props.WithValues(schema.Values...))
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

	return nil, errors.New(fmt.Sprintf("invalid type %s", schema.Type))
}
