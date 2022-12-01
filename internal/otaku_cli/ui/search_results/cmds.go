package search_results

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
	tea "github.com/charmbracelet/bubbletea"
)

func (m UI) getAnimeDetails() tea.Msg {
	url := fmt.Sprintf(constants.ApiUrl+"/anime-details/%s", m.selected.AnimeId)

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
