package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"bytes"

	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
)

// prepareTemplateSnapshot returns the latest snapshot that was made with the prefix
// if none is available, it creates a new one
func prepareTemplateSnapshot(
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) (string, error) {
	snapshots, err := getTemplateSnapshots(vmrun, config)
	if err != nil {
		return "", err
	}

	// Check for the latest snapshot with the given prefix
	foundSnapshot := ""
	foundSnapshotTimestamp := -1
	for _, snapshot := range snapshots {
		parts := strings.Split(snapshot, "-")
		if len(parts) != 2 {
			continue
		}
		if parts[0] != config.Prefix {
			continue
		}
		newTimestamp, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		if newTimestamp > foundSnapshotTimestamp {
			foundSnapshot = parts[0]
			foundSnapshotTimestamp = newTimestamp
		}
	}
	// Check if a snapshot has been found
	if foundSnapshotTimestamp > 0 {
		return fmt.Sprintf("%s-%d", foundSnapshot, foundSnapshotTimestamp), nil
	}

	// Otherwise create a new snapshot
	return createTemplateSnapshot(vmrun, config)
}

// getRunningVMPaths returns the paths of running VMs
func getTemplateSnapshots(
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) ([]string, error) {
	list, err := vmrun.ListSnapshots(config.TemplatePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(list, "\n")

	// Check if at least one line is there
	if len(lines) < 1 {
		return nil, fmt.Errorf("Failed to parse the listSnapshots command")
	}
	// Then remove the first line
	lines = lines[1:]
	var result []string

	// Now remove empty lines
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		result = append(result, line)
	}
	return result, nil
}

// createTemplateSnapshot creates a prefixed snapshot of the template
func createTemplateSnapshot(
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) (string, error) {
	timestamp := int(time.Now().Unix())

	// Create snapshot name
	snapshotName := fmt.Sprintf("%s-%d", config.Prefix, timestamp)
	err := vmrun.Snapshot(config.TemplatePath, snapshotName)

	if err != nil {
		return "", err
	}
	return snapshotName, nil
}

// Snapshot creates a new template snapshot
func Snapshot(
	buffer *bytes.Buffer,
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) error {
	// Fetch all vms that can be discovered easily (clone folder && running)
	vms, err := getVMs(vmrun, config)
	if err != nil {
		return err
	}

	// Find template
	var template *virtualMachine
	for _, vm := range vms {
		if vm.template {
			template = vm
		}
	}

	// Check if template has been found
	if template == nil {
		return fmt.Errorf("Could not find template %s", template)
	}

	// I wont hard shutdown the template for a snapshot as the template is more
	// vulnerable to hard shutdowns
	if template.running {
		util.TryWriteVerification(buffer, "Template Offline", false)

		var warning bytes.Buffer
		buffer.WriteString("Aborting: Please shutdown the template first.\n")
		buffer.WriteString("That needs to be done manually as the template is ")
		buffer.WriteString("probably linked in the Virtual Machine Library ")
		buffer.WriteString("and should not be force stopped.\n")
		return fmt.Errorf(warning.String())
	}
	util.TryWriteVerification(buffer, "Template Offline", true)

	// Then create the snapshot
	snapshotName, err := createTemplateSnapshot(vmrun, config)
	if err != nil {
		return fmt.Errorf(
			"Failed to create a template snapshot.\nError:\n%s", err.Error())
	}
	buffer.WriteString("Created Snapshot: ")
	buffer.WriteString(snapshotName)
	buffer.WriteString("\n")
	return nil
}
