package vmware

import (
	"bytes"
	"io/ioutil"
	"fmt"
)

// MockVmrun provides access to many of the VMware Fusion vmrun API through the command line
type MockVmrun struct {
	TemplateVM string
	RunningVMs []string
	CloneFolderVMs []string
	TemplateSnapshots []string

	// Actions
	TriggeredVMStarts []string
	TriggeredVMStops []string
	TriggeredVMDeletes []string

}

// NewMockVmrun returns a new CLIVmrun object
func NewMockVmrun() *MockVmrun {
	vmrun := new(MockVmrun)
	return vmrun
}

// ClearRunningVMs clears the mocked running virtual machines
func (vmrun *MockVmrun) ClearRunningVMs() {
	vmrun.RunningVMs = vmrun.RunningVMs[:0]
}

// ClearCloneFolderVMs clears the mocked clone folder VMs
func (vmrun *MockVmrun) ClearCloneFolderVMs() {
	vmrun.CloneFolderVMs = vmrun.CloneFolderVMs[:0]
}

// ClearTemplateSnapshots clears the mocked template snapshots
func (vmrun *MockVmrun) ClearTemplateSnapshots() {
	vmrun.TemplateSnapshots = vmrun.TemplateSnapshots[:0]
}

// ClearActions clears the action slices (TriggeredVMStarts etc.)
func (vmrun *MockVmrun) ClearActions() {
	vmrun.TriggeredVMStarts = vmrun.TriggeredVMStarts[:0]
	vmrun.TriggeredVMStops = vmrun.TriggeredVMStops[:0]
	vmrun.TriggeredVMDeletes = vmrun.TriggeredVMDeletes[:0]
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
	var buffer bytes.Buffer
	// Write first line
	firstLine := fmt.Sprintf("Total running VMs: %d\n", len(vmrun.RunningVMs))
	buffer.WriteString(firstLine)
	// Now append the vm paths of each running vm
	for _, vmpath := range vmrun.RunningVMs {
		vmLine := fmt.Sprintf("%s\n", vmpath)
		buffer.WriteString(vmLine)
	}
	// Return the built string
	return buffer.String(), nil
}

// Start starts a VM
func (vmrun *MockVmrun) Start(vmx string) error {
	for _, runningVM := range vmrun.RunningVMs {
		if runningVM == vmx {
			// Found a running VM
			vmrun.TriggeredVMStarts = append(vmrun.TriggeredVMStarts, vmx)
			return nil
		}
	}
	for _, knownVM := range vmrun.CloneFolderVMs {
		if knownVM == vmx {
			// Found a clone folder VM
			vmrun.TriggeredVMStarts = append(vmrun.TriggeredVMStarts, vmx)
			return nil
		}
	}
	return fmt.Errorf("Couldn't find the vmx")
}

// Stop stops a VM
func (vmrun *MockVmrun) Stop(vmx string, force bool) error {
	for _, runningVM := range vmrun.RunningVMs {
		if runningVM == vmx {
			// Found a running VM
			vmrun.TriggeredVMStops = append(vmrun.TriggeredVMStops, vmx)
			return nil
		}
	}
	for _, knownVM := range vmrun.CloneFolderVMs {
		if knownVM == vmx {
			// Found a clone folder VM
			vmrun.TriggeredVMStops = append(vmrun.TriggeredVMStops, vmx)
			return nil
		}
	}
	return fmt.Errorf("Couldn't find the vmx")
}

// Delete starts a VM
func (vmrun *MockVmrun) Delete(vmx string) error {
	for _, runningVM := range vmrun.RunningVMs {
		if runningVM == vmx {
			// Found a running VM
			vmrun.TriggeredVMDeletes = append(vmrun.TriggeredVMStops, vmx)
			return nil
		}
	}
	for _, knownVM := range vmrun.CloneFolderVMs {
		if knownVM == vmx {
			// Found a clone folder VM
			vmrun.TriggeredVMDeletes = append(vmrun.TriggeredVMStops, vmx)
			return nil
		}
	}
	return fmt.Errorf("Couldn't find the vmx")
}

// ListSnapshots lists the snapshots of a VM
func (vmrun *MockVmrun) ListSnapshots(vmx string) (string, error) {
	var buffer bytes.Buffer
	// Write first line
	firstLine := fmt.Sprintf("Total snapshots: %d\n", len(vmrun.TemplateSnapshots))
	buffer.WriteString(firstLine)
	// Now append the vm paths of each running vm
	for _, snapshot := range vmrun.TemplateSnapshots {
		snapshotLine := fmt.Sprintf("%s\n", snapshot)
		buffer.WriteString(snapshotLine)
	}
	// Return the built string
	return buffer.String(), nil
}

// Snapshot creates a snapshot of a VM
func (vmrun *MockVmrun) Snapshot(vmx string, name string) error {
	// The mocked snapshot function only creates snapshots of the template
	if vmx != vmrun.TemplateVM {
		return fmt.Errorf("Tried to create a snapshot of an unknown VM!")
	}
	// Check if the snapshot name already exists
	for _, existingSnapshot := range vmrun.TemplateSnapshots {
		if existingSnapshot == name {
			return fmt.Errorf("Tried to create a snapshot with a name collision!")
		}
	}
	vmrun.TemplateSnapshots = append(vmrun.TemplateSnapshots, name)
	return nil
}

// CloneLinked clones a linked VM
func (vmrun *MockVmrun) CloneLinked(template string, cloneDir string, cloneName string, snapshot string) error {
	// Check for path collisions
  clonePath := fmt.Sprintf("%s%s.vmwarevm/%s.vmx", cloneDir, cloneName, cloneName)
	for _, cloneFolderVM := range vmrun.CloneFolderVMs {
		if clonePath == cloneFolderVM {
			return fmt.Errorf("Tried to create a VM that already exists!")
		}
	}
	// Check if template exists
	if template != vmrun.TemplateVM {
		return fmt.Errorf("Template does not exist!")
	}
	// Check if snapshot exists
	found := false
	for _, existingSnapshot := range vmrun.TemplateSnapshots {
		if existingSnapshot == snapshot {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("Snapshot does not exist!")
	}
	// Add the clone and return
	vmrun.CloneFolderVMs = append(vmrun.CloneFolderVMs, clonePath)
	return nil
}
