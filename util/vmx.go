package util

import (
  "fmt"
  "io/ioutil"
  "regexp"
  "bytes"

)

// Pattern that matches the eth0 connection type in a vmx file
var eth0ConnectionTypePattern = regexp.MustCompile(
  "ethernet0.connectionType = \"[a-zA-Z]+\"",
)

// Pattern that matches the eth0 address in a vmx file
var eth0AddressPattern = regexp.MustCompile(
  "ethernet0.address = \"(([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2}))\"",
)

// UpdateVMX updates a given vmx configuration file
// It sets the ethernet0.connectionType and the ethernet0.address
// This approach is faster and much more reliable th
func UpdateVMX(src string, dst string, address string) error {
  // First read the file
  fileBytes, fileError := ioutil.ReadFile(src)
  if fileError != nil {
    err := fmt.Errorf("Failed to read the file at path %s.", src)
    return err
  }

  eth0ConnectionType := []byte("ethernet0.connectionType = \"bridged\"")
  eth0Address := []byte(fmt.Sprintf("ethernet0.address = \"%s\"", address))

  // The connectionType should be included already
  // Just always set the connectionType to bridged
  fileBytes = eth0ConnectionTypePattern.ReplaceAll(
    fileBytes, eth0ConnectionType)

  // Then check if it has the eth0Address defined
  if !eth0AddressPattern.Match(fileBytes) {
    // Simply append the Mac address setting if its not
    fileBuffer := new(bytes.Buffer)
    fileBuffer.Write(fileBytes)
    fileBuffer.Write(eth0Address)
    fileBuffer.WriteString("\n")
    fileBytes = fileBuffer.Bytes()
  } else {
    // Otherwise we need to replace as well
    fileBytes = eth0AddressPattern.ReplaceAll(
      fileBytes, eth0Address)
  }

  // Finally write out the new configuration file
  writeError := ioutil.WriteFile(dst, fileBytes, 0755)
  if writeError != nil {
    err := fmt.Errorf("Failed to write the file at path %s.", dst)
    return err
  }
  return nil
}
