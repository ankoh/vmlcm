package util

import (
	"fmt"
)

// LogVerification prints a validation message
func LogVerification(message string, successful bool) {
	var status = "[Failed]"
	if successful {
		status = "[Ok]"
	}
	fmt.Printf("\t%-50s %s\n", message, status)
}
