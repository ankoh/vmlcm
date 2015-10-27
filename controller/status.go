package controller

import (
  "regexp"
  "fmt"
  "strconv"

  "github.com/ankoh/vmlcm/util"
  "github.com/ankoh/vmlcm/vmware"
)

// Status returns the VMLCM status
func Status(
  logger *util.Logger,
  vmrun vmware.VmrunWrapper,
  config *util.LCMConfiguration,
  silent bool) error {

  // Fetch vmrun version
  version, err := getVmrunVersion(vmrun)
  if err != nil {
    return err
  }
  // Fetch running vms number
  runningVms, err := getRunningVMNumber(vmrun)
  if err != nil {
    return err
  }

  // Print report
  if !silent {
    printHeader()
    fmt.Println()
    fmt.Printf("%-25s %s%s%s\n", "Vmrun version", util.ColorCyan, version.version, util.ColorNone)
    fmt.Printf("%-25s %s%s%s\n", "Vmrun build", util.ColorCyan, version.build, util.ColorNone)
    fmt.Printf("%-25s %s%d%s\n", "Running VMs", util.ColorCyan, runningVms, util.ColorNone)
    fmt.Println()
  }
  return nil
}

func printHeader() {
  fmt.Println("                  __              ")
  fmt.Println(" _   ______ ___  / /___ ___  _____")
  fmt.Println("| | / / __ `__ \\/ / __ `__ \\/ ___/")
  fmt.Println("| |/ / / / / / / / / / / / / /__  ")
  fmt.Println("|___/_/ /_/ /_/_/_/ /_/ /_/\\___/  ")
  fmt.Println("                                  ")
}


var helpVmrunVersion = regexp.MustCompile("vmrun version (\\d+\\.\\d+\\.\\d+) build-(\\d+)")
var listVMNumber = regexp.MustCompile("Total running VMs: (\\d+)")
var listVMPaths = regexp.MustCompile("/.*\\.vmx")

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
