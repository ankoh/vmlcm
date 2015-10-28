package util

import (
	"fmt"
	"strings"
)

// MacAddressToVMId converts a Mac Address to a VM ID
func MacAddressToVMId(address string) string {
	splitted := strings.Split(address, ":")
	joined := strings.Join(splitted, "")
	return joined
}

// VMIDToMacAddress converts a VMID into a Mac Address
func VMIDToMacAddress(vmID string) (string, error) {
	if len(vmID) != 12 {
		return "", fmt.Errorf("Invalid length of Virtual Machine ID")
	}
	one := vmID[0:2]
	two := vmID[2:4]
	three := vmID[4:6]
	four := vmID[6:8]
	fife := vmID[8:10]
	six := vmID[10:12]
	all := []string{
		one,
		two,
		three,
		four,
		fife,
		six,
	}
	return strings.Join(all, ":"), nil
}

// VMNameToMacAddress converts a VM name intio a Mac Address
func VMNameToMacAddress(vmName string) (string, error) {
	splitted := strings.Split(vmName, "-")
	if len(splitted) != 2 {
		return "", fmt.Errorf("Invalid format of Virtual Machine length")
	}
	splitted[1] = strings.TrimSuffix(splitted[1], ".vmx")
	return VMIDToMacAddress(splitted[1])
}
