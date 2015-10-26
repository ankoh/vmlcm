package vmware

import "fmt"

// Command to start a VM
const startCommand string = "%s start %s nogui"
func getStartCommand(vmrun string, vmx string) string {
  return fmt.Sprintf(startCommand, vmrun, vmx)
}

// Commands to stop a VM
const stopSoftCommand string = "%s stop %s soft"
const stopHardCommand string = "%s stop %s hard"
func getStopCommand(vmrun string, vmx string, hard bool) string {
  if hard {
    return fmt.Sprintf(stopHardCommand, vmrun, vmx)
  }
  return fmt.Sprintf(stopSoftCommand, vmrun, vmx)
}

// Commands to reset a VM
const resetSoftCommand string = "%s reset %s soft"
const resetHardCommand string = "%s reset %s hard"
func getResetCommand(vmrun string, vmx string, hard bool) string {
  if hard {
    return fmt.Sprintf(resetHardCommand, vmrun, vmx)
  }
  return fmt.Sprintf(resetSoftCommand, vmrun, vmx)
}

// Commands to suspend a VM
const suspendSoftCommand string = "%s suspend %s soft"
const suspendHardCommand string = "%s suspend %s hard"
func getSuspendCommand(vmrun string, vmx string, hard bool) string {
  if hard {
    return fmt.Sprintf(suspendHardCommand, vmrun, vmx)
  }
  return fmt.Sprintf(suspendSoftCommand, vmrun, vmx)
}

// Command to pause a VM
const pauseCommand string = "%s pause %s"
func getPauseCommand(vmrun string, vmx string) string {
  return fmt.Sprintf(pauseCommand, vmrun, vmx)
}

// Command to unpause a VM
const unpauseCommand string = "%s unpause %s"
func getUnpauseCommand(vmrun string, vmx string) string {
  return fmt.Sprintf(unpauseCommand, vmrun, vmx)
}

// Command to list the VMs snapshot
const listSnapshotsCommand string = "%s listSnapshots %s"
func getListSnapshotsCommand(vmrun string, vmx string) string {
  return fmt.Sprintf(listSnapshotsCommand, vmrun, vmx)
}

// Command to create a VM snapshot
const snapshotCommand string = "%s snapshot %s %s"
func getSnapshotCommand(vmrun string, vmx string, name string) string {
  return fmt.Sprintf(snapshotCommand, vmrun, vmx, name)
}

// Command to delete a snapshot
const deleteSnapshotCommand string = "%s deleteSnapshot %s"
func getDeleteSnapshotCommand(vmrun string, vmx string) string {
  return fmt.Sprintf(deleteSnapshotCommand, vmrun, vmx)
}

// Command to revert to a snapshot
const revertToSnapshotCommand string = "%s revertToSnapshot %s %s"
func getRevertToSnapshotCommand(vmrun string, vmx string, name string) string {
  return fmt.Sprintf(revertToSnapshotCommand, vmrun, vmx, name)
}

// Command to delete a VM
const deleteCommand string = "%s deleteVM %s"
func getDeleteCommand(vmrun string, vmx string) string {
  return fmt.Sprintf(deleteCommand, vmrun, vmx)
}

// Command to create a linked clone of a VM
const cloneLinkedCommand string = "%s clone %s %s linked -snapshot=%s -cloneName=%s"
func getCloneLinkedCommand(vmrun string, template string, destination string, snapshot string, name string) string {
  return fmt.Sprintf(cloneLinkedCommand, vmrun, template, destination, snapshot, name)
}


// Vmrun provides access to many of the VMware Fusion vmrun API
type Vmrun struct {
  vmrunPath string
}

// Start starts a VM
func (*Vmrun) Start(vmx string) error {
  return nil
}

// Stop stops a VM
func (*Vmrun) Stop(vmx string, hard bool) error {
  return nil
}

// Reset resets a VM
func (*Vmrun) Reset(vmx string, hard bool) error {
  return nil
}

// Suspend suspends a VM
func (*Vmrun) Suspend(vmx string, hard bool) error {
  return nil
}

// Pause pauses a VM
func (*Vmrun) Pause(vmx string) error {
  return nil
}

// Unpause unpauses a VM
func (*Vmrun) Unpause(vmx string) error {
  return nil
}

// ListSnapshots lists the snapshots of a VM
func (*Vmrun) ListSnapshots(vmx string) error {
  return nil
}

// Snapshot creates a snapshot of a VM
func (*Vmrun) Snapshot(vmx string, name string) error {
  return nil
}

// DeleteSnapshot deletes a snapshot of a VM
func (*Vmrun) DeleteSnapshot(vmx string, name string) error {
  return nil
}

// RevertToSnapshot reverts a VM to a snapshot
func (*Vmrun) RevertToSnapshot(vmx string, name string) error {
  return nil
}

// Delete deletes a VM
func (*Vmrun) Delete(vmx string) error {
  return nil
}

// CloneLinked clones a linked VM
func (*Vmrun) CloneLinked(template string, destination string, snapshot string, name string) error {
  return nil
}
