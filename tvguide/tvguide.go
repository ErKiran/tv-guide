package tvguide

import (
	"fmt"

	"tv-guide/utils"
)

type TvWish struct {
	client *utils.Client
}

const (
	Upcoming     = "IN/Channels/UpcomingJson/%d"
	CategoryList = "IN/Channels/Category/List"
	ChannelJSON  = "IN/Channels/Categories/ChannelsJson/%d"
)

func NewGuideAPI() (*TvWish, error) {
	client := utils.NewClient(nil, "https://tvwish.com/")

	newTvWiz := &TvWish{
		client: client,
	}
	return newTvWiz, nil
}

func (t TvWish) buildUpcomingSlug(channelId int) string {
	return fmt.Sprintf(Upcoming, channelId)
}

func (t TvWish) buildChannelJSON(categoryId int) string {
	return fmt.Sprintf(ChannelJSON, categoryId)
}

func (t TvWish) buildCategoryList() string {
	return CategoryList
}
