package controller

import (
	"testing"
  "io/ioutil"
  "os"

  "github.com/ankoh/vmlcm/vmware"
  "github.com/ankoh/vmlcm/util"
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
			So(isValidPath("/not/existing/path"), ShouldBeFalse)
			So(isValidPath("../samples/config/valid1.json"), ShouldBeFalse)
		})
	})

  Convey("Verify must successfully verify various configurations", t, func() {
    logger := util.NewLogger()

    createTestFolders()
    createTestTemplate()
    createTestVmrun()
    defer deleteAll()

    vmrun := vmware.NewMockVmrun()
    config := new(util.LCMConfiguration)
    config.ClonesDirectory = "/tmp/vmlcm/clones/"
    config.TemplatePath = "/tmp/vmlcm/test.vmx"
    config.Vmrun = "/tmp/vmlcm/vmrun"
		config.Addresses = []string {
			"a1:b1:c1:d1:e1:f1",
			"a2:b2:c2:d2:e2:f2",
		}

    // Success
    err := Verify(logger, vmrun, config, true)
    So(err, ShouldBeNil)

		// Adding invalid MAC address
		config.Addresses = append(config.Addresses, "keine_valide_mac")
    err = Verify(logger, vmrun, config, true)
    So(err, ShouldNotBeNil)

    // Vmrun deletion
		config.Addresses = config.Addresses[:len(config.Addresses)-1]
    os.Remove("/tmp/vmlcm/vmrun")
    err = Verify(logger, vmrun, config, true)
    So(err, ShouldNotBeNil)

    // Template deletion
    createTestVmrun()
    os.Remove("/tmp/vmlcm/test.vmx")
    err = Verify(logger, vmrun, config, true)
    So(err, ShouldNotBeNil)

    // Clones directory deletion
    createTestTemplate()
    os.Remove("/tmp/vmlcm/clones")
    err = Verify(logger, vmrun, config, true)
    So(err, ShouldNotBeNil)

    // Invalid template file extension
    createTestFolders()
    os.Remove("/tmp/vmlcm/test.vmx")
    ioutil.WriteFile("/tmp/vmlcm/test", []byte(""), 0644)
    config.TemplatePath = "/tmp/vmlcm/test"
    err = Verify(logger, vmrun, config, true)
    So(err, ShouldNotBeNil)

    // File as clone folder
    os.Remove("/tmp/vmlcm/test")
    createTestTemplate()
    config.TemplatePath = "/tmp/vmlcm/test.vmx"
    os.Remove("/tmp/vmlcm/clones")
    ioutil.WriteFile("/tmp/vmlcm/clones", []byte(""), 0644)
    err = Verify(logger, vmrun, config, true)
    So(err, ShouldNotBeNil)
  })
}

func createTestFolders() {
  os.Mkdir("/tmp/vmlcm", 0755)
  os.Mkdir("/tmp/vmlcm/clones", 0755)
}

func createTestTemplate() {
  testBuffer := []byte("vmlcm test vmx\n")
  ioutil.WriteFile("/tmp/vmlcm/test.vmx", testBuffer, 0644)
}

func createTestVmrun() {
  testBuffer := []byte("vmlcm test vmrun\n")
  ioutil.WriteFile("/tmp/vmlcm/vmrun", testBuffer, 0755)
}

func deleteAll() {
  os.RemoveAll("/tmp/vmlcm")
}
