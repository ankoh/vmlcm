package vmware

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExecute(t *testing.T) {
	Convey("Given a standard UNIX environment", t, func(context C) {
		// Should be the case...

		Convey("executeCommand should be able to execute 'ls -la ./'", func() {
			outChan := make(chan string)
			errChan := make(chan error)
			defer close(outChan)
			defer close(errChan)

			go executeCommand(outChan, errChan, "ls", "-la", "./")

			// Chooses either out or err
			select {
			case out := <-outChan:
				So(out, ShouldNotBeNil)
			case err := <-errChan:
				So(err, ShouldBeNil)
			}
		})

		Convey("executeCommand should throw an error when executing 'notexistingcommand -foo -bar 42'", func() {
			outChan := make(chan string)
			errChan := make(chan error)
			defer close(outChan)
			defer close(errChan)

			go executeCommand(outChan, errChan, "notexistingcommand", "-foo", "-bar", "42")

			// Chooses either out or err
			select {
			case out := <-outChan:
				So(out, ShouldBeNil)
			case err := <-errChan:
				So(err, ShouldNotBeNil)
			}
		})
	})
}
