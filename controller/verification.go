package controller

import (
	"fmt"
	"os"
	"regexp"
  "io/ioutil"

	"gitlab.kohn.io/ankoh/vmlcm/util"
	"gitlab.kohn.io/ankoh/vmlcm/vmware"
)

// Verify verifies the provided settings
func Verify(
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration) error {
  // First check the paths
  err := verifyConfigurationPaths(config)
  if err != nil {
    return err
  }

  // Test vmrun help
  err = testVmrunHelp(vmrun)
  if err != nil {
    util.LogVerification("Testing vmrun executable", false)
    return err
  }
  util.LogVerification("Testing vmrun executable", true)

  // Test clone read
  err = testCloneRead(config)
  if err != nil {
    util.LogVerification("Testing clone list", false)
    return err
  }
  util.LogVerification("Testing clone list", true)

  // Test clone write
  err = testCloneWrite(config)
  if err != nil {
    util.LogVerification("Testing clone write", false)
    return err
  }
  util.LogVerification("Testing clone write", true)

  // Delete test file
  err = deleteTestFile(config)
  if err != nil {
    util.LogVerification("Deleting test file", false)
    return err
  }
  util.LogVerification("Deleting test file", true)
  return nil
}

// Check if vmrun help returns <<some>> output (and not an error)
func testVmrunHelp(vmrun vmware.VmrunWrapper) error {
  vmrunOut := vmrun.GetOutputChannel()
  vmrunErr := vmrun.GetErrorChannel()
  go vmrun.Help()
  select {
    case <- vmrunOut:
    case err := <- vmrunErr:
      return err
  }
  return nil
}

func testCloneRead(config *util.LCMConfiguration) error {
  // Try to read from the clones directory to check read permissions
  _, err := listDirectory(config.ClonesDirectory)
  return err
}

// Tries to create a dummy file in the clones directory to check write permissions
func testCloneWrite(config *util.LCMConfiguration) error {
  testBuffer := []byte("vmlcm write test\n")
  testFilePath := fmt.Sprintf("%s%s", config.ClonesDirectory, "test")
  err := ioutil.WriteFile(testFilePath, testBuffer, 0644)
  return err
}

// Deletes the test file again
func deleteTestFile(config *util.LCMConfiguration) error {
  testFilePath := fmt.Sprintf("%s%s", config.ClonesDirectory, "test")
  err := os.Remove(testFilePath)
  return err
}

// Regular expressions
var validMacRegEx, _ = regexp.Compile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$")
var absolutePathRegEx, _ = regexp.Compile("^/.*$")

// Returns whether the given address is a valid mac address
func isValidMacAddress(address string) bool {
	if validMacRegEx == nil {
		return false
	}
	return validMacRegEx.MatchString(address)
}

// Returns whether the given path is an absolute path
func isAbsolutePath(path string) bool {
	if absolutePathRegEx == nil {
		return false
	}
	return absolutePathRegEx.MatchString(path)
}

// Returns whether the given path is valid (absolute and existing)
func isValidPath(path string) bool {
	if !isAbsolutePath(path) {
		return false
	}
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

// validateConfigurationPaths validates a given LCM configuration
func verifyConfigurationPaths(config *util.LCMConfiguration) error {
	// Check Vmrun executable
	if !isValidPath(config.Vmrun) {
		util.LogVerification("Verifying vmrun path", false)
		return fmt.Errorf("Invalid vmrun path: %s", config.Vmrun)
	}
	util.LogVerification("Verifying vmrun path", true)

	// Check Clones directory
	if !isValidPath(config.ClonesDirectory) {
		util.LogVerification("Verifying clones directroy", false)
		return fmt.Errorf("Invalid clones directory: %s", config.ClonesDirectory)
	}
	util.LogVerification("Verifying clones directroy", true)

  // Check if Clones directory is a trailing slash
  matches, _ := regexp.MatchString(".*/$", config.ClonesDirectory)
  if !matches {
    util.LogVerification("Verifying directory trailing slash", false)
    return fmt.Errorf("The clones directory path must have a trailing slash")
  }
  util.LogVerification("Verifying directory trailing slash", true)

	// Check Template path
	if !isValidPath(config.TemplatePath) {
		util.LogVerification("Verifying template path", false)
		return fmt.Errorf("Invalid template path: %s", config.TemplatePath)
	}
	util.LogVerification("Verifying template path", true)

	// Check if the template path ends with vmx
	matches, _ = regexp.MatchString(".*\\.vmx$", config.TemplatePath)
	if !matches {
		util.LogVerification("Verifying template extension", false)
		return fmt.Errorf("The template path must end with '.vmx'")
	}
	util.LogVerification("Verifying template extension", true)

	return nil
}