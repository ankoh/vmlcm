package util

import (
  "os"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestParseArguments(t *testing.T) {
  assert := assert.New(t);

  os.Args = []string {
    "vmlcm",
    "-f", "./agents.json",
    "verify"}

  args, err := ParseArguments()
  assert.Nil(err)
  assert.NotNil(args)
}
