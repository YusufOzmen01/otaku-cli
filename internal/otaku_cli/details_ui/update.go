package details_ui

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Watch):
			//exec.Command("explorer", m.AnimeResult.AnimeUrl).Start()

			return m, m.getAnimeStreamingURL

		case key.Matches(msg, m.keys.GoBack):
			return m.ParentModel, nil

		}

	case constants.StreamResultData:
		err := exec.Command("vlc", msg.Data.Sources[0].File).Start()
		if err != nil {
			return m, tea.Quit
		}

		return m, tea.Quit

	case constants.ErrMsg:
		fmt.Println(msg.Err)

		return m, tea.Quit
	}

	return m, nil
}
