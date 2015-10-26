package controller

import (
	"testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestVerification(t *testing.T) {
  Convey("Given a verification controller", t, func(context C) {
    // $REPO/samples/config

    Convey("isValidMacAddress must identify valid mac addresses", func() {
      So(isValidMacAddress("a1:b1:c1:d1:e1:f1"), ShouldBeTrue)
      So(isValidMacAddress("11:22:33:44:55:66"), ShouldBeTrue)
    })

    Convey("isValidMacAddress must detect invalid mac addresses", func() {
      So(isValidMacAddress(""), ShouldBeFalse)
      So(isValidMacAddress(":"), ShouldBeFalse)
      So(isValidMacAddress("a1:b1:c1:d1:e1"), ShouldBeFalse)
    })

    Convey("isAbsolutePath must identify valid absolute paths", func() {
      So(isAbsolutePath("/"), ShouldBeTrue)
      So(isAbsolutePath("/blabla/foo/32/bar.xyz/.4"), ShouldBeTrue)
    })

    Convey("isAbsolutePath must detect invalid absolute paths", func() {
      So(isAbsolutePath("./"), ShouldBeFalse)
      So(isAbsolutePath("foo"), ShouldBeFalse)
      So(isAbsolutePath(".foo"), ShouldBeFalse)
      So(isAbsolutePath("./blabla/foo/32/bar.xyz/.4"), ShouldBeFalse)
    })

    Convey("isValidPath must identify valid absolute paths", func() {
      So(isValidPath("/tmp"), ShouldBeTrue)
      So(isValidPath("/tmp/"), ShouldBeTrue)
    })

    Convey("isValidPath must detect invalid absolute paths", func() {
      So(isValidPath("./"), ShouldBeFalse)
      So(isValidPath("foo"), ShouldBeFalse)
      So(isValidPath("../samples/config/valid1.json"), ShouldBeFalse)
    })
  })
}
