package controller

import (
	"testing"

	"github.com/ankoh/vmlcm/vmware"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStatus(t *testing.T) {
	Convey("Given a mocked Vmrun Wrapper", t, func() {
		vmrun := vmware.NewMockVmrun()

		Convey("getVmrunVersion should be able to parse the version information", func() {
			result, err := getVmrunVersion(vmrun)

			So(err, ShouldBeNil)
			So(result.version, ShouldEqual, "1.14.2")
			So(result.build, ShouldEqual, "2779224")
		})

		Convey("getRunningVMNumber should be able to parse the number of running vms", func() {
			result, err := getRunningVMNumber(vmrun)

			So(err, ShouldBeNil)
			So(result, ShouldEqual, 13)
		})
	})
}
