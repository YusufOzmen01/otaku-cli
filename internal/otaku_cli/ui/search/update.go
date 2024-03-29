package search

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/ui/search_results"
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/YusufOzmen01/otaku-cli/lib/cmds"
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
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Enter):
			m.loading = true

			return m, cmds.SearchAnime(m.textInput.Value())

		case key.Matches(msg, m.keys.GoBack):
			return constants.ReturnUI(m.UUID)

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

	case anime.ResultMsg:
		m.httpErr = nil

		if len(msg.Data) == 0 {
			m.nothing = true

			return m, cmds.Wait(time.Second * 2)
		}

		m.searchDone = true

		return m, updateAnimes(msg.Data)

	case anime.SearchDoneMsg:
		m.switched = true
		m.loading = false
		m.searchDone = false

		ui := search_results.NewUI(msg.Data, m.textInput.Value())
		m.textInput.SetValue("")

		return constants.SwitchUI(m, ui, ui.UUID)

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
