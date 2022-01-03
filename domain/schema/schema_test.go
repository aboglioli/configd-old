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
			name: "invalid object",
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
