package controller


import (
	"io/ioutil"
	"os"
	"testing"
  "strings"
	"fmt"

	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClones(t *testing.T) {
	Convey("Given a mocked Vmrun Wrapper", t, func() {
		// Prepare mock vmrun
		vmrun := vmware.NewMockVmrun()
		vmrun.TemplateVM = "/tmp/vmlcmclones/pom2015-template.vmwarevm/pom2015-template.vmx"
		vm1 := new(virtualMachine)
		vm2 := new(virtualMachine)
		vm3 := new(virtualMachine)
		vm1.path = "/foo/bar/pom2015-A1B1C1D1E1F1.vmwarevm/pom2015-A1B1C1D1E1F1.vmx"
		vm2.path = "/foo/bar/pom2015-A2B2C2D2E2F2.vmwarevm/pom2015-A2B2C2D2E2F2.vmx"
		vm3.path = "/foo/bar/pom2015-A3B3C3D3E3F3.vmwarevm/pom2015-A3B3C3D3E3F3.vmx"
		clones := []*virtualMachine {
			vm1,
			vm2,
			vm3,
		}
		vmrun.CloneFolderVMs = []string {
			"/foo/bar/pom2015-A1B1C1D1E1F1.vmwarevm/pom2015-A1B1C1D1E1F1.vmx",
			"/foo/bar/pom2015-A2B2C2D2E2F2.vmwarevm/pom2015-A2B2C2D2E2F2.vmx",
			"/foo/bar/pom2015-A3B3C3D3E3F3.vmwarevm/pom2015-A3B3C3D3E3F3.vmx",
		}

		// Prepare config
		config := new(util.LCMConfiguration)
		config.Vmrun = "/Applications/VMware Fusion.app/Contents/Library/vmrun"
		config.Prefix = "pom2015"
		config.ClonesDirectory = "/tmp/vmlcmclones/"
		config.TemplatePath = "/tmp/vmlcmclones/pom2015-template.vmwarevm/pom2015-template.vmx"
		config.Vmrun = ""
		config.Addresses = []string {
			"a1:b1:c1:d1:e1:f1",
			"a2:b2:c2:d2:e2:f2",
			"a3:b3:c3:d3:e3:f3",
			"a4:b4:c4:d4:e4:f4",
			"a5:b5:c5:d5:e5:f5",
		}
		for i, address := range config.Addresses {
			config.Addresses[i] = strings.ToUpper(address)
		}

		// Check mac addresses
		available := getAvailableMacAddresses(clones, config)
		So(available, ShouldNotBeNil)
		So(len(available), ShouldEqual, 2)

		// Test getAvailableMacAddresses
		Convey("getAvailableMacAddresses should only return available MAC addresses", func() {
	    // 2 available
	    available := getAvailableMacAddresses(clones, config)
	    So(available, ShouldNotBeNil)
	    So(len(available), ShouldEqual, 2)

	    // 1 available
	    config.Addresses = config.Addresses[0:4]
	    available = getAvailableMacAddresses(clones, config)
	    So(len(available), ShouldEqual, 1)

	    // 0 available
	    config.Addresses = config.Addresses[0:3]
	    available = getAvailableMacAddresses(clones, config)
	    So(len(available), ShouldEqual, 0)

	    // -1 available
	    config.Addresses = config.Addresses[0:2]
	    available = getAvailableMacAddresses(clones, config)
	    So(len(available), ShouldEqual, 0)

	    // None provided
	    config.Addresses = []string{}
	    available = getAvailableMacAddresses(clones, config)
	    So(len(available), ShouldEqual, 0)
		})

		Convey("cloneUpTo should clone no VM if up < existing", func() {
			created, err := cloneUpTo(vmrun, config, clones,3)
			So(err, ShouldBeNil)
			So(created, ShouldNotBeNil)
			So(len(created), ShouldEqual, 0)
			created, err = cloneUpTo(vmrun, config, clones, 2)
			So(err, ShouldBeNil)
			So(created, ShouldNotBeNil)
			So(len(created), ShouldEqual, 0)
			created, err = cloneUpTo(vmrun, config, clones, 1)
			So(err, ShouldBeNil)
			So(created, ShouldNotBeNil)
			So(len(created), ShouldEqual, 0)
			created, err = cloneUpTo(vmrun, config, clones, 0)
			So(err, ShouldBeNil)
			So(created, ShouldNotBeNil)
			So(len(created), ShouldEqual, 0)
			created, err = cloneUpTo(vmrun, config, clones, -10)
			So(err, ShouldBeNil)
			So(created, ShouldNotBeNil)
			So(len(created), ShouldEqual, 0)
		})

		Convey("cloneUpTo should correctly create 1 clone if allowed", func() {
			So(len(vmrun.TemplateSnapshots), ShouldEqual, 0)
			created, err := cloneUpTo(vmrun, config, clones, 4)
			So(err, ShouldBeNil)

			// Check cloneUpTo result array
			So(created, ShouldNotBeNil)
			So(len(created), ShouldEqual, 1)
			So(created[0], ShouldEqual, fmt.Sprintf(
				"%s-%s", config.Prefix, "A4B4C4D4E4F4"))

			// Check vmrun CloneFolder
			So(len(vmrun.CloneFolderVMs), ShouldEqual, 4)
			path := fmt.Sprintf("/tmp/vmlcmclones/%s-%s.vmwarevm/%s-%s.vmx",
				config.Prefix, "A4B4C4D4E4F4", config.Prefix, "A4B4C4D4E4F4")
			So(vmrun.CloneFolderVMs[3], ShouldEqual, path)

			// Check created snapshot
			So(vmrun.TemplateSnapshots, ShouldNotBeNil)
			So(len(vmrun.TemplateSnapshots), ShouldEqual, 1)
			So(vmrun.TemplateSnapshots[0], ShouldContainSubstring, config.Prefix)
		})

		Convey("cloneUpTo should correctly create 2 clones if allowed", func() {
			created, err := cloneUpTo(vmrun, config, clones, 5)
			So(err, ShouldBeNil)

			// Check cloneUpTo result array
			So(created, ShouldNotBeNil)
			So(len(created), ShouldEqual, 2)
			So(created[0], ShouldEqual, fmt.Sprintf(
				"%s-%s", config.Prefix, "A4B4C4D4E4F4"))
			So(created[1], ShouldEqual, fmt.Sprintf(
				"%s-%s", config.Prefix, "A5B5C5D5E5F5"))

			// Check vmrun CloneFolder
			path := fmt.Sprintf("/tmp/vmlcmclones/%s-%s.vmwarevm/%s-%s.vmx",
				config.Prefix, "A4B4C4D4E4F4", config.Prefix, "A4B4C4D4E4F4")
			So(vmrun.CloneFolderVMs[3], ShouldEqual, path)
			path = fmt.Sprintf("/tmp/vmlcmclones/%s-%s.vmwarevm/%s-%s.vmx",
				config.Prefix, "A5B5C5D5E5F5", config.Prefix, "A5B5C5D5E5F5")
			So(vmrun.CloneFolderVMs[4], ShouldEqual, path)

			// Check created snapshot
			So(vmrun.TemplateSnapshots, ShouldNotBeNil)
			So(len(vmrun.TemplateSnapshots), ShouldEqual, 1)
			So(vmrun.TemplateSnapshots[0], ShouldContainSubstring, config.Prefix)
		})

		Convey("cloneUpTo should create only 2 clones even though 3 are demanded", func() {
			created, err := cloneUpTo(vmrun, config, clones, 6)
			So(err, ShouldBeNil)

			// Check cloneUpTo result array
			So(created, ShouldNotBeNil)
			So(len(created), ShouldEqual, 2)
			So(created[0], ShouldEqual, fmt.Sprintf(
				"%s-%s", config.Prefix, "A4B4C4D4E4F4"))
			So(created[1], ShouldEqual, fmt.Sprintf(
				"%s-%s", config.Prefix, "A5B5C5D5E5F5"))

			// Check vmrun CloneFolder
			path := fmt.Sprintf("/tmp/vmlcmclones/%s-%s.vmwarevm/%s-%s.vmx",
				config.Prefix, "A4B4C4D4E4F4", config.Prefix, "A4B4C4D4E4F4")
			So(vmrun.CloneFolderVMs[3], ShouldEqual, path)
			path = fmt.Sprintf("/tmp/vmlcmclones/%s-%s.vmwarevm/%s-%s.vmx",
				config.Prefix, "A5B5C5D5E5F5", config.Prefix, "A5B5C5D5E5F5")
			So(vmrun.CloneFolderVMs[4], ShouldEqual, path)

			// Check created snapshot
			So(vmrun.TemplateSnapshots, ShouldNotBeNil)
			So(len(vmrun.TemplateSnapshots), ShouldEqual, 1)
			So(vmrun.TemplateSnapshots[0], ShouldContainSubstring, config.Prefix)
		})

		Convey("cloneUpTo should create a snapshot if only invalid snapshots are existing", func() {
			vmrun.TemplateSnapshots = []string {
				"existing-but-wrong-snapshot-1",
				"existing-but-wrong-snapshot-2",
			}

			_, err := cloneUpTo(vmrun, config, clones, 5)
			So(err, ShouldBeNil)

			// Check created snapshot
			So(vmrun.TemplateSnapshots, ShouldNotBeNil)
			So(len(vmrun.TemplateSnapshots), ShouldEqual, 3)
			So(vmrun.TemplateSnapshots[2], ShouldContainSubstring, config.Prefix)
		})

		Convey("cloneUpTo should not create a snapshot if a valid exists", func() {
			vmrun.TemplateSnapshots = []string {
				"existing-but-wrong-snapshot-1",
				"existing-but-wrong-snapshot-2",
				fmt.Sprintf("%s-1234567", config.Prefix),
			}
			_, err := cloneUpTo(vmrun, config, clones, 5)
			So(err, ShouldBeNil)
			So(vmrun.TemplateSnapshots, ShouldNotBeNil)
			So(len(vmrun.TemplateSnapshots), ShouldEqual, 3)
		})
	})
}

func createTestClonesFolders() {
	os.Mkdir("/tmp/vmlcmclones", 0755)
	os.Mkdir("/tmp/vmlcmclones/clones", 0755)
}

func createTestClonesTemplate() {
	testBuffer := []byte("vmlcm test vmx\n")
	ioutil.WriteFile("/tmp/vmlcmclones/test.vmx", testBuffer, 0644)
}

func createTestClonesVmrun() {
	testBuffer := []byte("vmlcm test vmrun\n")
	ioutil.WriteFile("/tmp/vmlcmclones/vmrun", testBuffer, 0755)
}

func deleteTestClonesFolders() {
	os.RemoveAll("/tmp/vmlcmclones")
}
