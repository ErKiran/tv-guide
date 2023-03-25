package tvguide

import (
	"context"
	"fmt"
	"net/http"
)

type Category struct {
	I  int         `json:"i"`
	N  string      `json:"n"`
	Ne interface{} `json:"ne"`
	So int         `json:"so"`
}

func (tv TvWish) GetCategories() ([]Category, error) {
	url := tv.buildCategoryList()

	req, err := tv.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error in new request", err)
		return nil, err
	}

	res := []Category{}

	if _, err := tv.client.Do(context.Background(), req, &res); err != nil {
		fmt.Println("error in do", err)
		return nil, err
	}

	return res, nil
}
