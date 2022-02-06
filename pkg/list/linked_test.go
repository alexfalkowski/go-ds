package list_test

import (
	"testing"

	"github.com/alexfalkowski/go-ds/pkg/list"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLinkedAddInvalidHead(t *testing.T) {
	Convey("Given I have an empty linked list", t, func() {
		lst := list.NewLinkedList()

		Convey("When I add an invalid head to the linked list", func() {
			head := lst.AddHead(nil)

			Convey("Then I should have no head", func() {
				So(head, ShouldBeNil)
				So(lst.Length(), ShouldEqual, 0)
			})
		})
	})
}

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

func TestLinkedAddMultipleHead(t *testing.T) {
	Convey("Given I have a multiple node linked list", t, func() {
		lst := list.NewLinkedList()
		lst.AddHead(1)
		lst.AddHead(2)

		Convey("When I add a new head to the linked list", func() {
			head := lst.AddHead(3)

			Convey("Then I should have a single node", func() {
				So(head, ShouldNotBeNil)
				So(head.Data, ShouldEqual, 3)
				So(lst.Length(), ShouldEqual, 3)
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
	Convey("Given I have a single element linked list", t, func() {
		lst := list.NewLinkedList()
		lst.AddHead(1)

		Convey("When I remove the head from the linked list", func() {
			head := lst.RemoveHead()

			Convey("Then I should have the head", func() {
				So(head, ShouldNotBeNil)
				So(head.Data, ShouldEqual, 1)
				So(lst.Length(), ShouldEqual, 0)
			})
		})
	})
}

func TestLinkedRemoveMultipleHead(t *testing.T) {
	Convey("Given I have a multiple element linked list", t, func() {
		lst := list.NewLinkedList()
		lst.AddHead(1)
		lst.AddHead(2)

		Convey("When I remove the head from the linked list", func() {
			head := lst.RemoveHead()

			Convey("Then I should have the head", func() {
				So(head, ShouldNotBeNil)
				So(head.Data, ShouldEqual, 2)
				So(lst.Length(), ShouldEqual, 1)
			})
		})
	})
}

func TestLinkedAddInvalidTail(t *testing.T) {
	Convey("Given I have an empty linked list", t, func() {
		lst := list.NewLinkedList()

		Convey("When I add an invalid tail to the linked list", func() {
			tail := lst.AddTail(nil)

			Convey("Then I should have no tail", func() {
				So(tail, ShouldBeNil)
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

		Convey("When I add a new tail to the linked list", func() {
			tail := lst.AddTail(2)

			Convey("Then I should have a single node", func() {
				So(tail, ShouldNotBeNil)
				So(tail.Data, ShouldEqual, 2)
				So(lst.Length(), ShouldEqual, 2)
			})
		})
	})
}

func TestLinkedAddMultipleTail(t *testing.T) {
	Convey("Given I have a multiple node linked list", t, func() {
		lst := list.NewLinkedList()
		lst.AddTail(1)
		lst.AddTail(2)

		Convey("When I add a new tail to the linked list", func() {
			tail := lst.AddTail(3)

			Convey("Then I should have a single node", func() {
				So(tail, ShouldNotBeNil)
				So(tail.Data, ShouldEqual, 3)
				So(lst.Length(), ShouldEqual, 3)
			})
		})
	})
}

func TestLinkedRemoveTail(t *testing.T) {
	Convey("Given I have an empty linked list", t, func() {
		lst := list.NewLinkedList()

		Convey("When I remove the tail from the linked list", func() {
			tail := lst.RemoveTail()

			Convey("Then I should have no tail", func() {
				So(tail, ShouldBeNil)
				So(lst.Length(), ShouldEqual, 0)
			})
		})
	})
}

func TestLinkedRemoveSingleTail(t *testing.T) {
	Convey("Given I have a single element linked list", t, func() {
		lst := list.NewLinkedList()
		lst.AddTail(1)

		Convey("When I remove the tail from the linked list", func() {
			tail := lst.RemoveTail()

			Convey("Then I should have the tail", func() {
				So(tail, ShouldNotBeNil)
				So(tail.Data, ShouldEqual, 1)
				So(lst.Length(), ShouldEqual, 0)
			})
		})
	})
}

func TestLinkedRemoveMultipleTail(t *testing.T) {
	Convey("Given I have a single element linked list", t, func() {
		lst := list.NewLinkedList()
		lst.AddTail(1)
		lst.AddTail(2)

		Convey("When I remove the tail from the linked list", func() {
			tail := lst.RemoveTail()

			Convey("Then I should have the tail", func() {
				So(tail, ShouldNotBeNil)
				So(tail.Data, ShouldEqual, 2)
				So(lst.Length(), ShouldEqual, 1)
			})
		})
	})
}
