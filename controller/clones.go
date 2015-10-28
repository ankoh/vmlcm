package controller

import(
  "fmt"
  "regexp"
  "math"
  "os"

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
    _, err = cloneUpTo(vmrun, config, clones, use)
    return err
  }

  // Check if clones need to be deleted
  if len(clones) > use {
    _, err = deleteUpTo(vmrun, config, clones, use)
  }

  return nil
}

// Deletes up to <<use>> clones
func deleteUpTo(
  vmrun vmware.VmrunWrapper,
  config *util.LCMConfiguration,
  clones []*virtualMachine,
  use int) ([]string, error) {
  vmrunOut := vmrun.GetOutputChannel()
	vmrunErr := vmrun.GetErrorChannel()

  // First order the clones by status
  clonesOrdered := orderClonesByRunning(clones)
  toDelete := int(math.Min(
    float64(len(clonesOrdered)),
    math.Abs(float64(len(clones) - use))))

  var deletedClones []string

  // Then shutdown target clones
  for i := 0; i < toDelete; i++ {
    clone := clones[i]

    // First try to stop the clone if it is running
    if clone.running {
      // First try a soft stop
      vmrun.Stop(clone.path, false)
      succeeded := true
      select {
      case <- vmrunOut:
      case <- vmrunErr:
        succeeded = false
      }
      if !succeeded {
        // then try a hard stop
        vmrun.Stop(clone.path, false)
        select {
        case <- vmrunOut:
        case err := <- vmrunErr:
          // failed to shutdown the VM
          return nil, err
        }
      }
    }

    // Then delete the clone
    os.RemoveAll(clone.path)
    // TODO: The vmware folder still exsits
    deletedClones = append(deletedClones, clone.path)
  }

  return deletedClones, nil
}

// orderClonesByRunning orders the clones by their running state
// Offline virtual machines will come first, then the online ones
func orderClonesByRunning(clones []*virtualMachine) []*virtualMachine {
  var result []*virtualMachine
  // First append offline machines
  for _, clone := range clones {
    if !clone.running {
        result = append(result, clone)
    }
  }
  // Then append online machines
  for _, clone := range clones {
    if clone.running {
      result = append(result, clone)
    }
  }
  return result
}

// getAvailableMacAddresses checks which of the mac addresses is free
func getAvailableMacAddresses(
  clones []*virtualMachine,
  config *util.LCMConfiguration) []string {
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
  // Create slice with remaining addresses
  var remainingAddresses []string
  for address := range availableAddresses {
    remainingAddresses = append(remainingAddresses, address)
  }
  return remainingAddresses
}

// cloneUpTo clones the template up to <<use>> times
func cloneUpTo(
  vmrun vmware.VmrunWrapper,
  config *util.LCMConfiguration,
  clones []*virtualMachine,
  use int) ([]string, error) {
  // First get available mac addresses
  availableAddresses := getAvailableMacAddresses(clones, config)

  // calculate how many clones need to be created
  toCreate := int(math.Min(
      float64(len(availableAddresses)),
      math.Abs(float64(use - len(clones)))))
  if toCreate <= 0 {
    return nil, nil
  }

  // Get snapshot for the clones
  snapshot, err := prepareTemplateSnapshot(vmrun, config)
  if err != nil {
    return nil, err
  }

  var created []string

  // Now fire the clone command for each of the addresses
  for i := 0; i < toCreate; i++ {
    address := availableAddresses[i]
    vmID := util.MacAddressToVMId(address)
    vmName := fmt.Sprintf("%s-%s", config.Prefix, vmID)

    err := cloneTemplate(vmrun, config, vmName, snapshot)
    if err != nil {
      return nil, err
    }
    created = append(created, vmName)
  }
  return created, nil
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
