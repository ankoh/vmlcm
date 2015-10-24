package util_test

import (
	"os"
	"testing"
  "flag"
  . "github.com/smartystreets/goconvey/convey"
  "gitlab.kohn.io/ankoh/vmlcm/util"
)

func TestArguments(t *testing.T) {
	Convey("Given a clean command line environment", t, func() {
    flag.CommandLine = flag.NewFlagSet("vmlcm", flag.ContinueOnError)

		Convey("When the flag value is separated with whitespace", func() {
      os.Args = []string{
        "vmlcm",
        "-f", "./agents.json",
        "verify"}

      Convey("ParseArguments must return correct values", func() {
        args, err := util.ParseArguments()
        So(err, ShouldBeNil)
        So(args, ShouldNotBeNil)
        So(args.ConfigPath, ShouldNotBeNil)
        So(*args.ConfigPath, ShouldEqual, "./agents.json")
        So(args.Command, ShouldEqual, util.VerifyCommand)
      })
		})

    Convey("When the flag value is separated with an equal sign", func() {
      os.Args = []string{
        "vmlcm",
        "-f=./agents.json",
        "up", "3"}

      Convey("ParseArguments must return correct values", func() {
        args, err := util.ParseArguments()

        So(err, ShouldBeNil)
        So(args, ShouldNotBeNil)
        So(args.ConfigPath, ShouldNotBeNil)
        So(*args.ConfigPath, ShouldEqual, "./agents.json")
        So(args.Command, ShouldEqual, util.UpCommand)
        So(args.CommandParameter, ShouldEqual, 3)
      })
		})

    Convey("When no arguments are available", func() {
      os.Args = []string{ "vmlcm" }

      Convey("ParseArguments must return an error", func() {
        args, err := util.ParseArguments()
        So(err, ShouldNotBeNil)
        So(args, ShouldBeNil)
      })
    })
	})
}
