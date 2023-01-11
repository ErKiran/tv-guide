package tvguide

import (
	"context"
	"fmt"
	"net/http"
)

type Program struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Releasedate   int64  `json:"releasedate"`
	Releaseyear   int    `json:"releaseyear"`
	Episodeno     int    `json:"episodeno"`
	Seasonno      int    `json:"seasonno"`
	Language      string `json:"language"`
	Length        int    `json:"length"`
	Thumbnail     string `json:"thumbnail"`
	Thumbnailw92  string `json:"thumbnailw92"`
	Thumbnailw154 string `json:"thumbnailw154"`
	Thumbnailw780 string `json:"thumbnailw780"`
	Poster        string `json:"poster"`
	Posterw92     string `json:"posterw92"`
	Posterw154    string `json:"posterw154"`
	Genre         string `json:"genre"`
	Type          string `json:"type"`
	Imdbid        string `json:"imdbid"`
	Themoviedbid  int    `json:"themoviedbid"`
	Summary       string `json:"summary"`
	Tagline       string `json:"tagline"`
	Rating        int    `json:"rating"`
	Mainrating    int    `json:"mainrating"`
	Popularity    int    `json:"popularity"`
	Modifytime    int64  `json:"modifytime"`
	Castloaded    bool   `json:"castLoaded"`
	Sourcesloaded bool   `json:"sourcesLoaded"`
	Peopleloaded  bool   `json:"peopleLoaded"`
	Hastrailers   bool   `json:"hasTrailers"`
	Similarloaded bool   `json:"similarLoaded"`
}

type NextProgram struct {
	ID            int     `json:"id"`
	Channelid     int     `json:"channelid"`
	Program       Program `json:"program"`
	Programid     int     `json:"programid"`
	Starttime     int64   `json:"starttime"`
	Endtime       int64   `json:"endtime"`
	Inprogress    bool    `json:"inprogress"`
	Newontv       bool    `json:"newontv"`
	Modifytime    int64   `json:"modifytime"`
	Repeatsloaded bool    `json:"repeatsLoaded"`
}

type ProgramResult struct {
	ID            int         `json:"id"`
	Channelid     int         `json:"channelid"`
	Program       Program     `json:"program,omitempty"`
	Programid     int         `json:"programid"`
	Starttime     int64       `json:"starttime"`
	Endtime       int64       `json:"endtime"`
	Inprogress    bool        `json:"inprogress"`
	Newontv       bool        `json:"newontv"`
	Modifytime    int64       `json:"modifytime"`
	Nextprogram   NextProgram `json:"nextProgram"`
	Repeatsloaded bool        `json:"repeatsLoaded"`
}

type ProgramResponse struct {
	Results []ProgramResult `json:"results"`
}

func (t TvWiz) GetProgramme(channels string) (*ProgramResponse, error) {
	url := t.buildProgrammeSlug(channels)

	req, err := t.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error in new request", err)
		return nil, err
	}

	res := &ProgramResponse{}

	if _, err := t.client.Do(context.Background(), req, res); err != nil {
		fmt.Println("error in do", err)
		return nil, err
	}

	return res, nil
}
