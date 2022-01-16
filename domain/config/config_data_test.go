package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashConfigData(t *testing.T) {
	type test struct {
		name       string
		configData ConfigData
		expected   string
	}

	tests := []test{
		{
			name: " simple",
			configData: ConfigData{
				"message": "hello",
			},
			expected: "9b2d43affbf49a367028df2e1414f84c0e099ac98c3d54a8a80157fd7771af25",
		},
		{
			name: " simple",
			configData: ConfigData{
				"object": map[string]interface{}{
					"array": []interface{}{1, 2, 3},
				},
			},
			expected: "6e63cf6bba4ae2550efaed311c41c83b1a3b50d0bbb38e0594405e11d3e7d18e",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hash := test.configData.Hash()

			assert.Equal(t, test.expected, hash)
		})
	}
}
