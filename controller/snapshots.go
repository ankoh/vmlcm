package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
