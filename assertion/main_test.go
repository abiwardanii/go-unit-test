package assertion

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHaloAbi(t *testing.T){
	hasil := HaloAbi("Abi")
	assert.Equal(t, "Halo Abi", hasil, "Hasilnya harus Halo Abi")
	fmt.Println("Testing HaloAbi selesai")
}

func TestHaloWardani(t *testing.T){
	hasil := HaloWardani("Wardani")
	assert.Equal(t, "Halo Wardani", hasil, "Hasilnya harus Halo Wardani")
	fmt.Println("Testing HaloWardani selesai")

}