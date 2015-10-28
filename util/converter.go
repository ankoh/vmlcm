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
	one := vmID[0:1]
	two := vmID[2:3]
	three := vmID[4:5]
	four := vmID[6:7]
	fife := vmID[8:9]
	six := vmID[10:11]
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
	return VMIDToMacAddress(splitted[1])
}
