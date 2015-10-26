package vmware

import "fmt"
import "sync"

// VmrunWrapper provides access to the Vmrun binary
type VmrunWrapper interface {
	GetOutputChannel() chan string
	GetErrorChannel() chan error
	Close()

  Help()
	List()
	Start(vmx string)
	Stop(vmx string, hard bool)
	Reset(vmx string, hard bool)
	Suspend(vmx string, hard bool)
	Pause(vmx string)
	Unpause(vmx string)
	ListSnapshots(vmx string)
	Snapshot(vmx string, name string)
	DeleteSnapshot(vmx string)
	RevertToSnapshot(vmx string, name string)
	Delete(vmx string)
	CloneLinked(template string, destination string, snapshot string, name string)
}

// CLIVmrun provides access to many of the VMware Fusion vmrun API through the command line
type CLIVmrun struct {
	vmrunPath  string
	vmrunMutex sync.Mutex

	outputChannel chan string
	errorChannel  chan error
}

// NewCLIVmrun returns a new CLIVmrun object
func NewCLIVmrun() *CLIVmrun {
	vmrun := new(CLIVmrun)
	vmrun.outputChannel = make(chan string)
	vmrun.errorChannel = make(chan error)
	return vmrun
}

// Close closes the CLIVmrun channels
func (vmrun *CLIVmrun) Close() {
	close(vmrun.outputChannel)
	close(vmrun.errorChannel)
}

// GetOutputChannel returns the outputchannel of the CLIVmrun wrapper
func (vmrun *CLIVmrun) GetOutputChannel() chan string {
	return vmrun.outputChannel
}

// GetErrorChannel returns the errorchannel of the CLIVmrun wrapper
func (vmrun *CLIVmrun) GetErrorChannel() chan error {
	return vmrun.errorChannel
}

// Help runs the vmrun command without parameters resulting in
// help message
func (vmrun *CLIVmrun) Help() {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion")
}

// List lists all running vms
func (vmrun *CLIVmrun) List() {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"list")
}

// Start starts a VM
func (vmrun *CLIVmrun) Start(vmx string) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"start", vmx, "nogui")
}

// Stop stops a VM
func (vmrun *CLIVmrun) Stop(vmx string, hard bool) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	var force string
	if hard {
		force = "hard"
	} else {
		force = "soft"
	}

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"stop", vmx, force)
}

// Reset resets a VM
func (vmrun *CLIVmrun) Reset(vmx string, hard bool) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	var force string
	if hard {
		force = "hard"
	} else {
		force = "soft"
	}

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"reset", vmx, force)
}

// Suspend suspends a VM
func (vmrun *CLIVmrun) Suspend(vmx string, hard bool) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	var force string
	if hard {
		force = "hard"
	} else {
		force = "soft"
	}

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"suspend", vmx, force)
}

// Pause pauses a VM
func (vmrun *CLIVmrun) Pause(vmx string) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"pause", vmx)
}

// Unpause unpauses a VM
func (vmrun *CLIVmrun) Unpause(vmx string) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"unpause", vmx)
}

// ListSnapshots lists the snapshots of a VM
func (vmrun *CLIVmrun) ListSnapshots(vmx string) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"listSnapshots", vmx)
}

// Snapshot creates a snapshot of a VM
func (vmrun *CLIVmrun) Snapshot(vmx string, name string) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"snapshot", vmx, name)
}

// DeleteSnapshot deletes a snapshot of a VM
func (vmrun *CLIVmrun) DeleteSnapshot(vmx string) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"deleteSnapshot", vmx)
}

// RevertToSnapshot reverts a VM to a snapshot
func (vmrun *CLIVmrun) RevertToSnapshot(vmx string, name string) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"revertToSnapshot", vmx, name)
}

// Delete deletes a VM
func (vmrun *CLIVmrun) Delete(vmx string) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"deleteVM", vmx)
}

// CloneLinked clones a linked VM
func (vmrun *CLIVmrun) CloneLinked(template string, destination string, snapshot string, name string) {
	vmrun.vmrunMutex.Lock()
	defer vmrun.vmrunMutex.Unlock()

	executeCommand(
		vmrun.outputChannel,
		vmrun.errorChannel,
		vmrun.vmrunPath,
		"-T", "fusion",
		"clone", template, destination,
		"linked",
		fmt.Sprintf("-snapshot=%s", snapshot),
		fmt.Sprintf("-cloneName=%s", name))
}
