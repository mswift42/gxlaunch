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
