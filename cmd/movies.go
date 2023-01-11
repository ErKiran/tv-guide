package cmd

import (
	"fmt"
	"strconv"
	"time"

	"tv-guide/omdb"
	"tv-guide/tvguide"

	"github.com/brettski/go-termtables"
	"github.com/spf13/cobra"
)

var MovieCmd = &cobra.Command{
	Use:   "movies",
	Short: "List Current and Upcoming Movies in Hindi Channel",
	Long:  `For Movies`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		MovieCommand()
	},
}

var interestedChannels = map[string]string{}

func init() {
	interestedChannels = ChannelMap()
}

func ChannelMap() map[string]string {
	interestedChannels = map[string]string{
		"Star Gold Select HD": "https://tvwiz.in/channel/Star-Gold-Select-HD",
		"&pictures HD":        "https://tvwiz.in/channel/pictures-HD",
		"Cineplex HD":         "https://tvwiz.in/channel/Cineplex-HD",
		"Sony MAX HD":         "https://tvwiz.in/channel/Sony-MAX-HD",
		"Star Gold HD":        "https://tvwiz.in/channel/Star-Gold-HD",
		"Zee Cinema HD":       "https://tvwiz.in/channel/Zee-Cinema-HD",
	}
	return interestedChannels
}

type MovieInfo struct {
	Channel     int
	NowShowing  string
	ReleaseYear int
	IMDBID      string
	Genre       string
	Started     string
	EndsOn      string
	Eclipsed    string
	NextChange  string
	StartOn     string
	NextGenre   string
	NextIMDBID  string
}

func MovieCommand() {
	table := termtables.CreateTable()

	guide, err := tvguide.NewGuideAPI()
	if err != nil {
		fmt.Println(err)
	}
	channels, err := guide.GetChannels(tvguide.ChannelFilter{
		Type:     "Movies",
		Language: "Hindi",
	})
	if err != nil {
		fmt.Println(err)
	}

	channelIds := ""

	newChanMap := map[int]string{}

	for name, id := range channels {
		newChanMap[id] = name
		if _, ok := interestedChannels[name]; ok {
			channelIds += fmt.Sprintf("%s,", strconv.Itoa(id))
		}
	}

	programme, err := guide.GetProgramme(channelIds)
	if err != nil {
		fmt.Println(err)
	}

	var mInfos []MovieInfo
	for _, pInfo := range programme.Results {
		mInfos = append(mInfos, MovieInfo{
			Channel:     pInfo.Channelid,
			Started:     HumanizeTime(pInfo.Starttime),
			EndsOn:      HumanizeTime(pInfo.Endtime),
			NowShowing:  fmt.Sprintf("%s (%d)", pInfo.Program.Name, pInfo.Program.Releaseyear),
			ReleaseYear: pInfo.Program.Releaseyear,
			IMDBID:      GetImdbRating(pInfo.Program.Imdbid),
			Genre:       pInfo.Program.Genre,
			NextChange:  fmt.Sprintf("%s (%d)", pInfo.Nextprogram.Program.Name, pInfo.Nextprogram.Program.Releaseyear),
			NextGenre:   pInfo.Nextprogram.Program.Genre,
			NextIMDBID:  GetImdbRating(pInfo.Nextprogram.Program.Imdbid),
		})
	}

	table.AddHeaders("Channel", "Now Showing", "Genre", "IMDB", "Start", "End", "Next Change", "Genre", "IMDB")
	for _, mInfo := range mInfos {
		table.AddRow(newChanMap[mInfo.Channel], mInfo.NowShowing, mInfo.Genre, mInfo.IMDBID, mInfo.Started, mInfo.EndsOn, mInfo.NextChange, mInfo.NextGenre, mInfo.NextIMDBID)
	}

	table.AddTitle("Movie Channels")

	fmt.Println(table.Render())
}

func HumanizeTime(timez int64) string {
	return time.UnixMilli(timez).In(time.Local).Format("15:04:05")
}

func GetImdbRating(id string) string {
	if id == "" {
		return "N/A"
	}
	omdb, err := omdb.NewOmdbAPI()
	if err != nil {
		fmt.Println(err)
	}
	info, err := omdb.GetMovieInfo(id)
	if err != nil {
		return "N/A"
	}

	return fmt.Sprintf("%s (%s)", info.Imdbrating, info.Imdbvotes)
}
