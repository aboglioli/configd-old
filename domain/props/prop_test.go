package props

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ok(err error) {
	if err != nil {
		panic(err)
	}
}

func propsToInterface(props ...*prop) []Prop {
	ps := make([]Prop, len(props))
	for i, p := range props {
		ps[i] = p
	}
	return ps
}

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
			name: "different prop type and values types",
			makeProp: func() (*prop, error) {
				return NewString("env", WithValues(1, 2, 3))
			},
			err: true,
		},
		{
			name: "different prop type and values types",
			makeProp: func() (*prop, error) {
				return NewInteger("env", WithValues("str"))
			},
			err: true,
		},
		{
			name: "different prop type and values types",
			makeProp: func() (*prop, error) {
				return NewString("env", WithValues("str", 256))
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
			name: "object with values",
			makeProp: func() (*prop, error) {
				return NewObject("env", WithValues(1, 2, 3))
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
				ok(err)

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
					WithValues("str1", "str2"),
				)
				ok(err)

				num, err := NewInteger(
					"num",
					WithValues(1, 5, 10),
					WithInterval(0, 10),
				)
				ok(err)

				return NewObject("obj", WithProps(str, num))
			},
			expected: &prop{
				t:        OBJECT,
				name:     "obj",
				required: true,
				props: propsToInterface(
					&prop{
						t:        STRING,
						name:     "str",
						required: true,
						def:      "default",
						values:   []interface{}{"str1", "str2"},
					},
					&prop{
						t:        INT,
						name:     "num",
						values:   []interface{}{1, 5, 10},
						interval: &interval{0, 10},
					},
				),
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
