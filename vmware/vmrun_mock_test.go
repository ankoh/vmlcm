package vmware_test

import (
	"testing"

	"github.com/ankoh/vmlcm/vmware"
	. "github.com/smartystreets/goconvey/convey"
)

func TestExecute(t *testing.T) {
	Convey("Given a mock vmrun wrapper", t, func(context C) {
		vmrun := vmware.NewMockVmrun()
		vmrun.RunningVMs = append(vmrun.RunningVMs, "/Volumes/VM_SB3/VMware/webbruegge.vmwarevm/webbruegge.vmx")

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


		Convey("CloneLinked() should correctly clone in a clean environment", func() {
			vmrun.TemplateVM = "/tmp/vmlcmclones/pom2015-template.vmwarevm/pom2015-template.vmx"
			vmrun.RunningVMs = []string{}
			vmrun.CloneFolderVMs = []string{}
			vmrun.TemplateSnapshots = []string {
				"existing-snapshot-1",
				"existing-snapshot-2",
			}
			vmrun.CloneLinked(
				"/tmp/vmlcmclones/pom2015-template.vmwarevm/pom2015-template.vmx",
				"/tmp/vmlcmclones/",
				"pom2015-12345678",
				"existing-snapshot-1")
			So(vmrun.CloneFolderVMs, ShouldNotBeNil)
			So(len(vmrun.CloneFolderVMs), ShouldEqual, 1)
			So(vmrun.CloneFolderVMs[0], ShouldEqual,
				"/tmp/vmlcmclones/pom2015-12345678.vmwarevm/pom2015-12345678.vmx")
		})

		Convey("CloneLinked() should return an error if vm exists", func() {
			vmrun.TemplateVM = "/tmp/vmlcmclones/pom2015-template.vmwarevm/pom2015-template.vmx"
			vmrun.RunningVMs = []string{}
			vmrun.CloneFolderVMs = []string{
				"/tmp/vmlcmclones/pom2015-12345678.vmwarevm/pom2015-12345678.vmx",
			}
			vmrun.TemplateSnapshots = []string {
				"existing-snapshot-1",
				"existing-snapshot-2",
			}
			err := vmrun.CloneLinked(
				"/tmp/vmlcmclones/pom2015-template.vmwarevm/pom2015-template.vmx",
				"/tmp/vmlcmclones/",
				"pom2015-12345678",
				"existing-snapshot-1")
			So(err, ShouldNotBeNil)
		})

		Convey("CloneLinked() should return an error if snapshot does not exist", func() {
			vmrun.TemplateVM = "/tmp/vmlcmclones/pom2015-template.vmwarevm/pom2015-template.vmx"
			vmrun.RunningVMs = []string{}
			vmrun.CloneFolderVMs = []string{}
			vmrun.TemplateSnapshots = []string {
				"existing-snapshot-1",
				"existing-snapshot-2",
			}
			err := vmrun.CloneLinked(
				"/tmp/vmlcmclones/pom2015-template.vmwarevm/pom2015-template.vmx",
				"/tmp/vmlcmclones/",
				"pom2015-12345678",
				"not-existing-snapshot")
			So(err, ShouldNotBeNil)
		})

		Convey("CloneLinked() should return an error if template does not exist", func() {
			vmrun.TemplateVM = "/tmp/vmlcmclones/pom2015-template.vmwarevm/pom2015-template.vmx"
			vmrun.RunningVMs = []string{}
			vmrun.CloneFolderVMs = []string{}
			vmrun.TemplateSnapshots = []string {
				"existing-snapshot-1",
				"existing-snapshot-2",
			}
			err := vmrun.CloneLinked(
				"/not/existing",
				"/tmp/vmlcmclones/",
				"pom2015-12345678",
				"not-existing-snapshot")
			So(err, ShouldNotBeNil)
		})
	})
}
