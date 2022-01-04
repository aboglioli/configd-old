package schema

import (
	"testing"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/props"
	"github.com/aboglioli/configd/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidateSchema(t *testing.T) {
	type test struct {
		name string

		schema func(t *test) *Schema
		config config.ConfigData

		err bool
	}

	tests := []test{
		{
			name: "invalid simple schema",
			schema: func(t *test) *Schema {
				str, err := props.NewString("str")
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				s, err := NewSchema(n, str)

				return s
			},
			config: config.ConfigData{
				"str": 12,
			},
			err: true,
		},
		{
			name: "invalid object containing value",
			schema: func(t *test) *Schema {
				str, err := props.NewString("str")
				utils.Ok(err)

				obj, err := props.NewObject("obj", props.WithProps(str))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				s, err := NewSchema(n, obj)

				return s
			},
			config: config.ConfigData{
				"obj": map[string]interface{}{
					"str": 12,
				},
			},
			err: true,
		},
		{
			name: "invalid interval",
			schema: func(t *test) *Schema {
				integer, err := props.NewInteger("integer", props.WithInterval(4, 6))
				utils.Ok(err)

				obj, err := props.NewObject("obj", props.WithProps(integer))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				s, err := NewSchema(n, obj)

				return s
			},
			config: config.ConfigData{
				"obj": map[string]interface{}{
					"integer": 3,
				},
			},
			err: true,
		},
		{
			name: "invalid enum",
			schema: func(t *test) *Schema {
				env, err := props.NewString("env", props.WithEnum("dev", "prod"))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				s, err := NewSchema(n, env)

				return s
			},
			config: config.ConfigData{
				"env": "staging",
			},
			err: true,
		},
		{
			name: "invalid array",
			schema: func(t *test) *Schema {
				strs, err := props.NewString("strs", props.WithArray(), props.WithEnum("dev", "prod"))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				s, err := NewSchema(n, strs)

				return s
			},
			config: config.ConfigData{
				"strs": "prod",
			},
			err: true,
		},
		{
			name: "valid enum",
			schema: func(t *test) *Schema {
				env, err := props.NewString("env", props.WithEnum("dev", "prod"))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				s, err := NewSchema(n, env)

				return s
			},
			config: config.ConfigData{
				"env": "prod",
			},
			err: false,
		},
		{
			name: "valid object",
			schema: func(t *test) *Schema {
				integer, err := props.NewInteger("integer", props.WithInterval(4, 6))
				utils.Ok(err)

				str, err := props.NewString("str", props.WithRequired(true))
				utils.Ok(err)

				obj, err := props.NewObject("obj", props.WithProps(integer, str))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				s, err := NewSchema(n, obj)

				return s
			},
			config: config.ConfigData{
				"obj": map[string]interface{}{
					"integer": 5,
					"str":     "hello",
				},
			},
			err: false,
		},
		{
			name: "valid array",
			schema: func(t *test) *Schema {
				integer, err := props.NewInteger(
					"integers",
					props.WithArray(),
					props.WithInterval(2, 6),
				)
				utils.Ok(err)

				obj, err := props.NewObject("obj", props.WithProps(integer))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				s, err := NewSchema(n, obj)

				return s
			},
			config: config.ConfigData{
				"obj": map[string]interface{}{
					"integers": []interface{}{2, 3, 4, 5, 6},
				},
			},
			err: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			schema := test.schema(&test)

			err := schema.Validate(test.config)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
