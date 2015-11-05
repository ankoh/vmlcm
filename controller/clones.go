package controller

import (
	"fmt"
	"math"
	"strings"
	"regexp"
	"sort"
	"bytes"

	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
)

var vmxName = regexp.MustCompile("([A-Za-z0-9-]+\\.vmx)$")

// Use ensures that <<use>> number of clones exist
// Attention buffer may be nil
func Use(
	buffer *bytes.Buffer,
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration,
	use int) error {
	// Check use parameter
	if use < 0 {
		return fmt.Errorf("Parameter must be >= zero")
	}

	// First check mac addresses
	if len(config.Addresses) < use {
		return fmt.Errorf("Tried to use %d clones with only %d mac addresses",
			use, len(config.Addresses))
	}

	// Then check the number of clones
	// Get all existing clones
	clones, err := getClones(vmrun, config)
	if err != nil {
		return err
	}

	util.TryWrite2Columns(buffer, 20, "Old clones", fmt.Sprint(len(clones)))
	util.TryWrite2Columns(buffer, 20, "New clones", fmt.Sprint(use))

	// Check if number of existing clones equals the parameter (== noop)
	if len(clones) == use {
		util.TryWrite(buffer, "\nNothing to do...\t:'(\n\n")
		return nil
	}

	// Check if clones need to be created
	if len(clones) < use {
		createdClones, err := cloneUpTo(vmrun, config, clones, use)
		if err != nil {
			return err
		}
		util.TryWriteln(buffer, "")
		for _, createdClone := range createdClones {
			util.TryWrite2Columns(buffer, 20, "  Create clone", createdClone)
		}
		util.TryWriteln(buffer, "")
	}

	// Check if clones need to be deleted
	if len(clones) > use {
		deletedClones, err := deleteUpTo(vmrun, config, clones, use)
		if err != nil {
			return err
		}
		util.TryWrite2Columns(buffer, 20, "Deleted clones", fmt.Sprint(len(deletedClones)))
		util.TryWriteln(buffer, "")
		for _, deletedClone := range deletedClones {
			util.TryWrite2Columns(buffer, 20, "  Deleted clone", deletedClone)
		}
		util.TryWriteln(buffer, "")
	}

	return nil
}

// Deletes up to <<use>> clones
func deleteUpTo(
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration,
	clones []*virtualMachine,
	use int) ([]string, error) {

	// First order the clones by status
	clonesOrdered := orderClonesByRunning(clones)
	toDelete := int(math.Min(
		float64(len(clonesOrdered)),
		math.Abs(float64(len(clones)-use))))

	var deletedClones []string

	// Then shutdown target clones
	for i := 0; i < toDelete; i++ {
		clone := clones[i]

		// First try to stop the clone if it is running
		if clone.running {
			// First try a soft stop
			err := vmrun.Stop(clone.path, false)
			if err != nil {
				// then try a hard stop
				err = vmrun.Stop(clone.path, true)
				if err != nil {
					return nil, err
				}
			}
		}

		// Run Vmrun delete
		err := vmrun.Delete(clone.path)
		if err != nil {
			return nil, err
		}

		// Try to find the vmName
		vmxMatch := vmxName.FindStringSubmatch(clone.path)
		if len(vmxMatch) != 2 {
			deletedClones = append(deletedClones, clone.path)
		} else {
			vmName := strings.TrimSuffix(vmxMatch[1], ".vmx")
			deletedClones = append(deletedClones, vmName)
		}
	}

	return deletedClones, nil
}

// orderClonesByRunning orders the clones by their running state
// Offline virtual machines will come first, then the online ones
func orderClonesByRunning(clones []*virtualMachine) []*virtualMachine {
	var result []*virtualMachine
	// First append offline machines
	for _, clone := range clones {
		if !clone.running {
			result = append(result, clone)
		}
	}
	// Then append online machines
	for _, clone := range clones {
		if clone.running {
			result = append(result, clone)
		}
	}
	return result
}

// getAvailableMacAddresses checks which of the mac addresses is free
func getAvailableMacAddresses(
	clones []*virtualMachine,
	config *util.LCMConfiguration) []string {
	// Create map with addresses
	availableAddresses := make(map[string]bool)
	for _, address := range config.Addresses {
		availableAddresses[address] = true
	}

	// Loop through clones and delete addresses that are used
	for _, clone := range clones {
		matches := vmxName.FindStringSubmatch(clone.path)
		if len(matches) != 2 {
			continue
		}
		macAddress, err := util.VMNameToMacAddress(matches[1])
		if err != nil {
			continue
		}
		delete(availableAddresses, macAddress)
	}
	// Create slice with remaining addresses
	var remainingAddresses []string
	for address := range availableAddresses {
		remainingAddresses = append(remainingAddresses, address)
	}
	sort.Strings(remainingAddresses)
	return remainingAddresses
}

// cloneUpTo clones the template up to <<use>> times
func cloneUpTo(
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration,
	clones []*virtualMachine,
	use int) ([]string, error) {
	// Assert use greater zero
	if use <= 0 {
		return []string {}, nil
	}

	// First get available mac addresses
	availableAddresses := getAvailableMacAddresses(clones, config)

	// calculate how many clones need to be created
	diff := math.Max(0, float64(use - len(clones)))
	toCreate := int(math.Min(float64(len(availableAddresses)), diff))
	if toCreate <= 0 {
		return []string {}, nil
	}

	// Get snapshot for the clones
	snapshot, err := prepareTemplateSnapshot(vmrun, config)
	if err != nil {
		return nil, fmt.Errorf("Failed to prepare snapshot\n%s", err.Error())
	}
	var created []string

	// Now fire the clone command for each of the addresses
	for i := 0; i < toCreate; i++ {
		address := availableAddresses[i]
		vmID := util.MacAddressToVMId(address)
		vmName := fmt.Sprintf("%s-%s", config.Prefix, vmID)

		err := vmrun.CloneLinked(
			config.TemplatePath,
			config.ClonesDirectory,
			vmName, snapshot)
		if err != nil {
			return nil, fmt.Errorf("Failed to clone vm %s\n%s", vmName, err.Error())
		}
		created = append(created, vmName)
	}
	return created, nil
}
