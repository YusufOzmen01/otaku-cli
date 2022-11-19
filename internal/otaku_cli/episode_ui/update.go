package episode_ui

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.init {
		m.init = true
		m.episodeLoading = true

		return m, m.getAnimeStreamingURL
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Next):
			if m.episodeLoading {
				return m, nil
			}

			if m.currentEpisodeIndex-1 < 0 {
				return m, nil
			}

			ui := NewUI(m.parentUUID, m.episodes, m.currentEpisodeIndex-1, m.details)

			anime := &database.Anime{
				ID:                 m.details.AnimeId,
				Name:               m.details.AnimeTitle,
				LastWatchedEpisode: m.episodes[m.currentEpisodeIndex-1].EpisodeNum,
			}

			if err := database.WatchAnime(anime); err != nil {
				panic(err)
			}

			return constants.SwitchUI(m, ui, ui.UUID)

		case key.Matches(msg, m.keys.Previous):
			if m.episodeLoading {
				return m, nil
			}

			if m.currentEpisodeIndex+1 == len(m.episodes) {
				return m, nil
			}

			ui := NewUI(m.parentUUID, m.episodes, m.currentEpisodeIndex+1, m.details)

			anime := &database.Anime{
				ID:                 m.details.AnimeId,
				Name:               m.details.AnimeTitle,
				LastWatchedEpisode: m.episodes[m.currentEpisodeIndex+1].EpisodeNum,
			}

			if err := database.WatchAnime(anime); err != nil {
				panic(err)
			}

			return constants.SwitchUI(m, ui, ui.UUID)

		case key.Matches(msg, m.keys.GoBack):
			return constants.ReturnUI(m.parentUUID)
		}

	case constants.StreamResultData:
		constants.KillProcessByNameWindows("vlc.exe")

		err := exec.Command("vlc", msg.Data.Sources[0].File).Start()
		if err != nil {
			return m, tea.Quit
		}

		m.episodeLoading = false

		return m, nil

	case constants.ErrMsg:
		fmt.Println(msg.Err)

		return m, tea.Quit
	}

	return m, nil
}
