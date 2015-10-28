package vmware_test

import (
	"testing"

	"github.com/ankoh/vmlcm/vmware"
	. "github.com/smartystreets/goconvey/convey"
)

func TestExecute(t *testing.T) {
	Convey("Given a mock vmrun wrapper", t, func(context C) {
		vmrun := vmware.NewMockVmrun()

		Convey("Help() should return usage information and the version number", func() {

			out, err := vmrun.Help()
			So(out, ShouldNotBeNil)
			So(out, ShouldContainSubstring, "vmrun version 1.14.2 build-2779224")
			So(err, ShouldBeNil)
		})

		Convey("List() should return running vms", func() {

			out, err := vmrun.List()
			So(out, ShouldNotBeNil)
			So(out, ShouldContainSubstring, "/Volumes/VM_SB3/VMware/webbruegge.vmwarevm/webbruegge.vmx")
			So(err, ShouldBeNil)
		})

	})
}
