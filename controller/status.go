package controller

import (
  "regexp"
  "fmt"
  "strconv"

  "github.com/ankoh/vmlcm/vmware"
)

var helpVmrunVersion = regexp.MustCompile("vmrun version (\\d+\\.\\d+\\.\\d+) build-(\\d+)")
var listVMNumber = regexp.MustCompile("Total running VMs: (\\d+)")
var listRunningVMs = regexp.MustCompile("/.*\\.vmx")

type vmrunVersion struct {
  version string
  build string
}

// getVmrunVersion returns version information of the used vmrun executable
func getVmrunVersion(
  vmrun vmware.VmrunWrapper) (*vmrunVersion, error) {
  vmrunOut := vmrun.GetOutputChannel()
  vmrunErr := vmrun.GetErrorChannel()
  go vmrun.Help()

  var response string
  select {
    case response = <- vmrunOut:
    case err := <- vmrunErr:
      return nil, err
  }

  matches := helpVmrunVersion.FindStringSubmatch(response)
  if len(matches) < 3 {
    return nil, fmt.Errorf("Could not parse vmrun version information")
  }
  result := new(vmrunVersion)
  // index 0 is the whole match itself
  result.version = matches[1]
  result.build = matches[2]
  return result, nil
}

// getRunningVMNumber returns the number of running vms
func getRunningVMNumber(
  vmrun vmware.VmrunWrapper) (int, error) {
  vmrunOut := vmrun.GetOutputChannel()
  vmrunErr := vmrun.GetErrorChannel()
  go vmrun.List()

  var response string
  select {
    case response = <- vmrunOut:
    case err := <- vmrunErr:
      return 0, err
  }

  matches := listVMNumber.FindStringSubmatch(response)
  if len(matches) < 2 {
    return 0, fmt.Errorf("Could not parse vm number information")
  }
  number, err := strconv.Atoi(matches[1])
  if err != nil {
    return 0, fmt.Errorf("Could not parse regex match as integer")
  }

  return number, nil
}
