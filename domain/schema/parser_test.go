package schema

import (
	"testing"

	"github.com/aboglioli/configd/domain/props"
	"github.com/stretchr/testify/assert"
)

func ok(err error) {
	if err != nil {
		panic(err)
	}
}

func TestSchemaFromJson(t *testing.T) {
	type args struct {
		json string
	}

	tests := []struct {
		name     string
		args     args
		expected func() Schema
		err      bool
	}{
		{
			name: "valid",
			args: args{
				json: `
					{
					  "env":{
						"$schema":{
						  "type":"string",
						  "values":[
							"dev",
							"staging",
							"prod"
						  ]
						}
					  },
					  "version":{
						"$schema":{
						  "type":"string",
						  "regex":"v[0-9]+"
						}
					  },
					  "internal_service":{
						"url":{
						  "$schema":{
							"type":"string",
							"required":true
						  }
						},
						"port":{
						  "$schema":{
							"type":"integer",
							"default":8080,
							"interval":{
							  "min":80,
							  "max":18080
							}
						  }
						}
					  },
					  "circuit_breaker":{
						"threshold":{
						  "$schema":{
							"type":"float",
							"default":0.6
						  }
						}
					  },
					  "nested_object":{
						"inner_object":{
						  "value":{
							"$schema":{
							  "type":"float",
							  "default":12.75
							}
						  }
						}
					  }
					}
				`,
			},
			expected: func() Schema {
				// Env
				env, err := props.NewString("env", props.WithValues("dev", "staging", "prod"))
				ok(err)

				// Version
				version, err := props.NewString("version", props.WithRegex("v[0-9]+"))
				ok(err)

				// Internal service
				url, err := props.NewString("url", props.WithRequired(true))
				ok(err)

				port, err := props.NewInteger("port", props.WithDefault(8080), props.WithInterval(80, 18080))
				ok(err)

				internalService, err := props.NewObject(
					"internal_service",
					props.WithProps(url, port),
				)
				ok(err)

				// Circuit breaker
				threshold, err := props.NewFloat("threshold", props.WithDefault(0.6))
				ok(err)

				circuitBreaker, err := props.NewObject("circuit_breaker", props.WithProps(threshold))
				ok(err)

				// Nested objects
				value, err := props.NewFloat("value", props.WithDefault(12.75))
				ok(err)

				innerObject, err := props.NewObject("inner_object", props.WithProps(value))
				ok(err)

				nestedObject, err := props.NewObject("nested_object", props.WithProps(innerObject))
				ok(err)

				s, err := NewSchema(
					env,
					version,
					internalService,
					circuitBreaker,
					nestedObject,
				)
				ok(err)

				return s
			},
		},
		{
			name: "schema without type",
			args: args{
				json: `
					{
					  "env":{
						"$schema":{
						  "type":"string",
						  "values":[
							"dev",
							"staging",
							"prod"
						  ]
						}
					  },
					  "version":{
						"$schema":{
						  "regex":"v[0-9]+"
						}
					  },
					}
				`,
			},
			expected: func() Schema {
				return nil
			},
			err: true,
		},
		{
			name: "empty schema",
			args: args{
				json: `{}`,
			},
			expected: func() Schema {
				return nil
			},
			err: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			schema, err := FromJson(test.args.json)

			assert.Equal(t, test.expected(), schema)
			if test.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSchemaFromMAp(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}

	tests := []struct {
		name     string
		args     args
		expected func() Schema
		err      bool
	}{
		{
			name: "valid",
			args: args{
				m: map[string]interface{}{
					"env": map[string]interface{}{
						"$schema": map[string]interface{}{
							"type": "integer",
							"interval": map[string]interface{}{
								"min": 1,
								"max": 256,
							},
						},
					},
				},
			},
			expected: func() Schema {
				// Env
				env, err := props.NewInteger("env", props.WithInterval(1, 256))
				ok(err)

				s, err := NewSchema(
					env,
				)
				ok(err)

				return s
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			schema, err := FromMap(test.args.m)

			assert.Equal(t, test.expected(), schema)
			if test.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
