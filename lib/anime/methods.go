package anime

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
)

const ApiUrl = "https://gogoanime.consumet.org"

func GetAnimeStreamingUrls(animeId string) (*constants.StreamData, error) {
	url := fmt.Sprintf(ApiUrl+"/vidcdn/watch/%s", animeId)

	resp, status, err := network.ProcessGet(context.Background(), url, nil)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("server returned %d", status)
	}

	data := new(constants.StreamData)

	if err := json.Unmarshal(resp, data); err != nil {
		return nil, err
	}

	return data, nil
}

func SearchAnime(query string) ([]*Result, error) {
	url := fmt.Sprintf(ApiUrl+"/search?keyw=%s", query)

	resp, status, err := network.ProcessGet(context.Background(), url, nil)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("server returned %d", status)
	}

	data := new([]*Result)

	if err := json.Unmarshal(resp, data); err != nil {
		return nil, err
	}

	return *data, nil
}

func GetAnimeDetails(animeId string) (*Details, error) {
	url := fmt.Sprintf(ApiUrl+"/anime-details/%s", animeId)

	resp, status, err := network.ProcessGet(context.Background(), url, nil)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("server returned %d", status)
	}

	data := new(Details)

	if err := json.Unmarshal(resp, data); err != nil {
		return nil, err
	}

	data.EpisodesList = constants.ReverseSlice(data.EpisodesList)

	return data, nil
}
