package subtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T){
	t.Run("Abi", func(t *testing.T) {
		hasil := HaloAbi("Abi")
		require.Equal(t, "Halo Abi", hasil)
	})
	t.Run("Wardani", func(t *testing.T) {
		hasil := HaloWardani("Wardani")
		require.Equal(t, "Halo Wardani", hasil)
	})
}