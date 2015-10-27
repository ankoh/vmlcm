package controller

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/ankoh/vmlcm/util"
	"github.com/ankoh/vmlcm/vmware"
)

var validPrefix = regexp.MustCompile("^[A-Za-z0-9]+$")

// Verify verifies the provided settings
func Verify(
	logger *util.Logger,
	vmrun vmware.VmrunWrapper,
	config *util.LCMConfiguration,
	silent bool) error {
	// Disable logger if not verbose
	var loggerSilent = logger.Silent
	logger.Silent = silent
	defer func() { logger.Silent = loggerSilent }()

	// First check the paths
	err := verifyConfigurationPaths(logger, config)
	if err != nil {
		return err
	}

	// Verify Prefix
	ok := validPrefix.MatchString(config.Prefix)
	if !ok {
		logger.LogVerification("Verifying prefix", false)
		return fmt.Errorf("Prefix must match the RegEx /^[A-Za-z0-9]+$/")
	}
	logger.LogVerification("Verifying prefix", true)

	// Verify Mac addresses
	err = testMacAddresses(config)
	if err != nil {
		logger.LogVerification("Verifying MAC addresses", false)
		return err
	}
	logger.LogVerification("Verifying MAC addresses", true)

	// Test vmrun help
	err = testVmrunHelp(vmrun)
	if err != nil {
		logger.LogVerification("Verifying vmrun executable", false)
		return err
	}
	logger.LogVerification("Verifying vmrun executable", true)

	// Test clone read
	err = testCloneRead(config)
	if err != nil {
		logger.LogVerification("Verifying clone list", false)
		return err
	}
	logger.LogVerification("Verifying clone list", true)

	// Test clone write
	err = testCloneWrite(config)
	if err != nil {
		logger.LogVerification("Verifying clone write", false)
		return err
	}
	logger.LogVerification("Verifying clone write", true)

	// Delete test file
	err = deleteTestFile(config)
	if err != nil {
		logger.LogVerification("Deleting test file", false)
		return err
	}
	logger.LogVerification("Deleting test file", true)

	return nil
}

func testMacAddresses(config *util.LCMConfiguration) error {
	addresses := config.Addresses
	for _, address := range addresses {
		if !isValidMacAddress(address) {
			return fmt.Errorf("Invalid Mac Address %s", address)
		}
	}
	return nil
}

// Check if vmrun help returns <<some>> output (and not an error)
func testVmrunHelp(vmrun vmware.VmrunWrapper) error {
	_, err := getRunningVMNumber(vmrun)
	if err != nil {
		return err
	}
	return nil
}

func testCloneRead(config *util.LCMConfiguration) error {
	// Try to read from the clones directory to check read permissions
	_, err := util.ListDirectory(config.ClonesDirectory)
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
var validMacRegEx = regexp.MustCompile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$")
var absolutePathRegEx = regexp.MustCompile("^/.*$")

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
func verifyConfigurationPaths(
	logger *util.Logger,
	config *util.LCMConfiguration) error {
	// Check Vmrun executable
	if !isValidPath(config.Vmrun) {
		logger.LogVerification("Verifying vmrun path", false)
		return fmt.Errorf("Invalid vmrun path: %s", config.Vmrun)
	}
	logger.LogVerification("Verifying vmrun path", true)

	// Check Clones directory
	if !isValidPath(config.ClonesDirectory) {
		logger.LogVerification("Verifying clones directroy", false)
		return fmt.Errorf("Invalid clones directory: %s", config.ClonesDirectory)
	}
	logger.LogVerification("Verifying clones directroy", true)

	// Check if Clones directory is a trailing slash
	matches, err := regexp.MatchString(".*/$", config.ClonesDirectory)
	if err != nil || !matches {
		logger.LogVerification("Verifying directory trailing slash", false)
		return fmt.Errorf("The clones directory path must have a trailing slash")
	}
	logger.LogVerification("Verifying directory trailing slash", true)

	// Check Template path
	if !isValidPath(config.TemplatePath) {
		logger.LogVerification("Verifying template path", false)
		return fmt.Errorf("Invalid template path: %s", config.TemplatePath)
	}
	logger.LogVerification("Verifying template path", true)

	// Check if the template path ends with vmx
	matches, err = regexp.MatchString(".*\\.vmx$", config.TemplatePath)
	if err != nil || !matches {
		logger.LogVerification("Verifying template extension", false)
		return fmt.Errorf("The template path must end with '.vmx'")
	}
	logger.LogVerification("Verifying template extension", true)

	return nil
}
