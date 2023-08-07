package my_math

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSub(t *testing.T) {
	Convey("testSub", t, func() {
		result := Sub(1, 2)
		So(result, ShouldEqual, -1)
	})
}

func TestSum(t *testing.T) {
	Convey("testSum", t, func() {
		result := Sum(1, 2)
		So(result, ShouldEqual, 3)
	})
}
