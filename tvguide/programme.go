package tvguide

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ProgrammeData struct {
	Lastmodified    time.Time   `json:"lastModified"`
	Lastmodifiedutc string      `json:"lastModifiedUtc"`
	Timezone        interface{} `json:"timeZone"`
	Data            []Data      `json:"data"`
}

type Data struct {
	Pi   int         `json:"pi"`
	Ci   interface{} `json:"ci"`
	Tzo  interface{} `json:"tzo"`
	Tzn  interface{} `json:"tzn"`
	Cti  interface{} `json:"cti"`
	Spn  string      `json:"spn"`
	Pn   string      `json:"pn"`
	Pne  string      `json:"pne"`
	Cn   interface{} `json:"cn"`
	Cc   interface{} `json:"cc"`
	Cne  interface{} `json:"cne"`
	Cl   interface{} `json:"cl"`
	Pu   string      `json:"pu"`
	Rd   string      `json:"rd"`
	Sp   string      `json:"sp"`
	Sdtl interface{} `json:"sdtl"`
	Sdtu string      `json:"sdtu"`
	Stf  string      `json:"stf"`
	Ed   string      `json:"ed"`
	En   interface{} `json:"en"`
	Cs   interface{} `json:"cs"`
	Si   int         `json:"si"`
	Day  interface{} `json:"day"`
}

func (tv TvWish) GetProgrammeByChannelId(channelId int) (*ProgrammeData, error) {
	url := tv.buildUpcomingSlug(channelId)

	req, err := tv.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error in new request", err)
		return nil, err
	}

	res := &ProgrammeData{}

	if _, err := tv.client.Do(context.Background(), req, res); err != nil {
		fmt.Println("error in do", err)
		return nil, err
	}

	return res, nil
}
