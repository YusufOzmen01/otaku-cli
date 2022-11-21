package episode_ui

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
	tea "github.com/charmbracelet/bubbletea"
)

const ApiUrl = "https://gogoanime.consumet.org"

func (m UI) getAnimeStreamingURL() tea.Msg {
	url := fmt.Sprintf(ApiUrl+"/vidcdn/watch/%s", m.episodes[m.currentEpisodeIndex].EpisodeId)

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

func (m UI) vlcUpdate() tea.Msg {
	body, status, err := network.ProcessGet(context.Background(), "http://localhost:58000/requests/status.xml", map[string]string{"Authorization": "Basic OmFtb25ndXNfaXNfZnVubnk="})
	if err != nil {
		return m.vlcUpdate()
	}

	if status != 200 {
		return m.vlcUpdate()
	}

	data := new(Root)

	if err := xml.Unmarshal(body, data); err != nil {
		return m.vlcUpdate()
	}

	return VLCMsg{Data: data}
}
