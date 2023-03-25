package omdb

import (
	"fmt"

	"tv-guide/utils"
)

type OMDB struct {
	client *utils.Client
}

const (
	Info = "?t=%s&apikey=%s"
)

func NewOmdbAPI() (*OMDB, error) {
	client := utils.NewClient(nil, "http://www.omdbapi.com/")

	newOMDB := &OMDB{
		client: client,
	}
	return newOMDB, nil
}

func (t OMDB) buildInfoSlug(url, id, api string) string {
	return fmt.Sprintf(url, id, api)
}
