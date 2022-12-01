package search

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
	tea "github.com/charmbracelet/bubbletea"
)

func (m UI) searchAnime() tea.Msg {
	url := fmt.Sprintf(constants.ApiUrl+"/search?keyw=%s", m.textInput.Value())

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
