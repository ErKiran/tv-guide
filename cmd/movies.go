package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"tv-guide/omdb"
	"tv-guide/tvguide"

	"github.com/brettski/go-termtables"
	"github.com/spf13/cobra"
)

var layout = "2006-01-02T15:04:05"

var MovieCmd = &cobra.Command{
	Use:   "movies",
	Short: "List Current and Upcoming Movies in Hindi Channel",
	Long:  `For Movies`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		MovieCommand()
	},
}

type MovieInfo struct {
	Channel     string
	NowShowing  string
	ReleaseYear string
	IMDBID      string
	IMDBRating  string
	Genre       []string
	Started     string
	EndsOn      string
	Eclipsed    string
	NextChange  string
	StartOn     string
	NextGenre   []string
	NextIMDBID  string
}

type ProgrammeInfo struct {
	ChannelName string
	MovieName   string
	StartTime   string
	FullTime    time.Time
}

func MovieCommand() {
	table := termtables.CreateTable()

	guide, err := tvguide.NewGuideAPI()
	if err != nil {
		fmt.Println(err)
	}
	cat, err := guide.GetCategories()
	if err != nil {
		fmt.Println(err)
	}

	var movieCat int

	for _, ca := range cat {
		if ca.N == "Hindi Movies" {
			movieCat = ca.I
		}
	}
	channelMap := make(map[string]int)

	chans, err := guide.GetChannelsByCategoryID(movieCat)
	if err != nil {
		fmt.Println("err", err)
	}
	for _, channel := range chans {
		if strings.Contains(channel.N, "HD") || channel.N == "Star Gold Select" || channel.N == "Star Gold" {
			channelMap[channel.N] = channel.I
		}
	}

	var pInfos []ProgrammeInfo
	for k, v := range channelMap {
		prg, err := guide.GetProgrammeByChannelId(v)
		if err != nil {
			fmt.Println("err", err)
		}
		for _, p := range prg.Data {
			tm, err := time.Parse(layout, p.Sdtu)
			if err != nil {
				fmt.Println("err", err)
			}
			hour, min, sec := tm.In(time.Local).Clock()
			pInfos = append(pInfos, ProgrammeInfo{
				ChannelName: k,
				MovieName:   p.Spn,
				StartTime:   fmt.Sprintf("%d:%02d:%02d", hour, min, sec),
				FullTime:    tm.In(time.Local),
			})
		}
	}

	var mInfos []MovieInfo
	movieInfo := make(map[string]omdb.MovieInfo)
	for _, pInfo := range pInfos[0:1] {
		mdata, err := GetMovieDetails(pInfo.MovieName)
		if err != nil {
			fmt.Println("err", err)
		}
		movieInfo[pInfo.MovieName] = *mdata
		mInfos = append(mInfos, MovieInfo{
			Channel:    pInfo.ChannelName,
			Started:    pInfo.StartTime,
			NowShowing: fmt.Sprintf("%s (%s)", pInfo.MovieName, mdata.Year),
			// ReleaseYear: pInfo.Program.Releaseyear,
			IMDBRating: func() string {
				info := GetMovieInfo(movieInfo, pInfo.MovieName)
				return info.IMDBRating
			}(),
			Genre: func() []string {
				info := GetMovieInfo(movieInfo, pInfo.MovieName)
				return info.Genre
			}(),
			// NextChange:  fmt.Sprintf("%s (%d)", pInfo.Nextprogram.Program.Name, pInfo.Nextprogram.Program.Releaseyear),
			// NextGenre:   GenreTrim(pInfo.Nextprogram.Program.Genre),
			// NextIMDBID:  GetImdbRating(pInfo.Nextprogram.Program.Imdbid),
		})
	}

	table.AddHeaders("Channel", "Now Showing", "Genre", "IMDB", "Start", "End", "Next Change", "Genre", "IMDB")
	for _, mInfo := range mInfos {
		table.AddRow(mInfo.Channel, mInfo.NowShowing, mInfo.Genre, mInfo.IMDBRating, mInfo.Started, mInfo.EndsOn, mInfo.NextChange, mInfo.NextGenre, mInfo.NextIMDBID)
	}

	table.AddTitle("Movie Channels")

	fmt.Println(table.Render())
}

func GenreTrim(genre string) []string {
	genres := strings.Split(genre, ",")
	if len(genres) > 3 {
		return genres[:3]
	}
	return genres
}

func GetMovieDetails(name string) (*omdb.MovieInfo, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	omdb, err := omdb.NewOmdbAPI()
	if err != nil {
		return nil, err
	}
	info, err := omdb.GetMovieInfo(name)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func GetMovieInfo(data map[string]omdb.MovieInfo, name string) MovieInfo {
	movie, ok := data[name]
	if !ok {
		return MovieInfo{}
	}

	return MovieInfo{
		IMDBRating:  fmt.Sprintf("%s (%s)", movie.Imdbrating, movie.Imdbvotes),
		ReleaseYear: movie.Year,
		Genre:       GenreTrim(movie.Genre),
	}
}
