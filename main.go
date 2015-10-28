package main

import (
	"fmt"

	"github.com/ankoh/vmlcm/controller"
	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
	"github.com/briandowns/spinner"
	"time"
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
	logger.Silent = false
	spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	// Create vmrun wrapper
	var vmrun vmware.VmrunWrapper
	vmrun = vmware.NewCLIVmrun(config.Vmrun)

	// Switch commands
	switch args.Command {
	case util.VerifyCommand:
		err := controller.Verify(logger, vmrun, config)
		if err != nil {
			fmt.Println(err.Error())
		}
	case util.StatusCommand:
		spinner.Start()
		if verify(logger, vmrun, config) != nil {
			spinner.Stop()
			return
		}

		err := controller.Status(logger, vmrun, config, spinner)
		if err != nil {
			fmt.Println(err.Error())
		}
	case util.UseCommand:
		fmt.Println("Not implemented yet")
	case util.StartCommand:
		fmt.Println("Not implemented yet")
	case util.StopCommand:
		fmt.Println("Not implemented yet")
	}
}

func verify(
	logger *util.Logger,
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) error {
	logger.Silent = true
	err := controller.Verify(logger, vmrun, config)
	logger.Silent = false
	if err != nil {
		fmt.Println("Failed to verify settings. Please run 'verify' to get more details")
	}
	return err
}
