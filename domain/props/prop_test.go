package props

import (
	"testing"

	"github.com/aboglioli/configd/utils"
	"github.com/stretchr/testify/assert"
)

func TestBuildProps(t *testing.T) {
	tests := []struct {
		name     string
		makeProp func() (*prop, error)
		expected *prop
		err      bool
	}{
		{
			name: "empty name",
			makeProp: func() (*prop, error) {
				return NewString("")
			},
			err: true,
		},
		{
			name: "simple string",
			makeProp: func() (*prop, error) {
				return NewString("env")
			},
			expected: &prop{
				t:    STRING,
				name: "env",
			},
		},
		{
			name: "different prop type and enum types",
			makeProp: func() (*prop, error) {
				return NewString("env", WithEnum(1, 2, 3))
			},
			err: true,
		},
		{
			name: "different prop type and enum types",
			makeProp: func() (*prop, error) {
				return NewInteger("env", WithEnum("str"))
			},
			err: true,
		},
		{
			name: "different prop type and enum types",
			makeProp: func() (*prop, error) {
				return NewString("env", WithEnum("str", 256))
			},
			err: true,
		},
		{
			name: "different prop type and default type",
			makeProp: func() (*prop, error) {
				return NewString("env", WithDefault(12.55))
			},
			err: true,
		},
		{
			name: "different prop type and default type",
			makeProp: func() (*prop, error) {
				return NewInteger("env", WithDefault("str"))
			},
			err: true,
		},
		{
			name: "object not required",
			makeProp: func() (*prop, error) {
				return NewObject("env", WithRequired(false))
			},
			err: true,
		},
		{
			name: "object with default",
			makeProp: func() (*prop, error) {
				return NewObject("env", WithDefault("hello"))
			},
			err: true,
		},
		{
			name: "object with enum",
			makeProp: func() (*prop, error) {
				return NewObject("env", WithEnum(1, 2, 3))
			},
			err: true,
		},
		{
			name: "non-string with regex",
			makeProp: func() (*prop, error) {
				return NewInteger("env", WithRegex("[0-9]*"))
			},
			err: true,
		},
		{
			name: "non-numeric with interval",
			makeProp: func() (*prop, error) {
				return NewString("env", WithInterval(1, 10))
			},
			err: true,
		},
		{
			name: "non-object with props",
			makeProp: func() (*prop, error) {
				prop, err := NewString("str")
				utils.Ok(err)

				return NewString("env", WithProps(prop))
			},
			err: true,
		},
		{
			name: "object with props",
			makeProp: func() (*prop, error) {
				str, err := NewString(
					"str",
					WithRequired(true),
					WithDefault("default"),
					WithEnum("str1", "str2"),
				)
				utils.Ok(err)

				num, err := NewInteger(
					"num",
					WithEnum(1, 5, 10),
					WithInterval(0, 10),
				)
				utils.Ok(err)

				return NewObject("obj", WithProps(str, num))
			},
			expected: &prop{
				t:        OBJECT,
				name:     "obj",
				required: true,
				props: map[string]Prop{
					"str": &prop{
						t:        STRING,
						name:     "str",
						required: true,
						def:      "default",
						enum:     []interface{}{"str1", "str2"},
					},
					"num": &prop{
						t:        INT,
						name:     "num",
						enum:     []interface{}{1, 5, 10},
						interval: &interval{0, 10},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prop, err := test.makeProp()

			assert.Equal(t, test.expected, prop)
			if test.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
