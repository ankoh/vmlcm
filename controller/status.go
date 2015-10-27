package controller

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
)

// Status returns the VMLCM status
func Status(
	logger *util.Logger,
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration,
	silent bool) error {

	// Fetch vmrun version
	version, err := getVmrunVersion(vmrun)
	if err != nil {
		return err
	}

	// Fetch all vms that can be discovered easily (clone folder && running)
	vms, err := getVMs(vmrun, config)
	if err != nil {
		return err
	}

	// build clones Array
	var clones []*virtualMachine
	var running []*virtualMachine
	var template *virtualMachine
	for _, vm := range vms {
		if vm.clone {
			clones = append(clones, vm)
		}
		if vm.template {
			template = vm
		}
		if vm.running {
			running = append(running, vm)
		}
	}

	// Print report
	if !silent {
		util.PrintASCIIHeader()
		fmt.Println()
		fmt.Printf("%-20s %s%s%s\n", "Vmrun executable", util.ColorCyan, config.Vmrun, util.ColorNone)
		fmt.Printf("%-20s %s\n", "Vmrun version", version.version)
		fmt.Printf("%-20s %s\n", "Vmrun build", version.build)
		fmt.Println()
		fmt.Printf("%-20s %s%s%s\n", "Prefix", util.ColorCyan, config.Prefix, util.ColorNone)
		fmt.Printf("%-20s %s%s%s\n", "Template path", util.ColorCyan, config.TemplatePath, util.ColorNone)
		if template.running {
			fmt.Printf("%-20s %s%s%s\n", "Template status", util.ColorNone, "Online", util.ColorNone)
		} else {
			fmt.Printf("%-20s %s%s%s\n", "Template status", util.ColorNone, "Offline", util.ColorNone)
		}
		fmt.Println("MAC addresses")
		for _, address := range config.Addresses {
			fmt.Printf("%-20s %s%s%s\n", "", util.ColorCyan, address, util.ColorNone)
		}
		fmt.Println()
		fmt.Printf("%-20s %s%s%s\n", "Clones directory", util.ColorCyan, config.ClonesDirectory, util.ColorNone)
		fmt.Printf("%-20s %s%d%s\n", "Linked clones", util.ColorNone, len(clones), util.ColorNone)
		fmt.Println()
		if len(clones) == 0 {
			fmt.Printf("  No clones available\n")
		} else {
			for _, clone := range clones {
				if clone.running {
					name := strings.TrimPrefix(clone.path, config.ClonesDirectory)
					fmt.Printf("  %-65s [%s%s%s]\n", name, util.ColorCyan, "Online", util.ColorNone)
				} else {
					name := strings.TrimPrefix(clone.path, config.ClonesDirectory)
					fmt.Printf("  %-65s [%s%s%s]\n", name, util.ColorLightGray, "Offline", util.ColorNone)
				}
			}
		}
		fmt.Println()
	}
	return nil
}

var helpVmrunVersion = regexp.MustCompile("vmrun version (\\d+\\.\\d+\\.\\d+) build-(\\d+)")
var listVMNumber = regexp.MustCompile("Total running VMs: (\\d+)")
var listVMPaths = regexp.MustCompile("/.*\\.vmx")

type vmrunVersion struct {
	version string
	build   string
}

// getVmrunVersion returns version information of the used vmrun executable
func getVmrunVersion(
	vmrun vmware.VmrunWrapper) (*vmrunVersion, error) {
	vmrunOut := vmrun.GetOutputChannel()
	vmrunErr := vmrun.GetErrorChannel()
	go vmrun.Help()

	var response string
	select {
	case response = <-vmrunOut:
	case err := <-vmrunErr:
		return nil, err
	}

	matches := helpVmrunVersion.FindStringSubmatch(response)
	if len(matches) < 3 {
		return nil, fmt.Errorf("Could not parse vmrun version information")
	}
	result := new(vmrunVersion)
	// index 0 is the whole match itself
	result.version = matches[1]
	result.build = matches[2]
	return result, nil
}

// getRunningVMNumber returns the number of running vms
func getRunningVMNumber(
	vmrun vmware.VmrunWrapper) (int, error) {
	vmrunOut := vmrun.GetOutputChannel()
	vmrunErr := vmrun.GetErrorChannel()
	go vmrun.List()

	var response string
	select {
	case response = <-vmrunOut:
	case err := <-vmrunErr:
		return 0, err
	}

	matches := listVMNumber.FindStringSubmatch(response)
	if len(matches) < 2 {
		return 0, fmt.Errorf("Could not parse vm number information")
	}
	number, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("Could not parse regex match as integer")
	}

	return number, nil
}
