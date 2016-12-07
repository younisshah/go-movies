package omdb

import (
	"bytes"
	"net/http"
	"log"
)

const BASE_URL = "http://www.omdbapi.com/?"

func NewMovieRequest(title, year, movieType, plot string) (MovieRequest) {
	return MovieRequest{Title: title, Year: year, Plot: PlotType(plot), Type: MovieType(movieType)}
}

func (req MovieRequest)Search() /*(MovieResponse, error)*/ {
	buffer := new(bytes.Buffer)
	if req.Title == "" {
		log.Println("Title cannot be empty")
		//return nil, errors.New("Title cannot be empty")
	}
	buffer.WriteString(BASE_URL)
	buffer.WriteString("s=")
	buffer.WriteString(req.Title)

	log.Println(buffer.String())

	client := http.Client{}
	if req, err := http.NewRequest("GET", buffer.String(), nil); err != nil {
		log.Println("e", err)
	} else {
		req.Header.Add("content-type", "application/json; charset=utf-8")
		req.Header.Add("access-control-allow-origin", "*")
		log.Println(req.Header)
		if resp, err := client.Do(req); err != nil {
			log.Println(err)
		} else {
			log.Println(resp)
		}
	}
}
