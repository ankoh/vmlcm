package vmware

import(
  "fmt"
  "os/exec"
)

// ExecuteCommand executes a shell command and pipes the result to channels
// When finished, call back the waitGroup
// That will enable nice dot-progress indicators in the command line
func ExecuteCommand(
  outputChannel chan string,
  errorChannel chan string,
  binary string,
  arguments ...string) {

  // Run the command
  out, err := exec.Command(binary, arguments...).Output()

  // Pipe to respective channels
  if err != nil {
    errorChannel <- fmt.Sprintf("%s", err.Error())
  } else {
    outputChannel <- fmt.Sprintf("%s", out)
  }
}
