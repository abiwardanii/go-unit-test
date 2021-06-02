package beforeafter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M){
	//before
	fmt.Println("sebeleum unit test")
	m.Run()

	//after
	fmt.Println("sesudah unit test")

}	

func TestHaloAbi(t *testing.T){
	hasil := HaloAbi("Abi")
	require.Equal(t, "Halo Abi", hasil, "Hasilnya harus Halo Abi")
	fmt.Println("Testing HaloAbi selesai")
}