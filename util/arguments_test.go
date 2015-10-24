package util

import (
	"os"
	"testing"
  "flag"
  . "github.com/smartystreets/goconvey/convey"
)

func TestArguments(t *testing.T) {
	Convey("Given a clean command line environment", t, func() {
    flag.CommandLine = flag.NewFlagSet("vmlcm", flag.ContinueOnError)

		Convey("When the flag value is separated with whitespace", func() {
      Convey("ParseArguments should return correct values", func() {
        os.Args = []string{
          "vmlcm",
          "-f", "./agents.json",
          "verify"}
        args, err := ParseArguments()

        So(err, ShouldEqual, nil)
        So(args, ShouldNotEqual, nil)
        So(args.ConfigPath, ShouldNotBeNil)
        So(*args.ConfigPath, ShouldEqual, "./agents.json")
        So(args.Command, ShouldEqual, VerifyCommand)
      })
		})

    Convey("When the flag value is separated with an equal sign", func() {
      Convey("ParseArguments should return correct values", func() {
        os.Args = []string{
          "vmlcm",
          "-f=./agents.json",
          "up", "3"}
        args, err := ParseArguments()

        So(err, ShouldEqual, nil)
        So(args, ShouldNotEqual, nil)
        So(args.ConfigPath, ShouldNotBeNil)
        So(*args.ConfigPath, ShouldEqual, "./agents.json")
        So(args.Command, ShouldEqual, UpCommand)
        So(args.CommandParameter, ShouldEqual, 3)
      })
		})
	})
}
