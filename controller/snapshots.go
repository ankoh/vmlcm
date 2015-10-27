package controller

import(
  "strings"
  "fmt"
  "strconv"
  "time"

  "github.com/ankoh/vmlcm/util"
  "github.com/ankoh/vmlcm/vmware"
)

// getRunningVMPaths returns the paths of running VMs
func getTemplateSnapshots(
	vmrun vmware.VmrunWrapper,
  config *util.LCMConfiguration) ([]string, error) {
	vmrunOut := vmrun.GetOutputChannel()
	vmrunErr := vmrun.GetErrorChannel()
	go vmrun.ListSnapshots(config.TemplatePath)

	var response string
	select {
	case response = <-vmrunOut:
	case err := <-vmrunErr:
		return nil, err
	}
  lines := strings.Split(response, "\n")

  // Check if at least one line is there
  if len(lines) < 1 {
    return nil, fmt.Errorf("Failed to parse the listSnapshots command")
  }
  // Then remove the first line
  lines = lines[1:]
  var result []string

  // Now remove empty lines
  for _, line := range lines {
    if len(line) == 0 {
      continue
    }
    result = append(result, line)
  }
	return result, nil
}

// createTemplateSnapshot creates a prefixed snapshot of the template
func createTemplateSnapshot(
  vmrun vmware.VmrunWrapper,
  config *util.LCMConfiguration,
  withSpinner bool) error {
  timestamp := int(time.Now().Unix())

  // Create snapshot name
  timestampString := strconv.Itoa(timestamp)
  snapshotName := fmt.Sprintf("%s-%s", config.Prefix, timestampString)

  vmrunOut := vmrun.GetOutputChannel()
	vmrunErr := vmrun.GetErrorChannel()
  go vmrun.Snapshot(config.TemplatePath, snapshotName)

	select {
	case <-vmrunOut:
	case err := <-vmrunErr:
		return err
	}
  return nil
}
