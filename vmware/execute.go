package vmware

import (
	"fmt"
	"os/exec"
)

// executeCommand executes a shell command and pipes the result to channels
// When finished, call back the waitGroup
// That will enable nice dot-progress indicators in the command line
func executeCommand(
	outputChannel chan string,
	errorChannel chan error,
	binary string,
	arguments ...string) {

	// Run the command
	out, err := exec.Command(binary, arguments...).Output()

	// Pipe to respective channels
	if err != nil {
		errorChannel <- err
	} else {
		outputChannel <- fmt.Sprintf("%s", out)
	}
}

// forceExecuteCommand force-executes a shell command and pipes the result
// to the outputchannel. errors are ignored
func forceExecuteCommand(
	outputChannel chan string,
	binary string,
	arguments ...string) {
	out, _ := exec.Command(binary, arguments...).Output()
	outputChannel <- fmt.Sprintf("%s", out)
}
