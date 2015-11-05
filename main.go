package main

import (
	"fmt"

	"github.com/ankoh/vmlcm/controller"
	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
	"github.com/briandowns/spinner"
	"time"
	"bytes"
)

func main() {
	// Create spinner
	spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spinner.Color("cyan")
	spinner.Start()

	// First read arguments
	args, argError := util.ParseArguments()
	if argError != nil {
		spinner.Stop()
		fmt.Println(argError.Error())
		return
	}

	// Then read configuration
	config, configError := util.ParseConfiguration(*args.ConfigPath)
	if configError != nil {
		spinner.Stop()
		fmt.Println(configError.Error())
		return
	}

	// Create vmrun wrapper
	var vmrun vmware.VmrunWrapper
	vmrun = vmware.NewCLIVmrun(config.Vmrun)

	// Create output buffer
	buffer := new(bytes.Buffer)

	// Switch commands
	switch args.Command {
	case util.VerifyCommand:
		err := controller.Verify(buffer, vmrun, config)
		if err != nil {
			spinner.Stop()
			fmt.Println(err.Error())
			return
		}
	case util.StatusCommand:
		if controller.Verify(nil, vmrun, config) != nil {
			spinner.Stop()
			fmt.Println("Failed to verify settings. Please run 'verify' to get more details")
			return
		}
		err := controller.Status(buffer, vmrun, config)
		if err != nil {
			spinner.Stop()
			fmt.Println(err.Error())
			return
		}
	case util.UseCommand:
		if controller.Verify(nil, vmrun, config) != nil {
			spinner.Stop()
			fmt.Println("Failed to verify settings. Please run 'verify' to get more details")
			return
		}
		err := controller.Use(buffer, vmrun, config, args.CommandIntParameter)
		if err != nil {
			spinner.Stop()
			fmt.Print(buffer.String())
			fmt.Println(err.Error())
			return
		}
	case util.StartCommand:
		spinner.Stop()
		fmt.Println("Not implemented yet")
	case util.StopCommand:
		spinner.Stop()
		fmt.Println("Not implemented yet")
	}

	// Stop spinner
	spinner.Stop()

	// Print output buffer
	fmt.Print(buffer.String())
}
