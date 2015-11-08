package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
  "os"
  "io/ioutil"
)

func TestVMX(t *testing.T) {
  Convey("Given precompiled regexs", t, func() {
    Convey("eth0ConnectionTypePattern should match a correct string", func() {
      natString := "ethernet0.connectionType = \"nat\""
      natBytes := []byte(natString)
      bridgedString := "ethernet0.connectionType = \"bridged\""
      bridgedBytes := []byte(bridgedString)

      So(eth0ConnectionTypePattern.MatchString(natString), ShouldBeTrue)
      So(eth0ConnectionTypePattern.Match(natBytes), ShouldBeTrue)
      So(eth0ConnectionTypePattern.MatchString(bridgedString), ShouldBeTrue)
      So(eth0ConnectionTypePattern.Match(bridgedBytes), ShouldBeTrue)
    })

    Convey("eth0ConnectionTypePattern should replace correctly", func() {
      natString := "XYZethernet0.connectionType = \"nat\"ABC"
      natString = eth0ConnectionTypePattern.ReplaceAllString(
        natString, "ethernet0.connectionType = \"bridged\"")
      So(natString, ShouldEqual, "XYZethernet0.connectionType = \"bridged\"ABC")
    })

    Convey("eth0AddressPattern should match a correct string", func() {
      testString := "ethernet0.address = \"00:50:56:27:C4:66\""
      testBytes := []byte(testString)

      So(eth0AddressPattern.MatchString(testString), ShouldBeTrue)
      So(eth0AddressPattern.Match(testBytes), ShouldBeTrue)
    })

    Convey("eth0AddressPattern should replace correctly", func() {
      testString := "XYZethernet0.address = \"00:50:56:27:C4:66\"ABC"
      testString = eth0AddressPattern.ReplaceAllString(testString,
        "ethernet0.address = \"11:22:33:44:55:66\"")
      So(testString, ShouldEqual, "XYZethernet0.address = \"11:22:33:44:55:66\"ABC")
    })
  })

	Convey("Given various test files", t, func() {
    createTestVMXFolders()
    defer deleteTestVerifyFolders()

    Convey("bridged.txt should be updated correctly", func() {
      err := UpdateVMX(
        "../samples/vmx/bridged.txt",
        "/tmp/vmlcmvmx/bridged.txt",
        "A1:B1:C1:D1:E1:F1")
      So(err, ShouldBeNil)

      // Check created file
      fileBytes, fileErr := ioutil.ReadFile("/tmp/vmlcmvmx/bridged.txt")
      So(fileErr, ShouldBeNil)
      fileString := string(fileBytes)
      So(fileString, ShouldContainSubstring,
        "ethernet0.connectionType = \"bridged\"")
      So(fileString, ShouldContainSubstring,
         "ethernet0.address = \"A1:B1:C1:D1:E1:F1\"")
    })

    Convey("nat.txt should be updated correctly", func() {
      err := UpdateVMX(
        "../samples/vmx/nat.txt",
        "/tmp/vmlcmvmx/nat.txt",
        "A1:B1:C1:D1:E1:F1")
      So(err, ShouldBeNil)

      // Check created file
      fileBytes, fileErr := ioutil.ReadFile("/tmp/vmlcmvmx/nat.txt")
      So(fileErr, ShouldBeNil)
      fileString := string(fileBytes)
      So(fileString, ShouldContainSubstring,
        "ethernet0.connectionType = \"bridged\"")
      So(fileString, ShouldContainSubstring,
         "ethernet0.address = \"A1:B1:C1:D1:E1:F1\"")
    })

    Convey("natMacLess.txt should be updated correctly", func() {
      err := UpdateVMX(
        "../samples/vmx/natMacLess.txt",
        "/tmp/vmlcmvmx/natMacLess.txt",
        "A1:B1:C1:D1:E1:F1")
      So(err, ShouldBeNil)

      // Check created file
      fileBytes, fileErr := ioutil.ReadFile("/tmp/vmlcmvmx/natMacLess.txt")
      So(fileErr, ShouldBeNil)
      fileString := string(fileBytes)
      So(fileString, ShouldContainSubstring,
        "ethernet0.connectionType = \"bridged\"")
      So(fileString, ShouldContainSubstring,
         "ethernet0.address = \"A1:B1:C1:D1:E1:F1\"")
    })
	})
}

func createTestVMXFolders() {
  os.Mkdir("/tmp/vmlcmvmx", 0755)
}

func deleteTestVerifyFolders() {
	os.RemoveAll("/tmp/vmlcmvmx")
}
