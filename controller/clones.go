package controller

/*import(
  "fmt"

  "github.com/ankoh/vmlcm/util"
  "github.com/ankoh/vmlcm/vmware"
)

// Use ensures that <<use>> number of clones exist
func Keep(
  vmrun vmware.VmrunWrapper,
  config *util.LCMConfiguration,
  use int) (error) {
  vmrunOut := vmrun.GetOutputChannel()
	vmrunErr := vmrun.GetErrorChannel()

  // Check use parameter
  if use < 0 {
    return fmt.Errorf("Parameter must be >= zero")
  }

  // First check mac addresses
  if len(config.Addresses) < up {
    return fmt.Errorf("Tried to use %d clones with only %d mac addresses",
      use, len(config.Addresses))
  }

  // Then check the number of clones
  // Get all existing clones
  clones, err := getClones(vmrun, config)
  if err != nil {
    return err
  }

  // Check if number of existing clones equals the parameter (== noop)
  if len(clones) == up {
    return nil
  }

  // Check if clones need to be created
  if len(clones) < up {


    for i := 0; i < up - clones; i++ {
      snapshot, err := prepareTemplateSnapshot(vmrun, config)
      if err != nil {
        return err
      }

    }
  }

  // Check if clones need to be deleted
  if len(clones) > up {

  }

  return nil, nil
}

// createTemplateSnapshot creates a prefixed snapshot of the template
func cloneTemplate(
  vmrun vmware.VmrunWrapper,
  config *util.LCMConfiguration) (string, error) {

  // Create snapshot name
  snapshotName := fmt.Sprintf("%s-%d", config.Prefix, timestamp)

  vmrunOut := vmrun.GetOutputChannel()
	vmrunErr := vmrun.GetErrorChannel()
  go vmrun.Snapshot(config.TemplatePath, snapshotName)

	select {
	case <-vmrunOut:
	case err := <-vmrunErr:
		return "", err
	}
  return snapshotName, nil
}*/
