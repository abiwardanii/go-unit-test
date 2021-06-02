package subbenchmark

import "testing"

func BenchmarkSub(b *testing.B){
	b.Run("Abi", func (b *testing.B){
		for i := 0; i < b.N; i++ {
			HaloAbi("Abi")
		}
	})
	b.Run("Wardani", func (b *testing.B){
		for i := 0; i < b.N; i++ {
			HaloWardani("Wardani")
		}
	})
}