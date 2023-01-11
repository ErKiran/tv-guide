package omdb

import (
	"context"
	"fmt"
	"net/http"
)

type MovieInfo struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	Imdbrating string `json:"imdbRating"`
	Imdbvotes  string `json:"imdbVotes"`
	Imdbid     string `json:"imdbID"`
	Type       string `json:"Type"`
	Dvd        string `json:"DVD"`
	Boxoffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

func (o OMDB) GetMovieInfo(id string) (*MovieInfo, error) {
	url := o.buildInfoSlug(Info, id, "changeme")
	req, err := o.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error in new request", err)
		return nil, err
	}

	res := &MovieInfo{}

	if _, err := o.client.Do(context.Background(), req, res); err != nil {
		fmt.Println("error in do", err)
		return nil, err
	}
	return res, err
}
