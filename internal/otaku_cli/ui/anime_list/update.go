package anime_list

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.GoBack):
			return constants.ReturnUI(m.UUID)
		}
	}

	return m, nil
}
