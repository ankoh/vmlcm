package controller

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"bytes"
	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
)

// Status returns the VMLCM status
func Status(
	buffer *bytes.Buffer,
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) error {
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

	// get template snapshots
	snapshots, err := getTemplateSnapshots(vmrun, config)
	if err != nil {
		return err
	}

	// Print report
	buffer.WriteString(util.GenerateASCIIHeader())
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("%-20s %s%s%s\n", "Vmrun executable", util.ColorCyan, config.Vmrun, util.ColorNone))
	buffer.WriteString(fmt.Sprintf("%-20s %s\n", "Vmrun version", version.version))
	buffer.WriteString(fmt.Sprintf("%-20s %s\n\n", "Vmrun build", version.build))
	buffer.WriteString(fmt.Sprintf("%-20s %s%s%s\n", "Prefix", util.ColorCyan, config.Prefix, util.ColorNone))
	buffer.WriteString(fmt.Sprintf("%-20s %s%s%s\n", "Template path", util.ColorCyan, config.TemplatePath, util.ColorNone))
	if template.running {
		buffer.WriteString(fmt.Sprintf("%-20s %s%s%s\n", "Template status", util.ColorNone, "Online", util.ColorNone))
	} else {
		buffer.WriteString(fmt.Sprintf("%-20s %s%s%s\n", "Template status", util.ColorNone, "Offline", util.ColorNone))
	}
	buffer.WriteString("Template snapshots\n")
	for _, snapshot := range snapshots {
		buffer.WriteString(fmt.Sprintf("%-20s %s%s%s\n", "", util.ColorCyan, snapshot, util.ColorNone))
	}
	if len(snapshots) == 0 {
		buffer.WriteString(fmt.Sprintf("%-20s %s\n", "", "No snapshots existing\n"))
	}
	buffer.WriteString("\nMAC addresses\n")
	for _, address := range config.Addresses {
		buffer.WriteString(fmt.Sprintf("%-20s %s%s%s\n", "", util.ColorCyan, address, util.ColorNone))
	}
	if len(config.Addresses) == 0 {
		buffer.WriteString(fmt.Sprintf("%-20s %s\n", "", "No addresses existing\n"))
	}
	buffer.WriteString(fmt.Sprintf("\n%-20s %s%s%s\n", "Clones directory", util.ColorCyan, config.ClonesDirectory, util.ColorNone))
	buffer.WriteString(fmt.Sprintf("%-20s %s%d%s\n\n", "Linked clones", util.ColorNone, len(clones), util.ColorNone))
	if len(clones) == 0 {
		buffer.WriteString("  No clones exsting for given prefix\n")
	} else {
		for _, clone := range clones {
			if clone.running {
				name := strings.TrimPrefix(clone.path, config.ClonesDirectory)
				buffer.WriteString(fmt.Sprintf("  %-65s [%s%s%s]\n", name, util.ColorCyan, "Online", util.ColorNone))
			} else {
				name := strings.TrimPrefix(clone.path, config.ClonesDirectory)
				buffer.WriteString(fmt.Sprintf("  %-65s [%s%s%s]\n", name, util.ColorLightGray, "Offline", util.ColorNone))
			}
		}
	}
	buffer.WriteString("\n")

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
	help, err := vmrun.Help()
	if err != nil {
		return nil, err
	}

	matches := helpVmrunVersion.FindStringSubmatch(help)
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
	list, err := vmrun.List()
	if err != nil {
		return 0, err
	}

	matches := listVMNumber.FindStringSubmatch(list)
	if len(matches) < 2 {
		return 0, fmt.Errorf("Could not parse vm number information")
	}
	number, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("Could not parse regex match as integer")
	}

	return number, nil
}
