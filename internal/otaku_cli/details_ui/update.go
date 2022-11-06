package details_ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Watch):
			exec.Command("explorer", m.AnimeResult.AnimeUrl).Start()

			return m, tea.Quit

		case key.Matches(msg, m.keys.Cancel):
			return m.ParentModel, nil

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		}

	}

	return m, nil
}