package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCommand(t *testing.T) {
	assert := assert.New(t)
	cmd, err := findCommand("", "hallo")
	if err != nil {
		panic(err)
	}
	assert.Equal(cmd.Path, "/usr/bin/find")
	assert.Equal(cmd.Args, []string{"find", "/home/severin", "-iname", "'*hallo*'"})
}

func TestLocateCommand(t *testing.T) {
	assert := assert.New(t)
	cmd := locateCommand("hallo")
	assert.Equal(cmd.Path, "/usr/bin/locate")
	assert.Equal(cmd.Args, []string{"locate", "-l", "20",
		"-b", "-i", "hallo"})
}
