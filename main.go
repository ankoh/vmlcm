package main

import (
	"fmt"

	"github.com/ankoh/vmlcm/controller"
	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
)

func main() {
	// First read arguments
	args, argError := util.ParseArguments()
	if argError != nil {
		fmt.Println(argError.Error())
		return
	}

	// Then read configuration
	config, configError := util.ParseConfiguration(*args.ConfigPath)
	if configError != nil {
		fmt.Println(configError.Error())
		return
	}

	// Create logger
	logger := util.NewLogger()

	// Create vmrun wrapper
	var vmrun vmware.VmrunWrapper
	vmrun = vmware.NewCLIVmrun(config.Vmrun)

	// Switch commands
	switch args.Command {
	case util.VerifyCommand:
		err := controller.Verify(logger, vmrun, config, false)
		if err != nil {
			fmt.Println(err.Error())
		}
	case util.StatusCommand:
		err := controller.Status(logger, vmrun, config, false)
		if err != nil {
			fmt.Println(err.Error())
		}
	case util.UpCommand:
		fmt.Println("Not implemented yet")
	case util.KeepCommand:
		fmt.Println("Not implemented yet")
	case util.ResetCommand:
		fmt.Println("Not implemented yet")
	case util.StartCommand:
		fmt.Println("Not implemented yet")
	case util.StopCommand:
		fmt.Println("Not implemented yet")
	case util.SuspendCommand:
		fmt.Println("Not implemented yet")
	}
}
