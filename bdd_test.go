package zen

import (
	"testing"
)

func ExampleDesc(t *testing.T) {
	Desc(t, "Equality", func(it It) {
		it("any integer should equal itself", func(expect Expect) {
			expect(1).ToEqual(1)
		})
		it("any integer should not equal nil", func(expect Expect) {
			expect(1).ToNotEqual(nil)
			expect(1).ToExist() // Same as above
		})
	})
}

func TestSetupAndTeardown(t *testing.T) {
	count := 0

	before := func() {
		count++
	}

	after := func() {
		count--
	}

	setup := Setup(before, after)

	Desc(t, "Setup", func(it It) {
		it("should execute before", setup(func(expect Expect) {
			expect(count).ToEqual(1)
		}))

		if count != 0 {
			t.Error("Count should have been reset to zero by the teardown func")
		}
	})
}
