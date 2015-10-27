package controller

import(
  "fmt"
  "regexp"
  "math"

  "github.com/ankoh/vmlcm/util"
  "github.com/ankoh/vmlcm/vmware"
)

var vmxName = regexp.MustCompile("([A-Za-z0-9-]+\\.vmx)$")

// Use ensures that <<use>> number of clones exist
func Use(
  vmrun vmware.VmrunWrapper,
  config *util.LCMConfiguration,
  use int) error {
  // Check use parameter
  if use < 0 {
    return fmt.Errorf("Parameter must be >= zero")
  }

  // First check mac addresses
  if len(config.Addresses) < use {
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
  if len(clones) == use {
    return nil
  }

  // Check if clones need to be created
  if len(clones) < use {
    // Get snapshot for the clones
    snapshot, err := prepareTemplateSnapshot(vmrun, config)
    if err != nil {
      return err
    }

    // Now we need to find the addresses that we can use

    // Create map with addresses
    availableAddresses := make(map[string]bool)
    for _, address := range config.Addresses {
      availableAddresses[address] = true
    }

    // Loop through clones and delete addresses that are used
    for _, clone := range clones {
      matches := vmxName.FindStringSubmatch(clone.path)
      if len(matches) != 2 {
        continue
      }
      macAddress, err := util.VMNameToMacAddress(matches[1])
      if err != nil {
        continue
      }
      delete(availableAddresses, macAddress)
    }

    // Remaining address number
    remaining := len(availableAddresses)
    toCreate := int(math.Min(float64(remaining), float64(use - len(clones))))

    // Create slice with remaining addresses
    var remainingAddresses []string
    for address := range availableAddresses {
      remainingAddresses = append(remainingAddresses, address)
    }

    // Now fire the clone command for each of the addresses
    for i := 0; i < toCreate; i++ {
      address := remainingAddresses[i]
      vmName := util.MacAddressToVMId(address)

      err := cloneTemplate(vmrun, config, vmName, snapshot)
      if err != nil {
        return err
      }
    }
  }

  // TODO: Delete
  // Check if clones need to be deleted
  /*if len(clones) > use {

  }*/

  return nil
}

// createTemplateSnapshot creates a prefixed snapshot of the template
func cloneTemplate(
  vmrun vmware.VmrunWrapper,
  config *util.LCMConfiguration,
  cloneName string,
  snapshot string) error {

  vmrunOut := vmrun.GetOutputChannel()
	vmrunErr := vmrun.GetErrorChannel()
  go vmrun.CloneLinked(
    config.TemplatePath,
		config.ClonesDirectory,
    cloneName, snapshot)

	select {
	case <-vmrunOut:
	case err := <-vmrunErr:
		return err
	}
  return  nil
}
