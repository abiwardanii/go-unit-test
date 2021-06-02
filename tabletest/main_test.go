package tabletest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHaloTable(t *testing.T) {
	tests := []struct {
		name 		string
		request 	string
		expected 	string
	}{
		{
			name: "Halo(Abi)",
			request: "Abi",
			expected: "Halo Abi",
		},
		{
			name: "Halo(Wardani)",
			request: "Wardani",
			expected: "Halo Wardani",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			hasil := Halo(test.request)
			assert.Equal(t, test.expected, hasil)
		})
	}
}