package controller

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
	. "github.com/smartystreets/goconvey/convey"
)

func TestVerification(t *testing.T) {
	Convey("Given a verification controller", t, func(context C) {
		// $REPO/samples/config

		Convey("isValidMacAddress should identify valid mac addresses", func() {
			So(isValidMacAddress("a1:b1:c1:d1:e1:f1"), ShouldBeTrue)
			So(isValidMacAddress("11:22:33:44:55:66"), ShouldBeTrue)
		})

		Convey("isValidMacAddress should detect invalid mac addresses", func() {
			So(isValidMacAddress(""), ShouldBeFalse)
			So(isValidMacAddress(":"), ShouldBeFalse)
			So(isValidMacAddress("a1:b1:c1:d1:e1"), ShouldBeFalse)
		})

		Convey("isAbsolutePath should identify valid absolute paths", func() {
			So(isAbsolutePath("/"), ShouldBeTrue)
			So(isAbsolutePath("/blabla/foo/32/bar.xyz/.4"), ShouldBeTrue)
		})

		Convey("isAbsolutePath should detect invalid absolute paths", func() {
			So(isAbsolutePath("./"), ShouldBeFalse)
			So(isAbsolutePath("foo"), ShouldBeFalse)
			So(isAbsolutePath(".foo"), ShouldBeFalse)
			So(isAbsolutePath("./blabla/foo/32/bar.xyz/.4"), ShouldBeFalse)
		})

		Convey("isValidPath should identify valid absolute paths", func() {
			So(isValidPath("/tmp"), ShouldBeTrue)
			So(isValidPath("/tmp/"), ShouldBeTrue)
		})

		Convey("isValidPath should detect invalid absolute paths", func() {
			So(isValidPath("./"), ShouldBeFalse)
			So(isValidPath("foo"), ShouldBeFalse)
			So(isValidPath("/not/existing/path"), ShouldBeFalse)
			So(isValidPath("../samples/config/valid1.json"), ShouldBeFalse)
		})
	})

	Convey("Verify should successfully verify various configurations", t, func() {
		logger := util.NewLogger()

		createTestVerifyFolders()
		createTestVerifyTemplate()
		createTestVerifyVmrun()
		defer deleteTestVerifyFolders()

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

		// Success
		err := Verify(logger, vmrun, config, true)
		So(err, ShouldBeNil)

		// Adding invalid MAC address
		config.Addresses = append(config.Addresses, "keine_valide_mac")
		err = Verify(logger, vmrun, config, true)
		So(err, ShouldNotBeNil)

		// Vmrun deletion
		config.Addresses = config.Addresses[:len(config.Addresses)-1]
		os.Remove("/tmp/vmlcmverify/vmrun")
		err = Verify(logger, vmrun, config, true)
		So(err, ShouldNotBeNil)

		// Template deletion
		createTestVerifyVmrun()
		os.Remove("/tmp/vmlcmverify/test.vmx")
		err = Verify(logger, vmrun, config, true)
		So(err, ShouldNotBeNil)

		// Clones directory deletion
		createTestVerifyTemplate()
		os.Remove("/tmp/vmlcmverify/clones")
		err = Verify(logger, vmrun, config, true)
		So(err, ShouldNotBeNil)

		// Invalid template file extension
		createTestVerifyFolders()
		os.Remove("/tmp/vmlcmverify/test.vmx")
		ioutil.WriteFile("/tmp/vmlcmverify/test", []byte(""), 0644)
		config.TemplatePath = "/tmp/vmlcmverify/test"
		err = Verify(logger, vmrun, config, true)
		So(err, ShouldNotBeNil)

		// File as clone folder
		os.Remove("/tmp/vmlcmverify/test")
		createTestVerifyTemplate()
		config.TemplatePath = "/tmp/vmlcmverify/test.vmx"
		os.Remove("/tmp/vmlcmverify/clones")
		ioutil.WriteFile("/tmp/vmlcmverify/clones", []byte(""), 0644)
		err = Verify(logger, vmrun, config, true)
		So(err, ShouldNotBeNil)
	})
}

func createTestVerifyFolders() {
	os.Mkdir("/tmp/vmlcmverify", 0755)
	os.Mkdir("/tmp/vmlcmverify/clones", 0755)
}

func createTestVerifyTemplate() {
	testBuffer := []byte("vmlcm test vmx\n")
	ioutil.WriteFile("/tmp/vmlcmverify/test.vmx", testBuffer, 0644)
}

func createTestVerifyVmrun() {
	testBuffer := []byte("vmlcm test vmrun\n")
	ioutil.WriteFile("/tmp/vmlcmverify/vmrun", testBuffer, 0755)
}

func deleteTestVerifyFolders() {
	os.RemoveAll("/tmp/vmlcmverify")
}
