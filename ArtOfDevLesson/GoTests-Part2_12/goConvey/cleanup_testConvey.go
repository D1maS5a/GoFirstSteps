package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExampleCleanup(t *testing.T) {
	x := 0
	Convey("A", t, func() {
		x++ // setup
		Convey("A-B", func() {
			x++
			Convey("A-B-C1", func() {
				So(x, ShouldEqual, 2)
				// So(x, ShouldAlmostEqual, 1.98, .03)
			})
			Convey("A-B-C2", func() {
				So(x, ShouldEqual, 4)
			})
			Convey("A-B-C3", func() {
				So(x, ShouldEqual, 6)
			})
		})
		Reset(func() {
			t.Log("Finish")
		})
	})
	// Convey("Something works properly", t, func() {
	// 	So(1, ShouldEqual, 1)
	// 	So(2*2, ShouldEqual, 4)
	// })
}

// $GOPATH/bin/goconvey
