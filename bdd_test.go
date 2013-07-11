package mao

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
