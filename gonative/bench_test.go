package gonative

import "testing"

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(35)
	}
}

func BenchmarkFibtc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibtc(35, 0, 1)
	}
}