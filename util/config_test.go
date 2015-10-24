package util_test

import (
	"testing"
  . "github.com/smartystreets/goconvey/convey"
  "gitlab.kohn.io/ankoh/vmlcm/util"
)

func TestConfig(t *testing.T) {
  Convey("Given various test files", t, func() {
    Convey("valid test files must be parsed correctly", func() {
      config1, err1 := util.ParseConfiguration("../samples/config/valid1.json")
      So(err1, ShouldBeNil)
      So(config1.Vmrun, ShouldEqual, "/Applications/VMware Fusion.app/Contents/Library/vmrun")
      So(config1.TemplatePath, ShouldEqual, "/tmp/samplevms/sample1.vmx")
      So(config1.ClonesDirectory, ShouldEqual, "/tmp/vmclones")
      So(len(config1.Addresses), ShouldEqual, 9)

      config2, err2 := util.ParseConfiguration("../samples/config/valid2.json")
      So(err2, ShouldBeNil)
      So(config2.Vmrun, ShouldEqual, "~/Applications/VMware Fusion.app/Contents/Library/vmrun")
      So(config2.TemplatePath, ShouldEqual, "./tmp/samplevms/sample2.vmx")
      So(config2.ClonesDirectory, ShouldEqual, "./tmp/vmclones2")
      So(len(config2.Addresses), ShouldEqual, 0)
    })

    Convey("invalid test files must result in an error", func() {
      config1, err1 := util.ParseConfiguration("../samples/config/invalid1.json")
      So(err1, ShouldNotBeNil)
      So(config1, ShouldBeNil)
      config2, err2 := util.ParseConfiguration("../samples/config/invalid2.json")
      So(err2, ShouldNotBeNil)
      So(config2, ShouldBeNil)
      config3, err3 := util.ParseConfiguration("../samples/config/invalid3.json")
      So(err3, ShouldNotBeNil)
      So(config3, ShouldBeNil)
      config4, err4 := util.ParseConfiguration("../samples/config/invalid4.json")
      So(err4, ShouldNotBeNil)
      So(config4, ShouldBeNil)
    })

    Convey("not existing file must result in an error", func() {
      config, err := util.ParseConfiguration("/path/not/exsting/hopefully")
      So(err, ShouldNotBeNil)
      So(config, ShouldBeNil)
    })
  })
}
