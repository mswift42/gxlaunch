package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameList(t *testing.T) {
	sr1 := NewSearchResult("Documents/GoBook.pdf")
	sr2 := NewSearchResult("Documents/Dive_Into_Python3.pdf")
	results := Searchresults{*sr1, *sr2}
	assert := assert.New(t)
	names := results.NameList()
	assert.Equal(len(names), 2)
	assert.Equal(names[0], "GoBook")
	assert.Equal(names[1], "Dive_Into_Python3")

}

func TestFindCommand(t *testing.T) {
	assert := assert.New(t)
	cmd, err := findCommandBookmarks("", "hallo")
	if err != nil {
		panic(err)
	}
	assert.Equal(cmd.Path, "/usr/bin/find")
	assert.Equal(cmd.Args, []string{"find", "/home/severin", "-maxdepth", "1", "-iname", "*hallo*"})
	cmd2, err := findCommandBookmarks("/Documents", "hallo")
	if err != nil {
		panic(err)
	}
	assert.Equal(cmd2.Path, "/usr/bin/find")
	assert.Equal(cmd2.Args, []string{"find", "/home/severin/Documents", "-maxdepth", "2",
		"-iname", "*hallo*"})
}

func TestFindCommandBinaries(t *testing.T) {
	assert := assert.New(t)
	cmd := findCommandBinaries("/usr/bin", "hallo")
	assert.Equal(cmd.Path, "/usr/bin/find")
	assert.Equal(cmd.Args, []string{"find", "/usr/bin", "-maxdepth", "2",
		"-iname", "*hallo*"})
}

func TestLocateCommand(t *testing.T) {
	assert := assert.New(t)
	cmd := locateCommand("hallo")
	assert.Equal(cmd.Path, "/usr/bin/locate")
	assert.Equal(cmd.Args, []string{"locate", "-l", "20",
		"-b", "-i", "hallo"})
}
func TestNewSearchResult(t *testing.T) {
	assert := assert.New(t)
	sr := NewSearchResult("/home/severin/Documents/GoBook.pdf")
	assert.Equal("GoBook", sr.name)
	assert.Equal(sr.fullpath, "/home/severin/Documents/GoBook.pdf")

}
