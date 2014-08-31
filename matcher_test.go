package zen

import (
	"testing"
)

func differsByOne(a, b interface{}) bool {
	diff := a.(int) - b.(int)
	return diff == 1 || diff == -1
}

type output struct {
	contextPrinted bool
	whenPrinted    bool
	titlePrinted   bool
	errors         string
}

func (out *output) PrintContext() {
	out.contextPrinted = true
}

func (out *output) PrintWhen() {
	out.whenPrinted = true
}

func (out *output) PrintTitle() {
	out.titlePrinted = true
}
func (out *output) PrintError(err string) {
	out.errors += err
}

// Prints title, even when test passes
func TestMatcherPrintsTitle(t *testing.T) {
	out := new(output)
	expectation := &Expectation{out, 1}
	expectation.ToEqual(1)
	if titlePrinted := out.titlePrinted; !titlePrinted {
		t.Errorf("should have printed title, but didn't")
	}
}

// Prints custom error message
func TestCustomMatcher(t *testing.T) {
	out := new(output)
	expectation := &Expectation{out, 1}
	expectation.To("differ by one from", differsByOne, 10)
	if err := out.errors; err != "Expected `1` to differ by one from `10`" {
		t.Errorf("incorrect error output for custom matcher: %s", err)
	}
}

func TestCustomMatcherPassesWithNoError(t *testing.T) {
	out := new(output)
	expectation := &Expectation{out, 1}
	expectation.To("differ by one from", differsByOne, 2)
	if err := out.errors; err != "" {
		t.Errorf("non-blank error '%s' even though it should have passed", err)
	}
}

var expect Expect

func ExampleExpectation_To() {
	divisibleBy := func(a, b interface{}) bool {
		return a.(int)%b.(int) == 0
	}
	expect(9).To("be divisible by", divisibleBy, 3)
}
