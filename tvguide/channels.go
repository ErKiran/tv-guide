package tvguide

import (
	"context"
	"fmt"
	"net/http"
)

type Channel struct {
	I  int         `json:"i"`
	N  string      `json:"n"`
	D  interface{} `json:"d"`
	Pu string      `json:"pu"`
	Ne interface{} `json:"ne"`
	Cc interface{} `json:"cc"`
	So interface{} `json:"so"`
}

func (tv TvWish) GetChannelsByCategoryID(categoryID int) ([]Channel, error) {
	url := tv.buildChannelJSON(categoryID)

	req, err := tv.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error in new request", err)
		return nil, err
	}

	res := []Channel{}

	if _, err := tv.client.Do(context.Background(), req, &res); err != nil {
		fmt.Println("error in do", err)
		return nil, err
	}

	return res, nil
}
