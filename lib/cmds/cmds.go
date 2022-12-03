package cmds

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func Wait(duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(duration)

		return constants.WaitMsg{}
	}
}

func GetAnimeStreamingUrls(animeId string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf(constants.ApiUrl+"/vidcdn/watch/%s", animeId)

		resp, status, err := network.ProcessGet(context.Background(), url, nil)
		if err != nil {
			return constants.ErrMsg{Err: err}
		}

		if status != 200 {
			return constants.ErrMsg{Err: fmt.Errorf("server returned %d", status)}
		}

		data := new(constants.StreamData)

		if err := json.Unmarshal(resp, data); err != nil {
			return constants.ErrMsg{Err: err}
		}

		return constants.StreamResultData{Data: data}
	}
}

func SearchAnime(query string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf(constants.ApiUrl+"/search?keyw=%s", query)

		resp, status, err := network.ProcessGet(context.Background(), url, nil)
		if err != nil {
			return constants.ErrMsg{Err: err}
		}

		if status != 200 {
			return constants.ErrMsg{Err: fmt.Errorf("server returned %d", status)}
		}

		data := new([]*styles.AnimeResult)

		if err := json.Unmarshal(resp, data); err != nil {
			return constants.ErrMsg{Err: err}
		}

		return constants.ResultMsg{Data: *data}
	}
}

func GetAnimeDetails(animeId string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf(constants.ApiUrl+"/anime-details/%s", animeId)

		resp, status, err := network.ProcessGet(context.Background(), url, nil)
		if err != nil {
			return constants.ErrMsg{Err: err}
		}

		if status != 200 {
			return constants.ErrMsg{Err: fmt.Errorf("server returned %d", status)}
		}

		data := new(styles.AnimeDetails)

		if err := json.Unmarshal(resp, data); err != nil {
			return constants.ErrMsg{Err: err}
		}

		data.EpisodesList = constants.ReverseSlice(data.EpisodesList)

		return constants.DetailMsg{Data: data}
	}
}
