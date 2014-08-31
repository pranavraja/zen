package zen

import (
	"testing"
)

type Test struct {
	T       *testing.T
	Context string
	When    string
	Title   string
	Fn      func(Expect)
}

func NewTest(t *testing.T, context string, when string, title string, fn func(Expect)) *Test {
	return &Test{
		t,
		context,
		when,
		title,
		fn,
	}
}

func (test *Test) Run() {
	test.Fn(func(val interface{}) *Expectation {
		return &Expectation{test, val}
	})
}

type When func(when string, fn func(It))
type It func(title string, fn func(Expect))

func Given(t *testing.T, context string, scenerioWrapper func(When)) {
	scenerioWrapper(func(when string, testWrapper func(It)) {
		testWrapper(func(it string, fn func(Expect)) {
			test := &Test{
				t,
				context,
				when,
				it,
				fn,
			}
			test.Run()
		})
	})
}

func Desc(t *testing.T, desc string, wrapper func(It)) {
	wrapper(func(it string, fn func(Expect)) {
		test := &Test{
			t,
			"",
			desc,
			it,
			fn,
		}
		test.Run()
	})

}

func Setup(before, after func()) func(fn func(Expect)) func(Expect) {
	return func(fn func(Expect)) func(Expect) {
		before()
		return func(expect Expect) {
			fn(expect)
			after()
		}
	}
}

func NotImplemented() func(Expect) {
	return func(expect Expect) { expect(nil).NotImplemented() }
}
