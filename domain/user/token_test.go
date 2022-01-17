package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenSigningAndValidation(t *testing.T) {
	type test struct {
		name            string
		data            TokenData
		errOnGenerating bool
		errOnParsing    bool
	}

	tests := []test{
		{
			name: "basic data",
			data: TokenData{
				"message": "hello",
			},
			errOnGenerating: false,
			errOnParsing:    false,
		},
		{
			name:            "nil data",
			errOnGenerating: true,
			errOnParsing:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			token, err := GenerateToken(test.data)
			if test.errOnGenerating {
				assert.Error(t, err)
			} else {
				if assert.NoError(t, err) {
					if assert.Greater(t, len(token.token), 30) {
						d, err := token.ParseData()
						if test.errOnParsing {
							assert.Error(t, err)
						} else {
							assert.NoError(t, err)
							assert.Equal(t, test.data, d)
						}
					}
				}
			}
		})
	}
}
