package tvguide

import (
	"fmt"

	"tv-guide/utils"
)

type TvWiz struct {
	client *utils.Client
}

const (
	Channels  = "channels?provider=other"
	Programme = "v2/currentProgram?channelids=%s&apiKey=useless&append=next"
)

func NewGuideAPI() (*TvWiz, error) {
	client := utils.NewClient(nil, "https://tvwiz.in/api/")

	newTvWiz := &TvWiz{
		client: client,
	}
	return newTvWiz, nil
}

func (t TvWiz) buildChannelSlug(url string) string {
	return url
}

func (t TvWiz) buildProgrammeSlug(channels string) string {
	return fmt.Sprintf(Programme, channels)
}
