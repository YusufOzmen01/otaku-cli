package anime

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
)

const (
	EnimeApiUrl = "https://api.enime.moe"
)

func SearchAnime(query string) ([]*Result, error) {
	url := fmt.Sprintf(EnimeApiUrl+"/search/%s", query)

	resp, status, err := network.ProcessGet(context.Background(), url, nil)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("server returned %d", status)
	}

	data := new(SearchResult)

	if err := json.Unmarshal(resp, data); err != nil {
		return nil, err
	}

	return data.Data, nil
}

func GetAnimeDetails(animeId string) (*Details, error) {
	url := fmt.Sprintf(EnimeApiUrl+"/anime/%s", animeId)

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

	return data, nil
}

func GetAnimeStreamingUrls(sourceId string) (*constants.StreamData, error) {
	url := fmt.Sprintf(EnimeApiUrl+"/source/%s", sourceId)

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
