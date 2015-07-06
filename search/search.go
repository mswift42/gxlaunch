package search

import (
	"bufio"
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

// Searchresults is a slice of Searchresults
type Searchresults []Searchresult

// NameList is a slice listing the field name for every Searchresult.
func (s Searchresults) NameList() []string {
	results := make([]string, len(s))
	for i := range s {
		results[i] = s[i].name
	}
	return results
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

// Search appends the results of calling findQuery and locateQuery for a given
// query string, and returns them.
func Search(query string) Searchresults {
	results := make([]Searchresult, 0)
	results = append(results, FindQuery(query)...)
	// results = append(results, LocateQuery(query)...)
	return results
}

// FindQuery uses the 'find' command to search a given string
// in an array of Places.
func FindQuery(query string) []Searchresult {
	results := make([]Searchresult, 0)
	c := make(chan []Searchresult)
	go findbinaries(query, c)
	go findbookmarks(query, c)
	bin, book := <-c, <-c
	results = append(results, bin...)
	results = append(results, book...)
	return results
}

// LocateQuery runs the locate command for a query string and returns
// a slice of []Searchresult with its results.
func LocateQuery(query string) []Searchresult {
	c := make(chan []Searchresult, 0)
	go commandOutput(locateCommand(query), c)
	res := <-c
	return res
}

func findbinaries(query string, c chan []Searchresult) {
	for _, i := range binaries {
		go findCommandOutput(findCommandBinaries(i.location, query), c)
	}
}

func findbookmarks(query string, c chan []Searchresult) {
	for _, i := range bookmarks {
		findbook, err := findCommandBookmarks(i.location, query)
		if err != nil {
			panic(err)
		}
		go findCommandOutput(findbook, c)
	}
}

// commandOutput runs an exec.Cmd, builds for every line of the output
// a new Searchresult, and passes these into channel c.
func commandOutput(cmd *exec.Cmd, c chan []Searchresult) {
	out, _ := cmd.Output()
	res := make([]Searchresult, 0)
	split := strings.Split(string(out), "\n")
	for _, i := range split {
		sr := NewSearchResult(i)
		res = append(res, *sr)
	}
	c <- res
}

func findCommandOutput(cmd *exec.Cmd, c chan []Searchresult) {
	results := make([]Searchresult, 0)
	head := exec.Command("head", "-5")
	head.Stdin, _ = cmd.StdoutPipe()
	reader, err := head.StdoutPipe()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			res := NewSearchResult(scanner.Text())
			results = append(results, *res)
		}
	}()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	if err := head.Start(); err != nil {
		panic(err)
	}
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
	c <- results
}

// findCommandBookmarks returns a Cmd struct for the find Command
// to search in a given location for a given value.
func findCommandBookmarks(loc, value string) (*exec.Cmd, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	if loc == "" {
		return exec.Command("find", usr.HomeDir+loc, "-maxdepth", "1",
			"-iname", "*"+value+"*"), nil
	}
	return exec.Command("find", usr.HomeDir+loc, "-maxdepth", "2",
		"-iname", "*"+value+"*"), nil
}

func findCommandBinaries(loc, value string) *exec.Cmd {
	return exec.Command("find", loc, "-maxdepth", "2", "-iname", "*"+value+"*")
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
