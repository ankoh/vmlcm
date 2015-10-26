package controller

import (
  "os"
  "regexp"
)

// Verify verifies the provided settings
func Verify() {

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

func isValidPath(path string) bool {
  if !isAbsolutePath(path) {
    return false
  }
  if _, err := os.Stat(path); err != nil {
    return false
  }
  return true
}
