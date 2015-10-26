package controller

import (
	"testing"
  "fmt"
  "io/ioutil"
  "os"

  "gitlab.kohn.io/ankoh/vmlcm/vmware"
  "gitlab.kohn.io/ankoh/vmlcm/util"
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
    createTestFolders()
    createTestTemplate()
    createTestVmrun()
    defer deleteAll()

    vmrun := vmware.NewMockVmrun()
    config := new(util.LCMConfiguration)
    config.ClonesDirectory = "/tmp/vmlcm/clones/"
    config.TemplatePath = "/tmp/vmlcm/test.vmx"
    config.Vmrun = "/tmp/vmlcm/vmrun"

    // Success
    fmt.Println()
    fmt.Println()
    err := Verify(vmrun, config)
    fmt.Printf("\t%-55s", "")
    So(err, ShouldBeNil)
    fmt.Println()

    // Vmrun deletion
    fmt.Println("\t-- Deleting vmrun executable")
    os.Remove("/tmp/vmlcm/vmrun")
    err = Verify(vmrun, config)
    fmt.Printf("\t%-55s", "")
    So(err, ShouldNotBeNil)
    fmt.Println()

    // Template deletion
    fmt.Println("\t-- Restoring vmrun, deleting vmx template")
    createTestVmrun()
    os.Remove("/tmp/vmlcm/test.vmx")
    err = Verify(vmrun, config)
    fmt.Printf("\t%-55s", "")
    So(err, ShouldNotBeNil)
    fmt.Println()

    // Clones directory deletion
    fmt.Println("\t-- Restoring template, deleting clones directory")
    createTestTemplate()
    os.Remove("/tmp/vmlcm/clones")
    err = Verify(vmrun, config)
    fmt.Printf("\t%-55s", "")
    So(err, ShouldNotBeNil)
    fmt.Println()

    // Invalid template file extension
    fmt.Println("\t-- Restoring clones, adding invalid template")
    createTestFolders()
    os.Remove("/tmp/vmlcm/test.vmx")
    ioutil.WriteFile("/tmp/vmlcm/test", []byte(""), 0644)
    config.TemplatePath = "/tmp/vmlcm/test"
    err = Verify(vmrun, config)
    fmt.Printf("\t%-55s", "")
    So(err, ShouldNotBeNil)
    fmt.Println()

    // File as clone folder
    fmt.Println("\t-- Restoring template, using file as clone folder")
    os.Remove("/tmp/vmlcm/test")
    createTestTemplate()
    config.TemplatePath = "/tmp/vmlcm/test.vmx"
    os.Remove("/tmp/vmlcm/clones")
    ioutil.WriteFile("/tmp/vmlcm/clones", []byte(""), 0644)
    err = Verify(vmrun, config)
    fmt.Printf("\t%-55s", "")
    So(err, ShouldNotBeNil)
    fmt.Println()
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
