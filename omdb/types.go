package omdb

type MovieRequest struct {
	Title, Year string
	Type MovieType
	Plot PlotType
}

type PlotType string
type MovieType string

const (
	SHORT PlotType = "short"
	FULL PlotType = "full"
	MOVIE MovieType = "movie"
	SERIES MovieType = "series"
	EPISODE MovieType = "episode"
)

type MovieResponse struct {
	Title        string `json:"Title"`
	Year         string `json:"Year"`
	Rated        string `json:"Rated"`
	Released     string `json:"Released"`
	Runtime      string `json:"Runtime"`
	Genre        string `json:"Genre"`
	Director     string `json:"Director"`
	Writer       string `json:"Writer"`
	Actors       string `json:"Actors"`
	Plot         string `json:"Plot"`
	Language     string `json:"Language"`
	Country      string `json:"Country"`
	Awards       string `json:"Awards"`
	Poster       string `json:"Poster"`
	Metascore    string `json:"Metascore"`
	ImdbRating   string `json:"imdbRating"`
	ImdbVotes    string `json:"imdbVotes"`
	ImdbID       string `json:"imdbID"`
	Type         string `json:"Type"`
	TotalSeasons string `json:"totalSeasons"`
	Response     string `json:"Response"`
}
