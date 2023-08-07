package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	x := 0
	Convey("test1", t, func() {
		x++
		So(1, ShouldEqual, 1)
		So(2*2, ShouldEqual, 4)
		Convey("test1-1", func() {
			x++
			So(1, ShouldEqual, 1)
			So(2*2, ShouldEqual, 4)
			Convey("test1-1-1", func() {
				So(x, ShouldEqual, 2)
			})
			Convey("test1-1-2", func() {
				So(x, ShouldEqual, 4)
			})

		})

		Convey("test1-2", func() {
			x++
			So(1, ShouldEqual, 1)
			So(2*2, ShouldEqual, 4)
			Convey("test1-2-1", func() {
				So(x, ShouldEqual, 6)
			})
			Convey("test1-2-2", func() {
				So(x, ShouldEqual, 8)
				//panic("PANIC AAAAAAAA")
			})

		})

		Reset(func() {
			t.Log("finish")
		})
	})
}
