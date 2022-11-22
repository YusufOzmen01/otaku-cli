package episodes_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/episode_ui"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.init {
		m.init = true

		items := make([]list.Item, 0)

		for _, episode := range m.episodes {
			items = append(items, list.Item(episode))
		}

		m.list = list.New(items, constants.AnimeEpisodesDelegate{AnimeID: m.details.AnimeId}, 0, 20)
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
		if !m.list.SettingFilter() {
			switch {
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit

			case key.Matches(msg, m.keys.Enter):
				ui := episode_ui.NewUI(m.UUID, m.episodes, m.list.Index(), m.details)

				anime := &database.Anime{
					ID:   m.details.AnimeId,
					Name: m.details.AnimeTitle,
					EpisodeProgress: &database.EpisodeProgress{
						CurrentEpisodeIndex: m.list.Index(),
						MaxEpisodes:         len(m.episodes),
					},
				}

				if err := database.WatchAnime(anime); err != nil {
					panic(err)
				}

				return constants.SwitchUI(m, ui, ui.UUID)

			case key.Matches(msg, m.keys.Return):
				return constants.ReturnUI(m.UUID)

			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}
