package yaegi

import (
	"io/ioutil"

	"github.com/traefik/yaegi/interp"
)

func read(filename string) (string, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func NewInterpreter(filename string) (i *interp.Interpreter, err error) {
	src, err := read(filename)
	i = interp.New(interp.Options{})
	_, err = i.Eval(src)
	return i, err
}

func initFib(i *interp.Interpreter) (fib func(int) int, err error) {
	v, err := i.Eval("yaegi.fib")
	fib = v.Interface().(func(int) int)
	return fib, err
}

func initFibtc(i *interp.Interpreter) (fibtc func(int, int, int) int, err error) {
	v, err := i.Eval("yaegi.fibtc")
	fibtc = v.Interface().(func(int, int, int) int)
	return fibtc, err
}
