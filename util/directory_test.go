package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDirectory(t *testing.T) {
	Convey("Given a folder of sample configurations", t, func(context C) {
		// $REPO/samples/config

		Convey("ListDirectory should return all files", func() {
			files, err := ListDirectory("../samples/config")

			So(err, ShouldBeNil)
			So(files, ShouldNotBeNil)
			So(len(files), ShouldEqual, 5)

			So(files[0], ShouldEqual, "invalid1.json")
			So(files[1], ShouldEqual, "invalid2.json")
			So(files[2], ShouldEqual, "invalid3.json")
			So(files[3], ShouldEqual, "invalid4.json")
			So(files[4], ShouldEqual, "valid1.json")
		})

		Convey("ListDirectory should return folders", func() {
			files, err := ListDirectory("../samples")

			So(err, ShouldBeNil)
			So(files, ShouldNotBeNil)
			So(len(files), ShouldEqual, 2)
		})
	})
}
