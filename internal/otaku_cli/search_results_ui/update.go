package search_results_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/details_ui"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.init {
		m.init = true

		return m, m.Init()
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			i, ok := m.list.SelectedItem().(*constants.AnimeResult)
			if ok {
				m.selected = i
			}

			m.loading = true

			return m, m.getAnimeDetails

		case key.Matches(msg, m.keys.Return):
			return m.ParentModel, nil

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}

	case constants.DetailMsg:
		selected := m.selected

		m.loading = false
		m.switched = true

		return details_ui.NewUI(m, msg.Data, selected), nil

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}
