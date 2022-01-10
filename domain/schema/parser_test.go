package schema

import (
	"testing"

	"github.com/aboglioli/configd/domain/props"
	"github.com/aboglioli/configd/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPropsFromJson(t *testing.T) {
	type test struct {
		name     string
		json     string
		expected func(t *test) []*props.Prop
		err      bool
	}

	tests := []test{
		{
			name: "valid",
			json: `
				{
				  "env":{
					"$schema":{
					  "type":"string",
					  "enum":[
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
			expected: func(t *test) []*props.Prop {
				// Env
				env, err := props.NewString("env", props.WithEnum("dev", "staging", "prod"))
				utils.Ok(err)

				// Version
				version, err := props.NewString("version", props.WithRegex("v[0-9]+"))
				utils.Ok(err)

				// Internal service
				url, err := props.NewString("url", props.WithRequired())
				utils.Ok(err)

				port, err := props.NewInteger("port", props.WithDefault(8080), props.WithInterval(80, 18080))
				utils.Ok(err)

				internalService, err := props.NewObject(
					"internal_service",
					props.WithProps(url, port),
				)
				utils.Ok(err)

				// Circuit breaker
				threshold, err := props.NewFloat("threshold", props.WithDefault(0.6))
				utils.Ok(err)

				circuitBreaker, err := props.NewObject("circuit_breaker", props.WithProps(threshold))
				utils.Ok(err)

				// Nested objects
				value, err := props.NewFloat("value", props.WithDefault(12.75))
				utils.Ok(err)

				innerObject, err := props.NewObject("inner_object", props.WithProps(value))
				utils.Ok(err)

				nestedObject, err := props.NewObject("nested_object", props.WithProps(innerObject))
				utils.Ok(err)

				return []*props.Prop{
					env,
					version,
					internalService,
					circuitBreaker,
					nestedObject,
				}
			},
		},
		{
			name: "valid with array",
			json: `
				{
				  "cache": {
					"addresses": [
					  {
						"$schema": {
						  "type": "string",
						  "regex": "^https://[a-z0-9.-]+$"
						}
					  }
					]
				  },
				  "inner": {
					"object": [
					  {
						"prop": {
						  "$schema": {
							"type": "string"
						  }
						}
					  }
					]
				  },
				  "arr1": [
					{
					  "arr2": [
						{
						  "float": {
							"$schema": {
							  "type": "float",
							  "interval": {
								"min": 1.5,
								"max": 6.8
							  }
							}
						  },
						  "integers": [
							{
							  "$schema": {
								"type": "integer",
								"enum": [4, 5, 6]
							  }
							}
						  ]
						}
					  ]
					}
				  ]
				}
			`,
			expected: func(t *test) []*props.Prop {
				// Cache
				addresses, err := props.NewString(
					"addresses",
					props.WithArray(),
					props.WithRegex("^https://[a-z0-9.-]+$"),
				)
				utils.Ok(err)

				cache, err := props.NewObject("cache", props.WithProps(addresses))
				utils.Ok(err)

				// Inner
				prop, err := props.NewString("prop")
				utils.Ok(err)

				object, err := props.NewObject("object", props.WithProps(prop), props.WithArray())
				utils.Ok(err)

				inner, err := props.NewObject("inner", props.WithProps(object))
				utils.Ok(err)

				// Nested arrays
				float, err := props.NewFloat("float", props.WithInterval(1.5, 6.8))
				utils.Ok(err)

				integeres, err := props.NewInteger(
					"integers",
					props.WithArray(),
					props.WithEnum(4, 5, 6),
				)
				utils.Ok(err)

				arr2, err := props.NewObject(
					"arr2",
					props.WithArray(),
					props.WithProps(float, integeres),
				)
				utils.Ok(err)

				arr1, err := props.NewObject(
					"arr1",
					props.WithArray(),
					props.WithProps(arr2),
				)

				return []*props.Prop{
					cache,
					inner,
					arr1,
				}
			},
		},
		{
			name: "invalid object with array schema",
			json: `
				{
				  "cache": {
					"addresses": [
					  {
						"$schema": {
						  "type": "string",
						  "regex": "^https://[a-z0-9.-]+$"
						}
					  },
					  {
						"$schema": {
						  "type": "string",
						  "regex": "^https://[a-z0-9.-]+$"
						}
					  }
					]
				  }
				}
			`,
			expected: func(t *test) []*props.Prop {
				return nil
			},
			err: true,
		},
		{
			name: "invalid object with multiple objects in array",
			json: `
				{
				  "cache": {
					"addresses": [
					  {
						"type": {
						  "$schema": {
							"type": "string"
						  }
						}
					  },
					  {
						"type": {
						  "$schema": {
							"type": "string"
						  }
						}
					  }
					]
				  }
				}
			`,
			expected: func(t *test) []*props.Prop {
				return nil
			},
			err: true,
		},
		{
			name: "schema without type",
			json: `
				{
				  "env":{
					"$schema":{
					  "type":"string",
					  "enut":[
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
			expected: func(t *test) []*props.Prop {
				return nil
			},
			err: true,
		},
		{
			name: "empty schema",
			json: `{}`,
			expected: func(t *test) []*props.Prop {
				return nil
			},
			err: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ps, err := PropsFromJson(test.json)
			expectedProps := test.expected(&test)

			if test.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)

				propsMap := make(map[string]*props.Prop)
				for _, p := range ps {
					propsMap[p.Name()] = p
				}

				expectedPropsMap := make(map[string]*props.Prop)
				for _, p := range expectedProps {
					expectedPropsMap[p.Name()] = p
				}

				assert.Equal(t, expectedPropsMap, propsMap)
			}
		})
	}
}

func TestPropsFromMap(t *testing.T) {
	type test struct {
		name     string
		m        map[string]interface{}
		expected func(t *test) []*props.Prop
		err      bool
	}

	tests := []test{
		{
			name: "valid",
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
			expected: func(t *test) []*props.Prop {
				// Env
				env, err := props.NewInteger("env", props.WithInterval(1, 256))
				utils.Ok(err)

				return []*props.Prop{env}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ps, err := PropsFromMap(test.m)
			expectedProps := test.expected(&test)

			if test.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)

				propsMap := make(map[string]*props.Prop)
				for _, p := range ps {
					propsMap[p.Name()] = p
				}

				expectedPropsMap := make(map[string]*props.Prop)
				for _, p := range expectedProps {
					expectedPropsMap[p.Name()] = p
				}

				assert.Equal(t, expectedPropsMap, propsMap)
			}
		})
	}
}
