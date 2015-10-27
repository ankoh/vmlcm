package util_test

import (
	"flag"
	"github.com/ankoh/vmlcm/util"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestArguments(t *testing.T) {
	Convey("Given a clean command line environment", t, func() {
		flag.CommandLine = flag.NewFlagSet("vmlcm", flag.ContinueOnError)

		Convey("whitespace flags should be parsed correctly", func() {
			os.Args = []string{
				"vmlcm",
				"-f", "./agents.json",
				"verify"}

			args, err := util.ParseArguments()
			So(err, ShouldBeNil)
			So(args, ShouldNotBeNil)
			So(args.ConfigPath, ShouldNotBeNil)
			So(*args.ConfigPath, ShouldEqual, "./agents.json")
			So(args.Command, ShouldEqual, util.VerifyCommand)
		})

		Convey("non-whitespace flags should be parsed correctly", func() {
			os.Args = []string{
				"vmlcm",
				"-f=./agents.json",
				"up", "3"}

			args, err := util.ParseArguments()
			So(err, ShouldBeNil)
			So(args, ShouldNotBeNil)
			So(args.ConfigPath, ShouldNotBeNil)
			So(*args.ConfigPath, ShouldEqual, "./agents.json")
			So(args.Command, ShouldEqual, util.UpCommand)
			So(args.CommandIntParameter, ShouldEqual, 3)
		})

		Convey("no arguments should result in an error", func() {
			os.Args = []string{"vmlcm"}

			args, err := util.ParseArguments()
			So(err, ShouldNotBeNil)
			So(args, ShouldBeNil)
		})
	})
}
