package vmware

import (
	"io/ioutil"
)

// MockVmrun provides access to many of the VMware Fusion vmrun API through the command line
type MockVmrun struct {
	outputChannel chan string
	errorChannel  chan error
}

// GetOutputChannel returns the outputchannel of the MockVmrun wrapper
func (vmrun *MockVmrun) GetOutputChannel() chan string {
	return vmrun.outputChannel
}

// GetErrorChannel returns the errorchannel of the MockVmrun wrapper
func (vmrun *MockVmrun) GetErrorChannel() chan error {
	return vmrun.errorChannel
}

// NewMockVmrun returns a new CLIVmrun object
func NewMockVmrun() *MockVmrun {
	vmrun := new(MockVmrun)
	vmrun.outputChannel = make(chan string)
	vmrun.errorChannel = make(chan error)
	return vmrun
}

// Close closes the MockVmrun channels
func (vmrun *MockVmrun) Close() {
	close(vmrun.outputChannel)
	close(vmrun.errorChannel)
}

// Help runs the vmrun command without parameters resulting in
// help message
func (vmrun *MockVmrun) Help() {
	buffer, err := ioutil.ReadFile("../samples/vmrun/help.txt")
	if err != nil {
		vmrun.errorChannel <- err
		return
	}
	vmrun.outputChannel <- string(buffer)
}

// List lists all running vms
func (vmrun *MockVmrun) List() {
	buffer, err := ioutil.ReadFile("../samples/vmrun/list.txt")
	if err != nil {
		vmrun.errorChannel <- err
		return
	}
	vmrun.outputChannel <- string(buffer)
}

// Start starts a VM
func (vmrun *MockVmrun) Start(vmx string) {

}

// Stop stops a VM
func (vmrun *MockVmrun) Stop(vmx string, hard bool) {

}

// Reset resets a VM
func (vmrun *MockVmrun) Reset(vmx string, hard bool) {

}

// Suspend suspends a VM
func (vmrun *MockVmrun) Suspend(vmx string, hard bool) {

}

// Pause pauses a VM
func (vmrun *MockVmrun) Pause(vmx string) {

}

// Unpause unpauses a VM
func (vmrun *MockVmrun) Unpause(vmx string) {

}

// ListSnapshots lists the snapshots of a VM
func (vmrun *MockVmrun) ListSnapshots(vmx string) {
	buffer, err := ioutil.ReadFile("../samples/vmrun/listSnapshots.txt")
	if err != nil {
		vmrun.errorChannel <- err
		return
	}
	vmrun.outputChannel <- string(buffer)
}

// Snapshot creates a snapshot of a VM
func (vmrun *MockVmrun) Snapshot(vmx string, name string) {
	vmrun.outputChannel <- ""
}

// DeleteSnapshot deletes a snapshot of a VM
func (vmrun *MockVmrun) DeleteSnapshot(vmx string) {

}

// RevertToSnapshot reverts a VM to a snapshot
func (vmrun *MockVmrun) RevertToSnapshot(vmx string, name string) {

}

// Delete deletes a VM
func (vmrun *MockVmrun) Delete(vmx string) {

}

// CloneLinked clones a linked VM
func (vmrun *MockVmrun) CloneLinked(template string, cloneDir string, cloneName string, snapshot string){
}
