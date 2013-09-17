package zen

import (
	"fmt"
	"testing"
)

type Test struct {
	T     *testing.T
	Title string
	Fn    func(Expect)
}

func (test *Test) Run() {
	test.Fn(func(val interface{}) *Expectation {
		return &Expectation{test, val}
	})
}

type It func(title string, fn func(Expect))

func Desc(t *testing.T, desc string, wrapper func(It)) {
	wrapper(func(it string, fn func(Expect)) {
		test := Test{t, fmt.Sprintf("%s %s", desc, it), fn}
		test.Run()
	})

}

func Setup(before, after func()) func(fn func(Expect)) (func(Expect)) {
	return func(fn func(Expect)) func(Expect) {
		before()
		return func(expect Expect) {
			fn(expect)
			after()
		}
	}
}
