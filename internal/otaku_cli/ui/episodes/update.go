package episodes

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/ui/episode"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.init {
		m.init = true

		items := make([]list.Item, 0)

		for _, e := range m.episodes {
			items = append(items, list.Item(e))
		}

		m.list = list.New(items, styles.AnimeEpisodesDelegate{AnimeID: m.details.Id}, 0, 20)
		m.list.Title = fmt.Sprintf("%s\n%s",
			styles.TitleStyle.Render("Episode List"),
			fmt.Sprintf("Found %d episodes for \"%s\"\nNote: Episode length will be 0 until you watch that episode.\nIt's a problem related to api.", len(m.episodes), m.details.AnimeTitle))
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
				ui := episode.NewUI(uuid.UUID{}, m.episodes, m.list.Index(), m.details, true)

				anime := &database.Anime{
					ID:   m.details.Id,
					Name: m.details.AnimeTitle.Romaji,
					CurrentEpisode: &database.Episode{
						Number: m.list.Index(),
					},
					MaxEpisodes: len(m.episodes),
				}

				if err := database.UpdateAnimeTracking(anime); err != nil {
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
