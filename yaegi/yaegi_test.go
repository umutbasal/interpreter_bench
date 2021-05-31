package yaegi

import (
	"testing"
)

func Test_newInterpreter(t *testing.T) {
	interpreter, err := NewInterpreter("fib.yaegi")
	fib, err := initFib(interpreter)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if fib(35) != 9227465 {
		t.Error("unexpected err")
	}
}
