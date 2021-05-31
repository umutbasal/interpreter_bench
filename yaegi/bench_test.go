package yaegi

import "testing"

func BenchmarkFib(b *testing.B) {
	interpreter, _ := NewInterpreter("fib.yaegi")
	fib, _ := initFib(interpreter)
	for n := 0; n < b.N; n++ {
		fib(35)
	}
}

func BenchmarkFibtc(b *testing.B) {
	interpreter, _ := NewInterpreter("fibtc.yaegi")
	fibtc, _ := initFibtc(interpreter)
	for n := 0; n < b.N; n++ {
		fibtc(35, 0, 1)
	}
}
