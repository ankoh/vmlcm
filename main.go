package main

import(
	"fmt"
	"github.com/ankoh/vmlcm/util"
)

func main() {
	// First read arguments
	args, argError := util.ParseArguments()
	if argError != nil {
		fmt.Println(argError.Error())
		return
	}

	// Then read configuration
	_, configError := util.ParseConfiguration(*args.ConfigPath)
	if configError != nil {
		fmt.Println(configError.Error())
		return
	}


}
