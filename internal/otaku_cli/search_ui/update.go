package search_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/search_results_ui"
	"github.com/YusufOzmen01/otaku-cli/lib/generic_cmds"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.switched || !m.init {
		m.switched = false
		m.init = true

		return m, m.spinner.Tick
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			m.loading = true

			return m, m.searchAnime

		case key.Matches(msg, m.keys.Quit):
			return m.ParentModel, nil

		default:
			if !m.loading {
				var cmd tea.Cmd
				m.textInput, cmd = m.textInput.Update(msg)

				return m, cmd
			}
		}

	case constants.ErrMsg:
		m.httpErr = msg.Err
		return m, tea.Quit

	case constants.ResultMsg:
		m.httpErr = nil
		m.loading = false

		if len(msg.Data) == 0 {
			m.nothing = true

			return m, generic_cmds.Wait(time.Second * 2)
		}

		m.switched = true
		m.textInput.SetValue("")

		return search_results_ui.NewUI(m, msg.Data), nil

	case constants.WaitMsg:
		m.nothing = false
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)

		return m, cmd
	}

	return m, nil
}
