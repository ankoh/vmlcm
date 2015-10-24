package util

// import "encoding/json"

// LCMConfiguration holds the configuration for the linked clones manager
// ClonesDirectory: Directory for the clones
// TemplateVMX: VMX of the template that is used to create the linked clones
// MACAddresses: Addresses that are used for the linked clones
type LCMConfiguration struct {
  VmrunExecutable     string
  ClonesDirectory     string
  TemplatePath        string
  MACAddresses        []string
}

// ParseConfiguration uses a string path to a json file
func ParseConfiguration(path string) *LCMConfiguration {
  return nil
}
