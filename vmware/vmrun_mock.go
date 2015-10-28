package vmware

import (
	"io/ioutil"
)

// MockVmrun provides access to many of the VMware Fusion vmrun API through the command line
type MockVmrun struct {
}

// NewMockVmrun returns a new CLIVmrun object
func NewMockVmrun() *MockVmrun {
	vmrun := new(MockVmrun)
	return vmrun
}

// Help runs the vmrun command without parameters resulting in
// help message
func (vmrun *MockVmrun) Help() (string, error) {
	buffer, err := ioutil.ReadFile("../samples/vmrun/help.txt")
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

// List lists all running vms
func (vmrun *MockVmrun) List() (string, error) {
	buffer, err := ioutil.ReadFile("../samples/vmrun/list.txt")
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

// Start starts a VM
func (vmrun *MockVmrun) Start(vmx string) error {
	return nil
}

// Stop stops a VM
func (vmrun *MockVmrun) Stop(vmx string, force bool) error {
	return nil
}

// Delete starts a VM
func (vmrun *MockVmrun) Delete(vmx string) error {
	return nil
}

// ListSnapshots lists the snapshots of a VM
func (vmrun *MockVmrun) ListSnapshots(vmx string) (string, error) {
	buffer, err := ioutil.ReadFile("../samples/vmrun/listSnapshots.txt")
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

// Snapshot creates a snapshot of a VM
func (vmrun *MockVmrun) Snapshot(vmx string, name string) error {
	return nil
}

// CloneLinked clones a linked VM
func (vmrun *MockVmrun) CloneLinked(template string, cloneDir string, cloneName string, snapshot string) error {
	return nil
}
