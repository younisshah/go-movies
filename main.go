package main

import "github.com/younisshah/go-movies/omdb"

func main() {
	req := omdb.NewMovieRequest("Game of thrones", "", "", "")
	req.Search()
}
