package vmware

import "fmt"
import "sync"

// Vmrun provides access to many of the VMware Fusion vmrun API
type Vmrun struct {
  vmrunPath string
  vmrunMutex sync.Mutex

  OutputChannel chan string
  ErrorChannel chan error
}

// Start starts a VM
func (vmrun *Vmrun) Start(vmx string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "start", vmx, "nogui")
}

// Stop stops a VM
func (vmrun *Vmrun) Stop(vmx string, hard bool) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  var force string
  if hard {
    force = "hard"
  } else {
    force = "soft"
  }

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "stop", vmx, force)
}

// Reset resets a VM
func (vmrun *Vmrun) Reset(vmx string, hard bool) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  var force string
  if hard {
    force = "hard"
  } else {
    force = "soft"
  }

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "reset", vmx, force)
}

// Suspend suspends a VM
func (vmrun *Vmrun) Suspend(vmx string, hard bool) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  var force string
  if hard {
    force = "hard"
  } else {
    force = "soft"
  }

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "suspend", vmx, force)
}

// Pause pauses a VM
func (vmrun *Vmrun) Pause(vmx string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "pause", vmx)
}

// Unpause unpauses a VM
func (vmrun *Vmrun) Unpause(vmx string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "unpause", vmx)
}

// ListSnapshots lists the snapshots of a VM
func (vmrun *Vmrun) ListSnapshots(vmx string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "listSnapshots", vmx)
}

// Snapshot creates a snapshot of a VM
func (vmrun *Vmrun) Snapshot(vmx string, name string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "snapshot", vmx, name)
}

// DeleteSnapshot deletes a snapshot of a VM
func (vmrun *Vmrun) DeleteSnapshot(vmx string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "deleteSnapshot", vmx)
}

// RevertToSnapshot reverts a VM to a snapshot
func (vmrun *Vmrun) RevertToSnapshot(vmx string, name string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "revertToSnapshot", vmx, name)
}

// Delete deletes a VM
func (vmrun *Vmrun) Delete(vmx string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "deleteVM", vmx)
}

// CloneLinked clones a linked VM
func (vmrun *Vmrun) CloneLinked(template string, destination string, snapshot string, name string) {
  vmrun.vmrunMutex.Lock()
  defer vmrun.vmrunMutex.Unlock()

  executeCommand(
    vmrun.OutputChannel,
    vmrun.ErrorChannel,
    vmrun.vmrunPath,
    "clone", template, destination,
    "linked",
    fmt.Sprintf("-snapshot=%s", snapshot),
    fmt.Sprintf("-cloneName=%s", name))
}
