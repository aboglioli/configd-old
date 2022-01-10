package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/aboglioli/configd/domain/props"
	"github.com/mitchellh/mapstructure"
)

const (
	SCHEMA_KEY = "$schema"
)

type propSchemaInterval struct {
	Min float64 `mapstructure:"min"`
	Max float64 `mapstructure:"max"`
}

type propSchema struct {
	Type     string              `mapstructure:"type"`
	Default  interface{}         `mapstructure:"default"`
	Required bool                `mapstructure:"required"`
	Enum     []interface{}       `mapstructure:"enum"`
	Regex    string              `mapstructure:"regex"`
	Interval *propSchemaInterval `mapstructure:"interval"`
}

func PropsFromJson(data string) (map[]*props.Prop, error) {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return nil, err
	}

	return PropsFromMap(m)
}

func PropsFromMap(data map[string]interface{}) ([]*props.Prop, error) {
	if len(data) == 0 {
		return nil, errors.New("empty schema")
	}

	return parseProps(data)
}

func parseProps(m map[string]interface{}) ([]*props.Prop, error) {
	ps := make([]*props.Prop, 0)

	for k, v := range m {
		opts := make([]props.Option, 0)

		// It's an array
		if arr, ok := v.([]interface{}); ok {
			if len(arr) != 1 {
				return nil, fmt.Errorf("%s is an invalid array", k)
			}

			m, ok := arr[0].(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("%s array has invalid object", k)
			}

			v = m
			opts = append(opts, props.WithArray())
		}

		// It's an schema
		if v, ok := v.(map[string]interface{}); ok {
			// Schema
			if v, ok := v[SCHEMA_KEY]; ok {
				var s propSchema
				if err := mapstructure.Decode(v, &s); err != nil {
					return nil, err
				}

				prop, err := parseSchemaProp(k, &s, opts...)
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

			opts = append(opts, props.WithProps(subProps...))

			prop, err := props.NewObject(
				k,
				opts...,
			)
			if err != nil {
				return nil, err
			}

			ps = append(ps, prop)
		}
	}

	return ps, nil
}

func parseSchemaProp(propName string, schema *propSchema, opts ...props.Option) (*props.Prop, error) {
	// Parse options
	if schema.Default != nil {
		opts = append(opts, props.WithDefault(schema.Default))
	}

	if schema.Required {
		opts = append(opts, props.WithRequired())
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
	case props.STRING.String():
		return props.NewString(propName, opts...)
	case props.INT.String():
		return props.NewInteger(propName, opts...)
	case props.FLOAT.String():
		return props.NewFloat(propName, opts...)
	case props.BOOL.String():
		return props.NewBool(propName, opts...)
	}

	return nil, fmt.Errorf("invalid type %s", schema.Type)
}
