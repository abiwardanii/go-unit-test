package helper

import (
	"fmt"
	"testing"
)

func TestHaloAbi(t *testing.T){
	hasil := HaloAbi("Abi")

	if hasil != "Halo Abi"{
		//error
		t.Error("Harusnya Halo Abi")
	}
	fmt.Println("Testing HaloAbi selesai")
}

func TestHaloWardani(t *testing.T){
	hasil := HaloWardani("Wardani")

	if hasil != "Halo Wardani"{
		//error
		t.Error("Harusnya Halo Wardani")
	}
	fmt.Println("Testing HaloWardani selesai")

}