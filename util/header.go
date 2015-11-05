package util

import (
	"bytes"
)

// GenerateASCIIHeader prints an ASCII header
func GenerateASCIIHeader() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("                    __             \n")
	buffer.WriteString("   _   ______ ___  / /________ ___ \n")
	buffer.WriteString("  | | / / __ `__ \\/ / ___/ __ `__ \\\n")
	buffer.WriteString("  | |/ / / / / / / / /__/ / / / / /\n")
	buffer.WriteString("  |___/_/ /_/ /_/_/\\___/_/ /_/ /_/ \n")
	buffer.WriteString("                                   \n")
	return buffer.String()
}
