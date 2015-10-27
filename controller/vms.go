package controller

import (
  "regexp"
  "fmt"
//  "strconv"

  "github.com/ankoh/vmlcm/util"
  "github.com/ankoh/vmlcm/vmware"
)

type virtualMachine struct {
  name string
  path string
  template bool
  clone bool
  running bool
}

// getRunningVMPaths returns the paths of running VMs
func getRunningVMPaths(
  vmrun vmware.VmrunWrapper) ([]string, error) {
  vmrunOut := vmrun.GetOutputChannel()
  vmrunErr := vmrun.GetErrorChannel()
  go vmrun.List()

  var response string
  select {
    case response = <- vmrunOut:
    case err := <- vmrunErr:
      return nil, err
  }

  matches := listVMPaths.FindAllString(response, -1)
  return matches, nil
}

var vmwareExtension = regexp.MustCompile(".*\\.vmwarevm$")
var vmxExtension = regexp.MustCompile(".*\\.vmx$")

// Tries to get the exact vmx paths of all vms that are in the clone directory
func getDirectoryVMs(
  basePath string) ([]string, error) {

  // First read all files within the clones directory
  elements, err := util.ListDirectory(basePath)
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
    elements, err := util.ListDirectory(fmt.Sprintf("%s%s", basePath, vmwarevm))
    if err != nil {
      continue
    }

    // Search for the vmx and take the first
    for _, element := range elements {
      if vmxExtension.MatchString(element) {
        vmxs = append(vmxs, fmt.Sprintf("%s%s/%s", basePath, vmwarevm, element))
        continue
      }
    }
  }
  return vmxs, nil
}
