package main

import (
	"os/exec"
	"os/user"
	"path"
	"strings"
)

// Searchresult represents the result of any local search.
type Searchresult struct {
	name      string
	fullpath  string
	thumbnail string
}

// Places represents the bookmarked locations of a Gnome desktop, e.g
// Videos, Documents, Music, Home... .
type Places struct {
	location string
}

// Bookmarks is a slice of PLaces
type Bookmarks []Places

// Binaries is a slice of Places reprenting the file directories holding
// binaries.
type Binaries []Places

var bookmarks = Bookmarks{
	{location: ""},
	{location: "/Documents"},
	{location: "/Downloads"},
	{location: "/Music"},
	{location: "/Pictures"},
	{location: "/Videos"},
}
var binaries = Binaries{
	{location: "/usr/bin"},
	{location: "/usr/local/bin"},
	{location: "/opt"},
}

// findQuery uses the 'find' command to search a given string
// in an array of Places.
// func findQuery(query string) ([]Searchresult, error) {
// }

// findCommand returns a Cmd struct for the find Command
// to search in a given location for a given value.
func findCommandBookmarks(loc, value string) (*exec.Cmd, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	return exec.Command("find", usr.HomeDir+loc, "-iname", "'*"+value+"*'"), nil
}

// locateCommand returns a Cmd struct for the locate Command.
// locate's output is limited to 20 results, case is ignored and
// only the base name of the path is matched.
func locateCommand(value string) *exec.Cmd {
	return exec.Command("locate", "-l", "20", "-b", "-i", value)
}

// NewSearchResult constructs from the output of a query a
// Searchresult struct with its name and path initialized.
func NewSearchResult(line string) *Searchresult {
	var sr Searchresult
	_, file := path.Split(line)
	sr.name = strings.Split(file, ".")[0]
	sr.fullpath = line
	return &sr

}
