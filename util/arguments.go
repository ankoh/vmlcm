package util

import (
	"flag"
	"fmt"
	"strconv"
)

// LCMCommand represents the primary command passed to vmlcm
// like "up 3" to use 3 clones
type LCMCommand int

const (
	// VerifyCommand checks the configuration of the LCM
	// The option shall be used to evaluate a working setup
	// * Checks whether the configuration file is valid
	// * Checks whether the template exists
	// * Checks whether the clones directory exists
	VerifyCommand LCMCommand = iota

	// StatusCommand returns the current state of the LCM
	// * Gathers information about running clones
	// * Gathers information about stopped/suspended clones
	// * Gathers information about the used template
	StatusCommand

	// UseCommand <number> ensures that specified the number of clones exists
	// * keep 10 with 10 existing == noop
	// * keep 10 with 4 clones existing == noop
	// * keep 4 with 10 clones existing == delete 6 clones
	// * keep 0 == delete all clones
	UseCommand

	// StartCommand starts all currently existing clones
	// This option shall be used when build agents are manually shutdown/suspended
	// or stopped with Stop
	StartCommand

	// StopCommand stops all currently existing clones
	// This option shall be used when build agents need to be stopped
	// in maintenance windows for example
	StopCommand

	// SnapshotCommand takes a snapshot of the template
	SnapshotCommand
)

// LCMArguments stores the options that have been passed to vmlcm
type LCMArguments struct {
	ConfigPath          *string    // vmlcm {-f agents.yml} up 3
	Command             LCMCommand // vmlcm -f agents.yml {up} 3
	CommandIntParameter int        // vmlcm -f agents.yml up {3}
	Test                bool       // vmlcm -test -f agents.yml verify
}

// ParseArguments parses the provided command line flags & aguments
func ParseArguments() (*LCMArguments, error) {
	// Parse flags
	configPath := flag.String("f", "", "path to the configuration file")
	flag.Parse()
	arguments := flag.Args()

	// Check if the flag has been provided
	if len(*configPath) == 0 {
		err := fmt.Errorf("You need to provide a configuration file with '-f'.")
		return nil, err
	}

	// Check if any argument has been provided
	if len(arguments) == 0 {
		err := fmt.Errorf("You have to provide one of these arguments: verify, status, start, stop, snapshot, use <number>")
		return nil, err
	}

	// Check if the command is valid
	commandString := arguments[0]
	var command LCMCommand

	switch commandString {
	case "verify":
		command = VerifyCommand
	case "status":
		command = StatusCommand
	case "use":
		command = UseCommand
	case "start":
		command = StartCommand
	case "stop":
		command = StopCommand
	case "snapshot":
		command = SnapshotCommand
	default:
		err := fmt.Errorf("Unknown command %s", commandString)
		return nil, err
	}

	var commandParameter int

	// If needed, check if command parameter has been provided
	if command == UseCommand {
		if len(arguments) <= 1 {
			err := fmt.Errorf("The command 'use' requires a number parameter. (vmlcm -f ./agents.yml use 3)")
			return nil, err
		}
		param, err := strconv.Atoi(arguments[1])
		if err != nil || param < 0 {
			err := fmt.Errorf("%s is not a valid parameter for the command 'use'", arguments[1])
			return nil, err
		}
		commandParameter = param
	}

	parameter := &LCMArguments{
		ConfigPath:          configPath,
		Command:             command,
		CommandIntParameter: commandParameter}

	return parameter, nil
}
