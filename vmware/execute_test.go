package vmware_test

import (
	"testing"
  . "github.com/smartystreets/goconvey/convey"
  "gitlab.kohn.io/ankoh/vmlcm/vmware"
)

func TestExecute(t *testing.T) {
  Convey("Given a standard UNIX environment", t, func(context C) {
    // Should be the case...

    Convey("executeCommand must be able to execute 'ls -la ./'", func() {
      outChan := make(chan string)
      errChan := make(chan string)
      defer close(outChan)
      defer close(errChan)

      go vmware.ExecuteCommand(outChan, errChan, "ls", "-la", "./")

			// Select different channels
      select {
        case out := <- outChan:
        So(out, ShouldNotBeNil)
				case err := <- errChan:
				So(err, ShouldBeNil)
      }
    })
  })
}
