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

	// UpCommand <number> ensures that the specified amount of clones is <<up>>
	// * up 10 creates 10 if no clones are existing
	// * up 10 creates 6 clones if 4 clones are existing
	// * up 0 creates 0 clones
	UpCommand

	// KeepCommand <number> keeps the number of clones and deletes the rest
	// * keep 10 with 10 existing == noop
	// * keep 10 with 4 clones existing == noop
	// * keep 4 with 10 clones existing == delete 6 clones
	// * keep 0 == delete all clones
	KeepCommand

	// ResetCommand resets all clones
	// This is probably faster than up 0 -> up 10
	ResetCommand

	// StartCommand starts all currently existing clones
	// This option shall be used when build agents are manually shutdown/suspended
	// or stopped with Stop
	StartCommand

	// StopCommand stops all currently existing clones
	// This option shall be used when build agents need to be stopped
	// in maintenance windows for example
	StopCommand

	// SuspendCommand suspends all currently existing clones
	// Similar to stop this option stops all clones in maintenance windows while
	// maintaining the vm state
	SuspendCommand
)

// LCMArguments stores the options that have been passed to vmlcm
type LCMArguments struct {
	ConfigPath       *string    // vmlcm {-f agents.yml} up 3
	Command          LCMCommand // vmlcm -f agents.yml {up} 3
	CommandParameter int        // vmlcm -f agents.yml up {3}
	Test						 bool				// vmlcm -test -f agents.yml verify
}

// ParseArguments parses the provided command line flags & aguments
func ParseArguments() (*LCMArguments, error) {
	// Parse flags
	configPath := flag.String("f", "", "path to the configuration file")
	flag.Parse()
  arguments := flag.Args()

	// Check if the flag has been provided
	if len(*configPath) == 0 {
		err := fmt.Errorf("You need to provide a valid path of the configuration file.")
		return nil, err
	}

	// Check if any argument has been provided
	if len(arguments) == 0 {
		err := fmt.Errorf("You have to provide one of the arguments {verify, status, reset, start, stop, suspend, up <number>}.")
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
	case "up":
		command = UpCommand
	case "keep":
		command = KeepCommand
	case "reset":
		command = ResetCommand
	case "start":
		command = StartCommand
	case "stop":
		command = StopCommand
	case "suspend":
		command = SuspendCommand
	default:
		err := fmt.Errorf("Unknown command %s", commandString)
		return nil, err
	}

	var commandParameter int

	// If needed, check if command parameter has been provided
	if command == UpCommand || command == KeepCommand {
		if len(arguments) <= 1 {
			err := fmt.Errorf("The commands up and keep require a number parameter. (vmlcm -f ./agents.yml up 3)")
      return nil, err
		}
		param, err := strconv.Atoi(arguments[1])
		if err != nil || param < 0 {
			err := fmt.Errorf("%s is not a valid parameter for the commands 'up' and 'keep'", arguments[1])
      return nil, err
		}
		commandParameter = param
	}

	parameter := &LCMArguments{
		ConfigPath:       configPath,
		Command:          command,
		CommandParameter: commandParameter}

	return parameter, nil
}
