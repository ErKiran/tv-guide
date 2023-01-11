package tvguide

import (
	"context"
	"fmt"
	"net/http"
)

type Channel struct {
	ID             int    `json:"id"`
	Channelname    string `json:"channelname"`
	Radio          bool   `json:"radio"`
	Language       string `json:"language"`
	Resolution     string `json:"resolution"`
	Thumbnail      string `json:"thumbnail"`
	Thumbnaillarge string `json:"thumbnaillarge"`
	Modifytime     int64  `json:"modifytime"`
	Type           string `json:"type"`
	Popularity     int    `json:"popularity"`
	Inprogress     bool   `json:"inprogress"`
	Channelsources []struct {
		ID         int    `json:"id"`
		Channelid  int    `json:"channelid"`
		Source     string `json:"source"`
		Paid       bool   `json:"paid"`
		Cost       int    `json:"cost"`
		Position   int    `json:"position"`
		Modifytime int64  `json:"modifytime"`
	} `json:"channelsources"`
}

type Results struct {
	ID             int     `json:"id"`
	Providername   string  `json:"providername"`
	Channelno      int     `json:"channelno"`
	Hidden         bool    `json:"hidden"`
	Modifytime     int64   `json:"modifytime"`
	Defaultchannel bool    `json:"defaultchannel"`
	Channel        Channel `json:"channel"`
}

type ChannelResponse struct {
	Results []Results `json:"results"`
	Curtime int64     `json:"curtime"`
}

type ChannelFilter struct {
	Type     string
	Language string
}

type channelMap map[string]int

func (t TvWiz) GetChannels(filter ChannelFilter) (channelMap, error) {
	url := t.buildChannelSlug(Channels)

	req, err := t.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error in new request", err)
		return nil, err
	}

	res := &ChannelResponse{}

	if _, err := t.client.Do(context.Background(), req, res); err != nil {
		fmt.Println("error in do", err)
		return nil, err
	}

	chanz := make(map[string]int)

	for _, channel := range res.Results {
		if filter.Type != "" {
			if channel.Channel.Type == filter.Type && channel.Channel.Resolution == "720p" && channel.Channel.Language == filter.Language {
				chanz[channel.Channel.Channelname] = channel.Channel.ID
			}
		}
	}

	return chanz, nil
}
