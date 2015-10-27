package controller

import (
	"testing"
	"os"

	"github.com/ankoh/vmlcm/vmware"
	"github.com/ankoh/vmlcm/util"
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

		Convey("getVMs should be able to detect all running vms and all vms in the clone directory", func() {
			// Create folders
			createTestStatusFolders()
			defer deleteTestStatusFolders()

			// Create configurations
			config := new(util.LCMConfiguration)
			config.Addresses = []string {
				"a1b1c1d1e1f1",
				"a2b2c2d2e2f2",
				"a3b3c3d3e3f3",
				"a4b4c4d4e4f4",
				"a5b5c5d5e5f5",
				"a6b6c6d6e6f6",
				"a7b7c7d7e7f7",
			}
			config.ClonesDirectory = "/tmp/vmlcmstatus/"
			config.Prefix = "pom2015"
			config.TemplatePath = "/tmp/vmlcmstatus/pom2015-template.vmwarevm/pom2015-template.vmx"

			// Now run getVMs and check for error first
			vms, err := getVMs(vmrun, config)

			So(err, ShouldBeNil)
			So(vms, ShouldNotBeNil)
			So(len(vms), ShouldEqual, 19)

			vmMap := make(map[string]*virtualMachine)

			// Build map with vms
			for _, vm := range vms {
				vmMap[vm.path] = vm
			}

			// First check all running vms
			runningPaths := []string {
				"/Volumes/VM_SB3/VMware/webbruegge.vmwarevm/webbruegge.vmx",
				"/Volumes/VM_SB3/VMware/dockerbruegge1.vmwarevm/dockerbruegge1.vmx",
				"/Volumes/VM_SB3/VMware/repoarchbruegge.vmwarevm/repoarchbruegge.vmx",
				"/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-5.vmwarevm/buildagent-mac-5.vmx",
				"/Volumes/VM_SB3/VMware/LS1Cloud.vmwarevm/LS1Cloud.vmx",
				"/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-6.vmwarevm/buildagent-mac-6.vmx",
				"/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-4.vmwarevm/buildagent-mac-4.vmx",
				"/Volumes/VM_SB3/VMware/backupbruegge.vmwarevm/backupbruegge.vmx",
				"/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-8.vmwarevm/buildagent-mac-8.vmx",
				"/Volumes/VM_SB3/VMware/build-agents-reserve/buildagent-mac-7.vmwarevm/buildagent-mac-7.vmx",
				"/Volumes/VM_SB3/VMware/mailbruegge.vmwarevm/mailbruegge.vmx",
				"/Volumes/VM_SB3/VMware/webbruegge_prelive.vmwarevm/webbruegge_prelive.vmx",
				"/Volumes/VM_SB3/VMware/monitorbruegge.vmwarevm/monitorbruegge.vmx",
				"/tmp/vmlcmstatus/pom2015-a1b1c1d1e1f1.vmwarevm/pom2015-a1b1c1d1e1f1.vmx",
				"/tmp/vmlcmstatus/pom2015-a2b2c2d2e2f2.vmwarevm/pom2015-a2b2c2d2e2f2.vmx",
				"/tmp/vmlcmstatus/pom2015-a3b3c3d3e3f3.vmwarevm/pom2015-a3b3c3d3e3f3.vmx",
			}

			for _, path := range runningPaths {
				vm, ok := vmMap[path]
				So(ok, ShouldBeTrue)
				So(vm, ShouldNotBeNil)
				So(vm.running, ShouldBeTrue)
				So(vm.template, ShouldBeFalse)
			}

			// Then check all clones
			clonePaths := []string {
				"/tmp/vmlcmstatus/pom2015-a1b1c1d1e1f1.vmwarevm/pom2015-a1b1c1d1e1f1.vmx",
				"/tmp/vmlcmstatus/pom2015-a2b2c2d2e2f2.vmwarevm/pom2015-a2b2c2d2e2f2.vmx",
				"/tmp/vmlcmstatus/pom2015-a3b3c3d3e3f3.vmwarevm/pom2015-a3b3c3d3e3f3.vmx",
				"/tmp/vmlcmstatus/pom2015-a4b4c4d4e4f4.vmwarevm/pom2015-a4b4c4d4e4f4.vmx",
				"/tmp/vmlcmstatus/pom2015-a5b5c5d5e5f5.vmwarevm/pom2015-a5b5c5d5e5f5.vmx",
			}

			for _, path := range clonePaths {
				vm, ok := vmMap[path]
				So(ok, ShouldBeTrue)
				So(vm, ShouldNotBeNil)
				So(vm.clone, ShouldBeTrue)
			}

			template, ok := vmMap["/tmp/vmlcmstatus/pom2015-template.vmwarevm/pom2015-template.vmx"]
			So(ok, ShouldBeTrue)
			So(template, ShouldNotBeNil)
			So(template.clone, ShouldBeFalse)
			So(template.running, ShouldBeFalse)
			So(template.template, ShouldBeTrue)
		})
	})
}

func createTestStatusFolders() {
	os.Mkdir("/tmp/vmlcmstatus", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-template.vmwarevm", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-template.vmwarevm/pom2015-template.vmx", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a1b1c1d1e1f1.vmwarevm", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a1b1c1d1e1f1.vmwarevm/pom2015-a1b1c1d1e1f1.vmx", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a2b2c2d2e2f2.vmwarevm", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a2b2c2d2e2f2.vmwarevm/pom2015-a2b2c2d2e2f2.vmx", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a3b3c3d3e3f3.vmwarevm", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a3b3c3d3e3f3.vmwarevm/pom2015-a3b3c3d3e3f3.vmx", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a4b4c4d4e4f4.vmwarevm", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a4b4c4d4e4f4.vmwarevm/pom2015-a4b4c4d4e4f4.vmx", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a5b5c5d5e5f5.vmwarevm", 0755)
	os.Mkdir("/tmp/vmlcmstatus/pom2015-a5b5c5d5e5f5.vmwarevm/pom2015-a5b5c5d5e5f5.vmx", 0755)
}

func deleteTestStatusFolders() {
	os.RemoveAll("/tmp/vmlcmstatus")
}
