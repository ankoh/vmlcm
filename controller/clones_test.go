package controller


import (
	"io/ioutil"
	"os"
	"testing"
  "strings"

	"github.com/ankoh/vmlcm/util"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClones(t *testing.T) {
	Convey("getAvailableMacAddresses should only return available MAC addresses", t, func() {
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

    config := new(util.LCMConfiguration)
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
