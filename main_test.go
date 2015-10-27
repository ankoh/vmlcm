package main

import testing "testing"
import assert "github.com/stretchr/testify/assert"

func TestMain(t *testing.T) {
	assert := assert.New(t)

	// assert equality
	assert.Equal(123, 123, "they should be equal")
}
