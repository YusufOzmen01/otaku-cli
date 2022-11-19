package details_ui

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/episode_ui"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/episodes_ui"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.EpisodeList):
			ui := episodes_ui.NewUI(m.EpisodesList, m.AnimeResult)

			return constants.SwitchUI(m, ui, ui.UUID)

		case key.Matches(msg, m.keys.GoBack):
			return constants.ReturnUI(m.UUID)

		case key.Matches(msg, m.keys.Watch):
			index := len(m.EpisodesList) - 1

			a, err := database.GetAnimeProgress(m.AnimeId)
			if err == nil {
				for i, ep := range m.EpisodesList {
					if ep.EpisodeNum == a.LastWatchedEpisode {
						index = i
					}
				}
			}

			ui := episode_ui.NewUI(m.UUID, m.EpisodesList, index, m.AnimeResult)

			anime := &database.Anime{
				ID:                 m.AnimeId,
				Name:               m.AnimeDetails.AnimeTitle,
				LastWatchedEpisode: m.EpisodesList[index].EpisodeNum,
			}

			if err := database.WatchAnime(anime); err != nil {
				panic(err)
			}

			return constants.SwitchUI(m, ui, ui.UUID)
		}

	case constants.StreamResultData:
		err := exec.Command("vlc", msg.Data.Sources[0].File).Start()
		if err != nil {
			return m, tea.Quit
		}

		return m, tea.Quit

	case constants.ErrMsg:
		fmt.Println(msg.Err)

		return m, tea.Quit
	}

	return m, nil
}
