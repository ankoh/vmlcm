package controller

import (
  "fmt"
	"bytes"
  "strings"

	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
)

// Start tries to start all associated VMs
func Start(
	buffer *bytes.Buffer,
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) error {

	// Get all existing clones
	clones, err := getClones(vmrun, config)
	if err != nil {
		return err
	}

  util.TryWrite2Columns(buffer, 20, "Clones", fmt.Sprint(len(clones)))
  util.TryWriteln(buffer, "")
  for _, clone := range clones {
    if clone.running {
      continue
    }
    err := vmrun.Start(clone.path)
    if err != nil {
      return err
    }
    vmName := tryVMNameExtraction(clone.path)
    util.TryWrite2Columns(buffer, 20, "  Started Clone", vmName)
  }
  util.TryWriteln(buffer, "")

  return nil
}


// Stop tries to start all associated VMs
func Stop(
	buffer *bytes.Buffer,
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) error {

	// Get all existing clones
	clones, err := getClones(vmrun, config)
	if err != nil {
		return err
	}

  util.TryWrite2Columns(buffer, 20, "Clones", fmt.Sprint(len(clones)))
  util.TryWriteln(buffer, "")
  for _, clone := range clones {
    if !clone.running {
      continue
    }
    err := vmrun.Stop(clone.path, true)
    if err != nil {
      return err
    }
    vmName := tryVMNameExtraction(clone.path)
    util.TryWrite2Columns(buffer, 20, "  Stopped Clone", vmName)
  }
  util.TryWriteln(buffer, "")

  return nil
}

// Given a path extractVMName tries to extract the VM name
func tryVMNameExtraction(path string) string {
  vmxMatch := vmxName.FindStringSubmatch(path)
  if len(vmxMatch) != 2 {
    return path
  }
  vmName := strings.TrimSuffix(vmxMatch[1], ".vmx")
  return vmName
}
