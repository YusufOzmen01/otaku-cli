package episodes_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.init {
		m.init = true

		items := make([]list.Item, 0)

		for _, result := range m.episodes {
			items = append(items, list.Item(result))
		}

		m.list = list.New(items, constants.AnimeResultDelegate{}, 0, 20)
		m.list.Title = titleStyle.Render("Episode List")
		m.list.SetShowStatusBar(true)
		m.list.SetFilteringEnabled(true)
		m.list.Styles.Title = lipgloss.NewStyle().MarginLeft(0)
		m.list.Styles.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(2)
		m.list.Styles.HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(2).PaddingBottom(0)

		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			_, ok := m.list.SelectedItem().(*constants.Episode)
			if ok {
				return m, tea.Quit
			}

			return m, nil

		case key.Matches(msg, m.keys.Return):
			return constants.ReturnUI(m.UUID)

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		}

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}
