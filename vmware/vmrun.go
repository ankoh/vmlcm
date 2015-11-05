package vmware

import (
	"fmt"
	"os"
	"os/exec"
)

// VmrunWrapper provides access to the Vmrun binary
type VmrunWrapper interface {
	Help() (string, error)
	List() (string, error)
	Start(vmx string) error
	Stop(vmx string, hard bool) error
	ListSnapshots(vmx string) (string, error)
	Snapshot(vmx string, name string) error
	Delete(vmx string) error
	CloneLinked(template string,
		cloneDir string, cloneName string, snapshot string) error
}

// CLIVmrun provides access to many of the VMware Fusion vmrun API through the command line
type CLIVmrun struct {
	vmrunPath string
}

// NewCLIVmrun returns a new CLIVmrun object
func NewCLIVmrun(vmrunPath string) *CLIVmrun {
	vmrun := new(CLIVmrun)
	vmrun.vmrunPath = vmrunPath
	return vmrun
}

// Help runs the vmrun command without parameters resulting in
// help message
func (vmrun *CLIVmrun) Help() (string, error) {
	out, _ := exec.Command(vmrun.vmrunPath, "-T", "fusion").Output()
	return fmt.Sprintf("%s", out), nil
}

// List lists all running vms
func (vmrun *CLIVmrun) List() (string, error) {
	out, err := exec.Command(vmrun.vmrunPath, "-T", "fusion", "list").Output()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", out), nil
}

// Start starts a VM
func (vmrun *CLIVmrun) Start(vmx string) error {
	_, err := exec.Command(vmrun.vmrunPath, "-T", "fusion", "start", vmx, "nogui").Output()
	return err
}

// Stop stops a VM
func (vmrun *CLIVmrun) Stop(vmx string, hard bool) error {
	var force string
	if hard {
		force = "hard"
	} else {
		force = "soft"
	}
	_, err := exec.Command(vmrun.vmrunPath, "-T", "fusion", "stop", vmx, force).Output()
	return err
}

// ListSnapshots lists the snapshots of a VM
func (vmrun *CLIVmrun) ListSnapshots(vmx string) (string, error) {
	out, err := exec.Command(vmrun.vmrunPath, "-T", "fusion", "listSnapshots", vmx).Output()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", out), nil
}

// Snapshot creates a snapshot of a VM
func (vmrun *CLIVmrun) Snapshot(vmx string, name string) error {
	_, err := exec.Command(vmrun.vmrunPath, "-T", "fusion", "snapshot", vmx, name).Output()
	return err
}

// Delete deletes a VM
func (vmrun *CLIVmrun) Delete(vmx string) error {
	_, err := exec.Command(vmrun.vmrunPath, "-T", "fusion", "deleteVM", vmx).Output()
	return err
}

// CloneLinked clones a linked VM
func (vmrun *CLIVmrun) CloneLinked(
	template string,
	cloneDir string,
	cloneName string,
	snapshot string) error {
	vmwarevmPath := fmt.Sprintf("%s%s.vmwarevm", cloneDir, cloneName)
	vmxPath := fmt.Sprintf("%s/%s.vmx", vmwarevmPath, cloneName)
	err := os.Mkdir(vmwarevmPath, 0755)
	if err != nil {
		return err
	}
	_, err = exec.Command(vmrun.vmrunPath,
		"-T", "fusion",
		"clone", template, vmxPath,
		"linked",
		fmt.Sprintf("-snapshot=%s", snapshot),
		fmt.Sprintf("-cloneName=%s", cloneName)).Output()
	return err
}
