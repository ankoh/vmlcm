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

		Convey("getRunningVMPaths should be able to parse all running vm paths", func() {
			result, err := getRunningVMPaths(vmrun)

			So(err, ShouldBeNil)
			So(len(result), ShouldEqual, 13)
			So(result[0], ShouldEqual, "/Volumes/VM_SB3/VMware/webbruegge.vmwarevm/webbruegge.vmx")
			So(result[1], ShouldEqual, "/Volumes/VM_SB3/VMware/dockerbruegge1.vmwarevm/dockerbruegge1.vmx")
			So(result[2], ShouldEqual, "/Volumes/VM_SB3/VMware/repoarchbruegge.vmwarevm/repoarchbruegge.vmx")
			So(result[3], ShouldEqual, "/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-5.vmwarevm/buildagent-mac-5.vmx")
			So(result[4], ShouldEqual, "/Volumes/VM_SB3/VMware/LS1Cloud.vmwarevm/LS1Cloud.vmx")
			So(result[5], ShouldEqual, "/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-6.vmwarevm/buildagent-mac-6.vmx")
			So(result[6], ShouldEqual, "/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-4.vmwarevm/buildagent-mac-4.vmx")
			So(result[7], ShouldEqual, "/Volumes/VM_SB3/VMware/backupbruegge.vmwarevm/backupbruegge.vmx")
			So(result[8], ShouldEqual, "/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-8.vmwarevm/buildagent-mac-8.vmx")
			So(result[9], ShouldEqual, "/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-7.vmwarevm/buildagent-mac-7.vmx")
			So(result[10], ShouldEqual, "/Volumes/VM_SB3/VMware/mailbruegge.vmwarevm/mailbruegge.vmx")
			So(result[11], ShouldEqual, "/Volumes/VM_SB3/VMware/webbruegge_prelive.vmwarevm/webbruegge_prelive.vmx")
			So(result[12], ShouldEqual, "/Volumes/VM_SB3/VMware/monitorbruegge.vmwarevm/monitorbruegge.vmx")
		})

		Convey("discoverVMsInDirectory should be able to discover vmx files in a directory", func() {
			createTestVmwareFolder()
			defer deleteTestVmwareFolder()

			result, err := discoverVMs("/tmp/")

			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			So(len(result), ShouldEqual, 2)
			So(result[0], ShouldEqual, "/tmp/test1.vmwarevm/test1.vmx")
			So(result[1], ShouldEqual, "/tmp/test2.vmwarevm/test2.vmx")
		})
	})
}

func createTestVmwareFolder() {
	os.Mkdir("/tmp/test1.vmwarevm", 0755)
	os.Mkdir("/tmp/test1.vmwarevm/test1.vmx", 0755)
	os.Mkdir("/tmp/test2.vmwarevm", 0755)
	os.Mkdir("/tmp/test2.vmwarevm/test2.vmx", 0755)
}

func deleteTestVmwareFolder() {
	os.RemoveAll("/tmp/test1.vmwarevm")
	os.RemoveAll("/tmp/test2.vmwarevm")
}
