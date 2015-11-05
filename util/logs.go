package util

import (
	"fmt"
	"bytes"
)

// CreateVerification creates a verification log message
func CreateVerification(message string, successful bool) string {
	status := "Failed"
	statusColor := ColorLightGray
	if successful {
		status = "Ok"
		statusColor = ColorCyan
	}
	return fmt.Sprintf("%-40s [%s%s%s]\n",
		message, statusColor, status, ColorNone)
}

// WriteVerification writes a verification message to an output buffer
func WriteVerification(buffer *bytes.Buffer, message string, successful bool) {
	// If buffer is nil abort
	// Someone wants us to be silent
	if buffer == nil {
		return
	}

	status := "Failed"
	statusColor := ColorLightGray
	if successful {
		status = "Ok"
		statusColor = ColorCyan
	}
	buffer.WriteString(fmt.Sprintf("%-40s [%s%s%s]\n",
		message, statusColor, status, ColorNone))
}
