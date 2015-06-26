package main

// Searchresult represents the result of any local search.
type Searchresult struct {
	name      string
	path      string
	thumbnail string
}

// Places represents the bookmarked locations of a Gnome desktop, e.g
// Videos, Documents, Music, Home... .
type Places struct {
	location string
}

// Bookmarks is a slice of PLaces
type Bookmarks []Places

var bookmarks = Bookmarks{
	{location: "/"},
	{location: "/Documents"},
	{location: "/Downloads"},
	{location: "/Music"},
	{location: "/Pictures"},
	{location: "/Videos"},
}

// findQuery uses the 'find' command to search a given string
// in an array of Places.
func findQuery(query string) (Searchresult, error) {
}
