package util

import (
	"fmt"
)

// Log padding
var verificationFormat = "%-50s %s"

// LogVerification prints a validation message
func LogVerification(message string, successful bool) {
	var status = "[Failed]"
	if successful {
		status = "[Ok]"
	}
	fmt.Println(verificationFormat, message, status)
}
