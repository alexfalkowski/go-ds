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
