## Zen

BDD Testing Framework For Go.

`go get github.com/pranavraja/zen`

Forked from [github.com/azer/mao](https://github.com/azer/mao).

Changes in my fork:
- Print each behaviour as it's being verified
- Actually fail the testrunner when there are expectation failures
- Small API changes
- Allow custom matchers

```go
package zen

import (
    "testing"
    . "github.com/pranavraja/zen"
)

func TestZen(t *testing.T) {
    Desc(t, "zen", func(it It) {
        it("should know when things exist", func(expect Expect) {
            expect("tree").ToExist()
        })
        it("should know when things don't exist", func(expect Expect) {
            expect(nil).ToNotExist()
        })
        it("should know that a thing is equal to itself", func(expect Expect) {
            expect(1).ToEqual(1)
        })
        it("should know when things are different", func(expect Expect) {
            expect(1).ToNotEqual(2)
        })
        it("should be able to learn about new things", func (expect Expect) {
            divisibleBy := func (a, b interface{}) bool {
                return a.(int) % b.(int) == 0
            }
            expect(1).To("be divisible by", divisibleBy, 1)
        })
    })
}
```

## Output

Pass:

![Test passed](http://i.imgur.com/zmkJOTW.png)

Fail:

![Test failure](http://i.imgur.com/CWiy8wi.png)


## Before and After test setup

```go
package zen

import (
    "testing"
    . "github.com/pranavraja/zen"
)

func TestZen(t *testing.T) {
    Desc(t, "before and after", func(it It) {
        count := 0

        before := func() {
          count++
        }

        after := func() {
          count--
        }

        setup := Setup(before, after)

        it("should execute before and after functions", setup(func(expect Expect) {
            expect(count).ToEqual(1)
        }))

        it("should execute before and after functions", setup(func(expect Expect) {
            expect(count).ToEqual(1)
        }))
    })
}
```
