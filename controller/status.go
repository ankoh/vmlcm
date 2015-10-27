package controller

import (
  "regexp"
  "fmt"

  "github.com/ankoh/vmlcm/vmware"
)

var helpVmrunVersion = regexp.MustCompile("vmrun version (\\d+\\.\\d+\\.\\d+) build-(\\d+)")
var listVMNumber = regexp.MustCompile("Total running VMs: (\\d+)")
var listRunningVMs = regexp.MustCompile("/.*\\.vmx")

type vmrunVersion struct {
  version string
  build string
}

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
