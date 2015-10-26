package controller

import (
  "os"
  "fmt"
)

// Verify verifies the provided settings
func Verify() {

}

func isValidMacAddress(address string) bool {
  return true
}

func isAbsolutePath(path string) bool {
  return true
}

func verifyPath(path string) error {
  if !isAbsolutePath(path) {
    return fmt.Errorf("Only absolute paths are allowed [%s]", path)
  }
  if _, err := os.Stat(path); err != nil {
    return fmt.Errorf("Invalid path [%s]", path)
  }
  return nil
}
