package controller

import (
	"os"
	"testing"

	"github.com/ankoh/vmlcm/vmware"
	. "github.com/smartystreets/goconvey/convey"
)

func TestVMs(t *testing.T) {
	Convey("Given a mocked Vmrun Wrapper", t, func() {
		vmrun := vmware.NewMockVmrun()
		vmrun.RunningVMs = append(vmrun.RunningVMs, "/Volumes/VM_SB3/VMware/webbruegge.vmwarevm/webbruegge.vmx")
		vmrun.RunningVMs = append(vmrun.RunningVMs, "/Volumes/VM_SB3/VMware/dockerbruegge1.vmwarevm/dockerbruegge1.vmx")
		vmrun.RunningVMs = append(vmrun.RunningVMs, "/Volumes/VM_SB3/VMware/repoarchbruegge.vmwarevm/repoarchbruegge.vmx")

		Convey("getRunningVMPaths should be able to parse all running vm paths", func() {
			result, err := getRunningVMPaths(vmrun)

			So(err, ShouldBeNil)
			So(len(result), ShouldEqual, 3)
			So(result[0], ShouldEqual, "/Volumes/VM_SB3/VMware/webbruegge.vmwarevm/webbruegge.vmx")
			So(result[1], ShouldEqual, "/Volumes/VM_SB3/VMware/dockerbruegge1.vmwarevm/dockerbruegge1.vmx")
			So(result[2], ShouldEqual, "/Volumes/VM_SB3/VMware/repoarchbruegge.vmwarevm/repoarchbruegge.vmx")
		})

		Convey("discoverVMsInDirectory should be able to discover vmx files in a directory", func() {
			createTestVMsVmwareFolder()
			defer deleteTestVMsVmwareFolder()

			result, err := discoverVMs("/tmp/vmlcmvms/")

			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			So(len(result), ShouldEqual, 2)
			So(result[0], ShouldEqual, "/tmp/vmlcmvms/test1.vmwarevm/test1.vmx")
			So(result[1], ShouldEqual, "/tmp/vmlcmvms/test2.vmwarevm/test2.vmx")
		})
	})
}

func createTestVMsVmwareFolder() {
	os.Mkdir("/tmp/vmlcmvms/", 0755)
	os.Mkdir("/tmp/vmlcmvms/test1.vmwarevm", 0755)
	os.Mkdir("/tmp/vmlcmvms/test1.vmwarevm/test1.vmx", 0755)
	os.Mkdir("/tmp/vmlcmvms/test2.vmwarevm", 0755)
	os.Mkdir("/tmp/vmlcmvms/test2.vmwarevm/test2.vmx", 0755)
}

func deleteTestVMsVmwareFolder() {
	os.RemoveAll("/tmp/vmlcmvms/")
}
