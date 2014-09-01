package zen

import (
	"fmt"
)

type Expect func(val interface{}) *Expectation

type formatter interface {
	PrintContext()
	PrintWhen()
	PrintTitle()
	PrintTitleNotImplemented()
	PrintError(string)
}

type Expectation struct {
	Output formatter
	Value  interface{}
}

type Matcher func(a, b interface{}) bool

func Equal(a, b interface{}) bool {
	return a == b
}

func NotEqual(a, b interface{}) bool {
	return a != b
}

func NotExist(a, b interface{}) bool {
	return a == nil
}

func Exist(a, b interface{}) bool {
	return a != nil
}

func (self *Expectation) To(desc string, match Matcher, value interface{}) {
	self.Output.PrintContext()
	self.Output.PrintWhen()
	self.Output.PrintTitle()
	if !match(self.Value, value) {
		self.Output.PrintError(fmt.Sprintf("Expected `%v` to %s `%v`", self.Value, desc, value))
	}
}

func (self *Expectation) ToEqual(b interface{}) {
	self.To("equal", Equal, b)
}

func (self *Expectation) ToNotEqual(b interface{}) {
	self.To("not equal", NotEqual, b)
}

func (self *Expectation) ToExist() {
	self.To("exist", Exist, nil)
}

func (self *Expectation) ToNotExist() {
	self.To("not exist", NotExist, nil)
}

func (self *Expectation) NotImplemented() {
	self.Output.PrintContext()
	self.Output.PrintWhen()
	self.Output.PrintTitleNotImplemented()
}
