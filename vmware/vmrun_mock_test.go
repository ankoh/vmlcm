package vmware_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ankoh/vmlcm/vmware"
)

func TestExecute(t *testing.T) {
	Convey("Given a mock vmrun wrapper", t, func(context C) {
		vmrun := vmware.NewMockVmrun()
		defer vmrun.Close()

		Convey("Help() must return usage information and the version number", func() {
			outChannel := vmrun.GetOutputChannel()
			errChannel := vmrun.GetErrorChannel()

			go vmrun.Help()

			select {
			case out := <-outChannel:
				So(out, ShouldNotBeNil)
        So(out, ShouldContainSubstring, "vmrun version 1.14.2 build-2779224")
			case err := <-errChannel:
				So(err, ShouldBeNil)
			}
		})

	})
}
