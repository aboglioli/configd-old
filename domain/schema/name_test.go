package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	type args struct {
		name string
	}

	type test struct {
		name     string
		args     args
		expected *Name
		err      bool
	}

	tests := []test{
		{
			name: "short name",
			args: args{
				name: "a",
			},
			err: true,
		},
		{
			name: "valid name",
			args: args{
				name: "My Distributed Service 01 !",
			},
			expected: &Name{
				name: "my-distributed-service-01",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n, err := NewName(test.args.name)

			assert.Equal(t, test.expected, n)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
