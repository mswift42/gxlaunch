package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCommand(t *testing.T) {
	assert := assert.New(t)
	cmd := findCommand("/", "hallo")
	assert.Equal(cmd.Path, "/usr/bin/find")
}
