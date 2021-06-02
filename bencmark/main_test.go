package bencmark

import "testing"

func BenchmarkHaloAbi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HaloAbi("Abi")
	}
}

func BenchmarkHaloWardani(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HaloWardani("Wardani")
	}
}