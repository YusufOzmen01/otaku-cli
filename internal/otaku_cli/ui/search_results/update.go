package search_results

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/ui/details"
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/YusufOzmen01/otaku-cli/lib/cmds"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.init {
		m.init = true

		items := make([]list.Item, 0)

		for _, result := range m.Results {
			items = append(items, list.Item(result))
		}

		m.list = list.New(items, styles.AnimeResultDelegate{}, 0, 20)
		m.list.Title = fmt.Sprintf("%s\n%s",
			styles.TitleStyle.Render("Anime Search Result"),
			fmt.Sprintf("Found %d results for \"%s\"", len(m.Results), m.SearchText))
		m.list.SetShowStatusBar(true)
		m.list.SetFilteringEnabled(true)
		m.list.Styles.Title = lipgloss.NewStyle().MarginLeft(0)
		m.list.Styles.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(2)
		m.list.Styles.HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(2).PaddingBottom(0)

		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !m.list.SettingFilter() {
			switch {
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit

			case key.Matches(msg, m.keys.Enter):
				i, ok := m.list.SelectedItem().(*anime.Result)
				if ok {
					m.selected = i
				}

				m.loading = true

				return m, cmds.GetAnimeDetails(i.AnimeId)

			case key.Matches(msg, m.keys.Return):
				return constants.ReturnUI(m.UUID)

			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			}
		}

	case anime.DetailsMsg:
		selected := m.selected

		m.loading = false
		m.switched = true

		ui := details.NewUI(msg.Data, selected)

		return constants.SwitchUI(m, ui, ui.UUID)

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}
