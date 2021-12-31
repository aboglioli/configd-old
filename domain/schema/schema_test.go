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

		props  func() []props.Prop
		config config.ConfigData

		err bool
	}

	tests := []test{
		{
			name: "invalid simple schema",
			props: func() []props.Prop {
				s, err := props.NewString("str")
				utils.Ok(err)

				return []props.Prop{s}
			},
			config: config.ConfigData{
				"str": 12,
			},
			err: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			name, err := NewName(test.name)
			assert.NoError(t, err)

			schema, err := NewSchema(name, test.props()...)
			assert.NoError(t, err)

			err = schema.Validate(test.config)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
