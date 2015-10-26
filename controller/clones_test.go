package controller

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClones(t *testing.T) {
	Convey("Given a folder of sample configurations", t, func(context C) {
		// $REPO/samples/config

		Convey("listDirectory must return all files", func() {
			files, err := listDirectory("../samples/config")

			So(err, ShouldBeNil)
			So(files, ShouldNotBeNil)
			So(len(files), ShouldEqual, 5)

			So(files[0], ShouldEqual, "invalid1.json")
			So(files[1], ShouldEqual, "invalid2.json")
			So(files[2], ShouldEqual, "invalid3.json")
			So(files[3], ShouldEqual, "invalid4.json")
			So(files[4], ShouldEqual, "valid1.json")
		})

		Convey("listDirectory must return folders", func() {
			files, err := listDirectory("../samples")

			So(err, ShouldBeNil)
			So(files, ShouldNotBeNil)
			So(len(files), ShouldEqual, 2)
		})
	})
}
