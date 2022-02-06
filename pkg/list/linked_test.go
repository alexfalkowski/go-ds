package list_test

import (
	"testing"

	"github.com/alexfalkowski/go-ds/pkg/list"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLinkedAddHead(t *testing.T) {
	Convey("Given I have an empty linked list", t, func() {
		lst := list.NewLinkedList()

		Convey("When I add a new head to the linked list", func() {
			head := lst.AddHead(1)

			Convey("Then I should have a single node", func() {
				So(head, ShouldNotBeNil)
				So(head.Data, ShouldEqual, 1)
				So(lst.Length(), ShouldEqual, 1)
			})
		})
	})
}

func TestLinkedAddDoubleHead(t *testing.T) {
	Convey("Given I have a single node linked list", t, func() {
		lst := list.NewLinkedList()
		lst.AddHead(1)

		Convey("When I add a new head to the linked list", func() {
			head := lst.AddHead(2)

			Convey("Then I should have a single node", func() {
				So(head, ShouldNotBeNil)
				So(head.Data, ShouldEqual, 2)
				So(lst.Length(), ShouldEqual, 2)
			})
		})
	})
}

func TestLinkedRemoveHead(t *testing.T) {
	Convey("Given I have an empty linked list", t, func() {
		lst := list.NewLinkedList()

		Convey("When I remove the head from the linked list", func() {
			head := lst.RemoveHead()

			Convey("Then I should have no head", func() {
				So(head, ShouldBeNil)
				So(lst.Length(), ShouldEqual, 0)
			})
		})
	})
}

func TestLinkedRemoveSingleHead(t *testing.T) {
	Convey("Given I have an empty linked list", t, func() {
		lst := list.NewLinkedList()
		lst.AddHead(1)

		Convey("When I remove the head from the linked list", func() {
			head := lst.RemoveHead()

			Convey("Then I should have no head", func() {
				So(head, ShouldNotBeNil)
				So(head.Data, ShouldEqual, 1)
				So(lst.Length(), ShouldEqual, 0)
			})
		})
	})
}

func TestLinkedAddTail(t *testing.T) {
	Convey("Given I have an empty linked list", t, func() {
		lst := list.NewLinkedList()

		Convey("When I add a new tail to the linked list", func() {
			tail := lst.AddTail(1)

			Convey("Then I should have a single node", func() {
				So(tail, ShouldNotBeNil)
				So(tail.Data, ShouldEqual, 1)
				So(lst.Length(), ShouldEqual, 1)
			})
		})
	})
}

func TestLinkedAddDoubleTail(t *testing.T) {
	Convey("Given I have a single node linked list", t, func() {
		lst := list.NewLinkedList()
		lst.AddTail(1)

		Convey("When I add a new head to the linked list", func() {
			tail := lst.AddTail(2)

			Convey("Then I should have a single node", func() {
				So(tail, ShouldNotBeNil)
				So(tail.Data, ShouldEqual, 2)
				So(lst.Length(), ShouldEqual, 2)
			})
		})
	})
}
