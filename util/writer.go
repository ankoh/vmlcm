package util

import (
	"fmt"
	"bytes"
)

// TryWrite writes a message to an output buffer if the buffer is not a
// nil pointer
func TryWrite(buffer *bytes.Buffer, message string) {
	if buffer == nil {
		return
	}
	buffer.WriteString(message)
}

// TryWriteln writes a message to an output buffer if the buffer is not a
// nil pointer
func TryWriteln(buffer *bytes.Buffer, message string) {
	if buffer == nil {
		return
	}
	buffer.WriteString(fmt.Sprintf("%s\n", message))
}

// TryWrite2Columns tries to write a status message to an output buffer
func TryWrite2Columns(buffer *bytes.Buffer, c1width int, c1 string, c2 string) {
	if buffer == nil {
		return
	}
	formatBuffer := new(bytes.Buffer)
	formatBuffer.WriteString("%-")
	formatBuffer.WriteString(fmt.Sprint(c1width))
	formatBuffer.WriteString("s %s%s%s\n")
	buffer.WriteString(fmt.Sprintf(formatBuffer.String(), c1,
		ColorCyan, c2, ColorNone))
}

// TryWriteVerification tries to write a verification message to an output buffer
func TryWriteVerification(buffer *bytes.Buffer, message string, successful bool) {
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
