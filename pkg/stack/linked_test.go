package stack_test

import (
	"testing"

	"github.com/alexfalkowski/go-ds/pkg/stack"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLinkedStackPush(t *testing.T) {
	Convey("Given I have an empty linked stack", t, func() {
		stk := stack.NewLinkedStack()

		Convey("When I do not push a node", func() {
			Convey("Then I should have no node", func() {
				So(stk.Peek(), ShouldBeNil)
				So(stk.Pop(), ShouldBeNil)
			})
		})

		Convey("When I push an node", func() {
			stk.Push(1)

			Convey("Then I should have a node", func() {
				So(stk.Peek(), ShouldEqual, 1)
				So(stk.Pop(), ShouldEqual, 1)
			})
		})
	})
}
