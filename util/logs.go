package util

import (
	"fmt"
)

// Logger allows to log messages in a unified way
type Logger struct {
  Silent bool
}

// NewLogger creates a new Logger
func NewLogger() *Logger {
  logger := new(Logger)
  logger.Silent = false
  return logger
}

// LogVerification prints a validation message
func (logger *Logger) LogVerification(message string, successful bool) {
  if logger.Silent {
    return
  }

	var status = "[Failed]"
	if successful {
		status = "[Ok]"
	}
	fmt.Printf("%-50s %s\n", message, status)
}
