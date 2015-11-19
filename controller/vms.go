package controller

import (
	"fmt"
	"regexp"
	//  "strconv"

	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
)

type virtualMachine struct {
	path     string
	template bool
	clone    bool
	running  bool
}

// getClones calls getVMs to get all the informations about the virtual machines
// It then filters the VMs for clones
func getClones(
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) ([]*virtualMachine, error) {
	vms, err := getVMs(vmrun, config)
	if err != nil {
		return nil, err
	}
	var result []*virtualMachine
	for _, vm := range vms {
		if vm.clone {
			result = append(result, vm)
		}
	}
	return result, nil
}

// getVMs checks the clones directory as well as running vms
// and returns virtualMachine objects
func getVMs(
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) ([]*virtualMachine, error) {

	prefix := config.Prefix
	cloneRegExString := fmt.Sprintf(".*/%s\\-[A-Fa-f0-9]+\\.vmwarevm/%s\\-[A-Fa-f0-9]+\\.vmx$", prefix, prefix)
	cloneRegEx, err := regexp.Compile(cloneRegExString)
	if err != nil {
		return nil, fmt.Errorf("Could not compile Clone RegEx with Prefix %s", config.Prefix)
	}

	clonesDirectoryPath := config.ClonesDirectory
	templatePath := config.TemplatePath
	var vms = make(map[string]*virtualMachine)

	// First get the running vms
	runningVMs, err := getRunningVMPaths(vmrun)
	if err != nil {
		return nil, err
	}

	// Then discover vms in the clones directory
	cloneDirectoryVMs, err := discoverVMs(clonesDirectoryPath)
	if err != nil {
		return nil, err
	}

	// Now iterate over all of them and decide if it is a clone, template or unknown
	for _, runningVM := range runningVMs {
		// Check if the path already exists (only for duplicates in vmrun output here)
		if vm, ok := vms[runningVM]; ok {
			vm.running = true
			continue
		}

		// Otherwise create new vm
		vm := new(virtualMachine)
		vm.path = runningVM
		vm.running = true
		vm.template = vm.path == templatePath
		vm.clone = cloneRegEx.MatchString(runningVM)
		vms[runningVM] = vm
	}

	for _, cloneDirectoryVM := range cloneDirectoryVMs {
		// Check if the path already exists (if the vm is running for instance)
		if _, ok := vms[cloneDirectoryVM]; ok {
			continue
		}

		// Otherwise create a new one
		vm := new(virtualMachine)
		vm.path = cloneDirectoryVM
		vm.running = false
		vm.template = vm.path == templatePath
		vm.clone = cloneRegEx.MatchString(cloneDirectoryVM)
		vms[cloneDirectoryVM] = vm
	}
	
	// Finally add the template VM itself if it is stored outside the clones folder
	if _, ok := vms[templatePath]; !ok {
		vm := new(virtualMachine)
		vm.path = templatePath
		vm.running = false
		vm.template = true
		vm.clone = false
		vms[templatePath] = vm
	}

	// Store VMs in result Array
	result := make([]*virtualMachine, 0, len(vms))
	for _, vm := range vms {
		result = append(result, vm)
	}
	return result, nil
}

// getRunningVMPaths returns the paths of running VMs
func getRunningVMPaths(
	vmrun vmware.VmrunWrapper) ([]string, error) {
	list, err := vmrun.List()
	if err != nil {
		return nil, err
	}
	matches := listVMPaths.FindAllString(list, -1)
	return matches, nil
}

var vmwareExtension = regexp.MustCompile(".*\\.vmwarevm$")
var vmxExtension = regexp.MustCompile(".*\\.vmx$")

// Tries to get the exact vmx paths of all vms that are in the clone directory
func discoverVMs(directoryPath string) ([]string, error) {

	// First read all files within the clones directory
	elements, err := util.ListDirectory(directoryPath)
	if err != nil {
		return nil, err
	}

	// Now filter the vmwarevms
	// Golang functional? Nope nope
	var vmwarevms []string
	for _, element := range elements {
		if vmwareExtension.MatchString(element) {
			vmwarevms = append(vmwarevms, element)
		}
	}

	// Now that we have vmwarepaths, lets search the vmx for each .vmwarevm
	var vmxs []string
	for _, vmwarevm := range vmwarevms {
		elements, err := util.ListDirectory(fmt.Sprintf("%s%s", directoryPath, vmwarevm))
		if err != nil {
			continue
		}

		// Search for the vmx and take the first
		for _, element := range elements {
			if vmxExtension.MatchString(element) {
				vmxs = append(vmxs, fmt.Sprintf("%s%s/%s", directoryPath, vmwarevm, element))
				continue
			}
		}
	}
	return vmxs, nil
}
