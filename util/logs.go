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

	var status = "Failed"
	var statusColor = ColorRed
	if successful {
		status = "Ok"
		statusColor = ColorGreen
	}
	fmt.Printf("%-50s [%s%s%s]\n", message, statusColor, status, ColorNone)
}
