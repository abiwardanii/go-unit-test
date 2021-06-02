package tablebenchmark

import "testing"

func BenchmarkHaloTable(b *testing.B) {
	benchmarks := []struct {
		name 		string
		request 	string
		expected 	string
	}{
		{
			name: "Halo(Abi)",
			request: "Abi",
		},
		{
			name: "Halo(Wardani)",
			request: "Wardani",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B){
			for i := 0; i < b.N; i++ {
				Halo(benchmark.name)
			}
		})
	}
}