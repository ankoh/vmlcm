package vmware

import "fmt"
import "sync"

// VmrunWrapper provides access to the Vmrun binary
type VmrunWrapper interface {
  GetOutputChannel() chan string
  GetErrorChannel() chan error

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
  vmrunPath string
  vmrunMutex sync.Mutex

  outputChannel chan string
  errorChannel chan error
}

// GetOutputChannel returns the outputchannel of the CLIVmrun wrapper
func (vmrun *CLIVmrun) GetOutputChannel() chan string {
  return vmrun.outputChannel
}

// GetErrorChannel returns the errorchannel of the CLIVmrun wrapper
func (vmrun *CLIVmrun) GetErrorChannel() chan error {
  return vmrun.errorChannel
}

// Start starts a VM
func (vmrun *CLIVmrun) Start(vmx string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.outputChannel,
    vmrun.errorChannel,
    vmrun.vmrunPath,
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
    "clone", template, destination,
    "linked",
    fmt.Sprintf("-snapshot=%s", snapshot),
    fmt.Sprintf("-cloneName=%s", name))
}
