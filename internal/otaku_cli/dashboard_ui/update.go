package dashboard_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/search_ui"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Search):
			ui := search_ui.NewUI()

			return constants.SwitchUI(m, ui, ui.UUID)

		case key.Matches(msg, m.keys.LastWatched):
			return m, tea.Quit
		}

	}

	return m, nil
}
