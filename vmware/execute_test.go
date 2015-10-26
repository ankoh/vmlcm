package vmware_test

import (
	"testing"
	"fmt"
  "gitlab.kohn.io/ankoh/vmlcm/vmware"
)

func TestExecute(t *testing.T) {
  // Should be the case...

  outChan := make(chan string)
  errChan := make(chan string)
  defer close(outChan)
  defer close(errChan)

  // Dont call as goroutine as goconvey does not support it
  go vmware.ExecuteCommand(outChan, errChan, "ls", "-la", "./")

  select {
    case out := <- outChan:
      fmt.Printf(out)
    case err := <- errChan:
			fmt.Printf(err)
      t.Fail()
  }
}
