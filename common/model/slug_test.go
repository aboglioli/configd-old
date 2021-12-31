package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlug(t *testing.T) {
	type args struct {
		name string
	}

	type test struct {
		name     string
		args     args
		expected *Slug
		err      bool
	}

	tests := []test{
		{
			name: "short name",
			args: args{
				name: "",
			},
			err: true,
		},
		{
			name: "valid name",
			args: args{
				name: "My Distributed Service 01 !",
			},
			expected: &Slug{
				slug: "my-distributed-service-01",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n, err := NewSlug(test.args.name)

			assert.Equal(t, test.expected, n)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
