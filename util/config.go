package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// LCMConfiguration holds the configuration for the linked clones manager
// ClonesDirectory: Directory for the clones
// TemplateVMX: VMX of the template that is used to create the linked clones
// MACAddresses: Addresses that are used for the linked clones
type LCMConfiguration struct {
	Vmrun           string
	ClonesDirectory string
	TemplatePath    string
	Addresses       []string
	Prefix          string
}

// ParseConfiguration uses a string path to a json file
func ParseConfiguration(path string) (*LCMConfiguration, error) {
	// First read the file
	fileString, fileError := ioutil.ReadFile(path)
	if fileError != nil {
		err := fmt.Errorf("Failed to read the file at path %s.", path)
		return nil, err
	}

	// Then parse the json
	var config LCMConfiguration
	jsonError := json.Unmarshal([]byte(fileString), &config)
	if jsonError != nil {
		err := fmt.Errorf("Failed to parse the json configuration at path %s.", path)
		return nil, err
	}

	// Now check the values of the strings.
	// They must not be empty
	if len(config.Vmrun) == 0 {
		err := fmt.Errorf("The configuration file does not contain a valid parameter 'Vmrun'")
		return nil, err
	}
	if len(config.ClonesDirectory) == 0 {
		err := fmt.Errorf("The configuration file does not contain a valid parameter 'ClonesDirectory'")
		return nil, err
	}
	if len(config.TemplatePath) == 0 {
		err := fmt.Errorf("The configuration file does not contain a valid parameter 'TemplatePath'")
		return nil, err
	}
	if len(config.Prefix) == 0 {
		err := fmt.Errorf("The configuration file does not contain a valid parameter 'Prefix'")
		return nil, err
	}

	// Uppercase all addresses
	for i, address := range config.Addresses {
		config.Addresses[i] = strings.ToUpper(address)
	}

	// If Unmarshal was successfull we're done
	return &config, nil
}
