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
	var statusColor = ColorLightGray
	if successful {
		status = "Ok"
		statusColor = ColorCyan
	}
	fmt.Printf("%-40s [%s%s%s]\n", message, statusColor, status, ColorNone)
}
