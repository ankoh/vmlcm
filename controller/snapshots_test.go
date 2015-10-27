package controller

import (
	"testing"

	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSnapshots(t *testing.T) {
	Convey("Given a mocked Vmrun Wrapper", t, func() {
		vmrun := vmware.NewMockVmrun()
    config := new(util.LCMConfiguration)
    config.ClonesDirectory = "/tmp/vmlcmverify/clones/"
    config.TemplatePath = "/tmp/vmlcmverify/test.vmx"
    config.Vmrun = "/tmp/vmlcmverify/vmrun"
    config.Prefix = "Pom2015"
    config.Addresses = []string{
      "a1:b1:c1:d1:e1:f1",
      "a2:b2:c2:d2:e2:f2",
    }

		Convey("getTemplateSnapshots should be able to parse the template snapshot names", func() {
			snapshots, err := getTemplateSnapshots(vmrun, config)
			So(err, ShouldBeNil)
      So(snapshots, ShouldNotBeNil)
      So(len(snapshots), ShouldEqual, 2)
      So(snapshots[0], ShouldEqual, "SetUp")
      So(snapshots[1], ShouldEqual, "POM 2015")
		})

    Convey("createTemplateSnapshot should be able to create a snapshot", func() {
			_, err := createTemplateSnapshot(vmrun, config)
			So(err, ShouldBeNil)
		})
  })
}
