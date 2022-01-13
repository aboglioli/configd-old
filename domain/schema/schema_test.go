package schema

import (
	"testing"

	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/props"
	"github.com/aboglioli/configd/pkg/models"
	"github.com/aboglioli/configd/pkg/utils"
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

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, str)
				utils.Ok(err)

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

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, obj)
				utils.Ok(err)

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

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, obj)
				utils.Ok(err)

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

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, env)
				utils.Ok(err)

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

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, strs)
				utils.Ok(err)

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

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, env)
				utils.Ok(err)

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

				str, err := props.NewString("str", props.WithRequired())
				utils.Ok(err)

				obj, err := props.NewObject("obj", props.WithProps(integer, str))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, obj)
				utils.Ok(err)

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
				integer, err := props.NewInteger("integers", props.WithArray(), props.WithInterval(2, 6))
				utils.Ok(err)

				obj, err := props.NewObject("obj", props.WithProps(integer))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, obj)
				utils.Ok(err)

				return s
			},
			config: config.ConfigData{
				"obj": map[string]interface{}{
					"integers": []interface{}{2, 3, 4, 5, 6},
				},
			},
			err: false,
		},
		{
			name: "invalid complex array",
			schema: func(t *test) *Schema {
				integers, err := props.NewInteger("integers", props.WithArray(), props.WithInterval(2, 6))
				utils.Ok(err)

				strings, err := props.NewString("strings", props.WithArray(), props.WithEnum("one", "two", "three"))
				utils.Ok(err)

				floats, err := props.NewFloat("floats", props.WithArray(), props.WithInterval(2.5, 15.65))
				utils.Ok(err)

				booleans, err := props.NewBool("booleans", props.WithArray(), props.WithRequired())
				utils.Ok(err)

				object, err := props.NewObject("object", props.WithProps(floats, booleans))
				utils.Ok(err)

				array, err := props.NewObject("array", props.WithArray(), props.WithProps(integers, strings, object))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, integers, strings, array)
				utils.Ok(err)

				return s
			},
			config: config.ConfigData{
				"integers": []interface{}{2, 4, 6},
				"strings":  []interface{}{"one", "three"},
				"array": []interface{}{
					map[string]interface{}{
						"integers": []interface{}{3},
						"strings":  []interface{}{"two"},
						"object": map[string]interface{}{
							"floats":   []interface{}{3.6, 4.8, 15.66}, // invalid number
							"booleans": []interface{}{true, true, false},
						},
					},
					map[string]interface{}{
						"integers": []interface{}{3, 2, 6},
						"strings":  []interface{}{"three"},
						"object": map[string]interface{}{
							"floats":   []interface{}{4.8},
							"booleans": []interface{}{false},
						},
					},
				},
			},
			err: true,
		},
		{
			name: "valid complex array",
			schema: func(t *test) *Schema {
				integers, err := props.NewInteger("integers", props.WithArray(), props.WithInterval(2, 6))
				utils.Ok(err)

				strings, err := props.NewString("strings", props.WithArray(), props.WithEnum("one", "two", "three"))
				utils.Ok(err)

				floats, err := props.NewFloat("floats", props.WithArray(), props.WithInterval(2.5, 15.65))
				utils.Ok(err)

				booleans, err := props.NewBool("booleans", props.WithArray(), props.WithRequired())
				utils.Ok(err)

				object, err := props.NewObject("object", props.WithProps(floats, booleans))
				utils.Ok(err)

				array, err := props.NewObject("array", props.WithArray(), props.WithProps(integers, strings, object))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, integers, strings, array)
				utils.Ok(err)

				return s
			},
			config: config.ConfigData{
				"integers": []interface{}{2, 4, 6},
				"strings":  []interface{}{"one", "three"},
				"array": []interface{}{
					map[string]interface{}{
						"integers": []interface{}{3},
						"strings":  []interface{}{"two"},
						"object": map[string]interface{}{
							"floats":   []interface{}{3.6, 4.8},
							"booleans": []interface{}{true, true, false},
						},
					},
					map[string]interface{}{
						"integers": []interface{}{3, 2, 6},
						"strings":  []interface{}{"three"},
						"object": map[string]interface{}{
							"floats":   []interface{}{4.8},
							"booleans": []interface{}{false},
						},
					},
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

func TestSchemaToMap(t *testing.T) {
	type test struct {
		name     string
		schema   func(t *test) *Schema
		expected map[string]interface{}
	}

	tests := []test{
		{
			name: "basic schema",
			schema: func(t *test) *Schema {
				str, err := props.NewString(
					"str",
					props.WithDefault("default"),
					props.WithEnum("default", "non-default"),
					props.WithRequired(),
				)
				utils.Ok(err)

				int, err := props.NewInteger("int", props.WithInterval(5, 89), props.WithDefault(7))
				utils.Ok(err)

				float, err := props.NewFloat("float", props.WithInterval(1.5, 7.66), props.WithRequired())
				utils.Ok(err)

				n, err := NewName(t.name)

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, str, int, float)
				utils.Ok(err)

				return s
			},
			expected: map[string]interface{}{
				"str": map[string]interface{}{
					"$schema": map[string]interface{}{
						"type":     props.STRING,
						"default":  "default",
						"required": true,
						"enum":     []interface{}{"default", "non-default"},
						"regex":    "",
						"interval": map[string]interface{}(nil),
					},
				},
				"int": map[string]interface{}{
					"$schema": map[string]interface{}{
						"type":     props.INT,
						"default":  7,
						"required": false,
						"enum":     []interface{}(nil),
						"regex":    "",
						"interval": map[string]interface{}{
							"min": 5.0,
							"max": 89.0,
						},
					},
				},
				"float": map[string]interface{}{
					"$schema": map[string]interface{}{
						"type":     props.FLOAT,
						"default":  nil,
						"required": true,
						"enum":     []interface{}(nil),
						"regex":    "",
						"interval": map[string]interface{}{
							"min": 1.5,
							"max": 7.66,
						},
					},
				},
			},
		},
		{
			name: "schema with object",
			schema: func(t *test) *Schema {
				str, err := props.NewString(
					"str",
					props.WithDefault("default"),
					props.WithEnum("default", "non-default"),
					props.WithRequired(),
				)
				utils.Ok(err)

				int, err := props.NewInteger("int", props.WithInterval(5, 89), props.WithDefault(7))
				utils.Ok(err)

				float, err := props.NewFloat("float", props.WithInterval(1.5, 7.66), props.WithRequired())
				utils.Ok(err)

				obj1, err := props.NewObject("obj1", props.WithProps(str, int))
				utils.Ok(err)

				obj2, err := props.NewObject("obj2", props.WithProps(str, float))
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, obj1, obj2)
				utils.Ok(err)

				return s
			},
			expected: map[string]interface{}{
				"obj1": map[string]interface{}{
					"str": map[string]interface{}{
						"$schema": map[string]interface{}{
							"type":     props.STRING,
							"default":  "default",
							"required": true,
							"enum":     []interface{}{"default", "non-default"},
							"regex":    "",
							"interval": map[string]interface{}(nil),
						},
					},
					"int": map[string]interface{}{
						"$schema": map[string]interface{}{
							"type":     props.INT,
							"default":  7,
							"required": false,
							"enum":     []interface{}(nil),
							"regex":    "",
							"interval": map[string]interface{}{
								"min": 5.0,
								"max": 89.0,
							},
						},
					},
				},
				"obj2": map[string]interface{}{
					"str": map[string]interface{}{
						"$schema": map[string]interface{}{
							"type":     props.STRING,
							"default":  "default",
							"required": true,
							"enum":     []interface{}{"default", "non-default"},
							"regex":    "",
							"interval": map[string]interface{}(nil),
						},
					},
					"float": map[string]interface{}{
						"$schema": map[string]interface{}{
							"type":     props.FLOAT,
							"default":  nil,
							"required": true,
							"enum":     []interface{}(nil),
							"regex":    "",
							"interval": map[string]interface{}{
								"min": 1.5,
								"max": 7.66,
							},
						},
					},
				},
			},
		},
		{
			name: "schema with array",
			schema: func(t *test) *Schema {
				strs, err := props.NewString(
					"strs",
					props.WithArray(),
					props.WithRequired(),
				)
				utils.Ok(err)

				objs, err := props.NewObject("objs", props.WithArray(), props.WithProps(strs))
				utils.Ok(err)

				ints, err := props.NewInteger("ints", props.WithInterval(5, 89), props.WithDefault(7), props.WithArray())
				utils.Ok(err)

				n, err := NewName(t.name)
				utils.Ok(err)

				id, err := models.NewSlug(n.Value())
				utils.Ok(err)

				s, err := NewSchema(id, n, objs, ints)
				utils.Ok(err)

				return s
			},
			expected: map[string]interface{}{
				"objs": []interface{}{
					map[string]interface{}{
						"strs": []interface{}{
							map[string]interface{}{
								"$schema": map[string]interface{}{
									"type":     props.STRING,
									"default":  nil,
									"required": true,
									"enum":     []interface{}(nil),
									"regex":    "",
									"interval": map[string]interface{}(nil),
								},
							},
						},
					},
				},
				"ints": []interface{}{
					map[string]interface{}{
						"$schema": map[string]interface{}{
							"type":     props.INT,
							"default":  7,
							"required": false,
							"enum":     []interface{}(nil),
							"regex":    "",
							"interval": map[string]interface{}{
								"min": 5.0,
								"max": 89.0,
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := test.schema(&test)

			assert.Equal(t, test.expected, s.ToMap())
		})
	}
}
