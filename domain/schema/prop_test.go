package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProp(t *testing.T) {
	type args struct {
		name   string
		config *PropConfig
		props  []*Prop
	}

	tests := []struct {
		name     string
		args     args
		expected *Prop
		err      bool
	}{
		{
			name: "without type",
			args: args{
				name: "env",
				config: &PropConfig{
					Required: true,
					Values:   []interface{}{"dev", "staging", "prod"},
					Regex:    "^[a-z0-9]{3,6}$",
				},
			},
			err: true,
		},
		{
			name: "empty name",
			args: args{
				name: "",
				config: &PropConfig{
					Type:     "string",
					Required: true,
					Values:   []interface{}{"dev", "staging", "prod"},
					Regex:    "^[a-z0-9]{3,6}$",
				},
			},
			err: true,
		},
		{
			name: "unknown type",
			args: args{
				name: "env",
				config: &PropConfig{
					Type: "other",
				},
			},
			err: true,
		},
		{
			name: "interval with string type",
			args: args{
				name: "env",
				config: &PropConfig{
					Type:     "string",
					Interval: &Interval{1, 10},
				},
			},
			err: true,
		},
		{
			name: "invalid interval",
			args: args{
				name: "env",
				config: &PropConfig{
					Type:     "float",
					Interval: &Interval{10, 1},
				},
			},
			err: true,
		},
		{
			name: "string type with integer values",
			args: args{
				name: "env",
				config: &PropConfig{
					Type:   "string",
					Values: []interface{}{1, 2, 3},
				},
			},
			err: true,
		},
		{
			name: "float type with integer values",
			args: args{
				name: "threshold",
				config: &PropConfig{
					Type:   "float",
					Values: []interface{}{1, 2, 3},
				},
			},
			err: true,
		},
		{
			name: "string type with bool values",
			args: args{
				name: "value",
				config: &PropConfig{
					Type:   "float",
					Values: []interface{}{true, false},
				},
			},
			err: true,
		},
		{
			name: "interger type with string values",
			args: args{
				name: "value",
				config: &PropConfig{
					Type:   "integer",
					Values: []interface{}{"hello", "bye"},
				},
			},
			err: true,
		},
		{
			name: "values of different types",
			args: args{
				name: "value",
				config: &PropConfig{
					Type:   "string",
					Values: []interface{}{"hello", 2},
				},
			},
			err: true,
		},
		{
			name: "valid string",
			args: args{
				name: "env",
				config: &PropConfig{
					Type:     "string",
					Required: true,
					Values:   []interface{}{"dev", "staging", "prod"},
					Regex:    "^[a-z0-9]{3,6}$",
				},
			},
			expected: &Prop{
				Name: "env",
				Config: &PropConfig{
					Type:     "string",
					Required: true,
					Values:   []interface{}{"dev", "staging", "prod"},
					Regex:    "^[a-z0-9]{3,6}$",
				},
			},
		},
		{
			name: "valid integer with interval",
			args: args{
				name: "num",
				config: &PropConfig{
					Type:     "integer",
					Required: true,
					Interval: &Interval{1, 10},
				},
			},
			expected: &Prop{
				Name: "num",
				Config: &PropConfig{
					Type:     "integer",
					Required: true,
					Interval: &Interval{1, 10},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prop, err := NewProp(test.args.name, test.args.config, test.args.props...)

			assert.Equal(t, test.expected, prop)
			if test.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
