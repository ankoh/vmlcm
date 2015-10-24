package util_test

import (
	"os"
	"testing"
  "flag"
  . "github.com/smartystreets/goconvey/convey"

  "gitlab.kohn.io/ankoh/vmlcm/util"
)

func TestArguments(t *testing.T) {
	Convey("Given a clean command line FlagSet", t, func() {
    flag.CommandLine = flag.NewFlagSet("vmlcm", flag.ContinueOnError)

		Convey("When the flag value is separated with whitespace", func() {
      Convey("ParseArguments should return correct values", func() {
        os.Args = []string{
          "vmlcm",
          "-f", "./agents.json",
          "verify"}
        args, err := util.ParseArguments()

        So(err, ShouldEqual, nil)
        So(args, ShouldNotEqual, nil)
        So(args.ConfigPath, ShouldNotBeNil)
        So(*args.ConfigPath, ShouldEqual, "./agents.json")
        So(args.Command, ShouldEqual, util.LCMVerify)
      })
		})

    Convey("When the flag value is separated with an equal sign", func() {
      Convey("ParseArguments should return correct values", func() {
        os.Args = []string{
          "vmlcm",
          "-f=./agents.json",
          "up", "3"}
        args, err := util.ParseArguments()

        So(err, ShouldEqual, nil)
        So(args, ShouldNotEqual, nil)
        So(args.ConfigPath, ShouldNotBeNil)
        So(*args.ConfigPath, ShouldEqual, "./agents.json")
        So(args.Command, ShouldEqual, util.LCMUp)
        So(args.CommandParameter, ShouldEqual, 3)
      })
		})
	})
}
