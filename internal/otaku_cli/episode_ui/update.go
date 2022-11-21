package episode_ui

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
	"strconv"
)

func (m UI) NextEpisode() (tea.Model, tea.Cmd) {
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
}

func (m UI) PreviousEpisode() (tea.Model, tea.Cmd) {
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
}

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.init {
		m.init = true
		m.episodeLoading = true

		return m, m.getAnimeStreamingURL
	}

	switch msg := msg.(type) {
	case constants.ErrMsg:
		fmt.Println(msg.Err)

		return m, tea.Quit

	case VLCMsg:

		pos, err := strconv.Atoi(msg.Data.Time)
		if err != nil {
			panic(err)
		}

		length, err := strconv.Atoi(msg.Data.Length)
		if err != nil {
			panic(err)
		}

		if m.currentVLCData != nil {
			posM, err := strconv.Atoi(m.currentVLCData.Time)
			if err != nil {
				panic(err)
			}

			if posM > 0 && pos == 0 && m.currentVLCData.Information.Text == msg.Data.Information.Text {
				return constants.ReturnUI(m.parentUUID)
			}
		}

		if pos+1 == length {
			return m.NextEpisode()
		}

		anime := &database.Anime{
			ID:                 m.details.AnimeId,
			Name:               m.details.AnimeTitle,
			LastWatchedEpisode: m.episodes[m.currentEpisodeIndex].EpisodeNum,
			Position:           pos,
		}

		if err := database.WatchAnime(anime); err != nil {
			panic(err)
		}

		m.currentVLCData = msg.Data

		return m, m.vlcUpdate

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			constants.KillProcessByNameWindows("vlc.exe")

			return m, tea.Quit

		case key.Matches(msg, m.keys.Next):
			return m.NextEpisode()

		case key.Matches(msg, m.keys.Previous):
			return m.PreviousEpisode()

		case key.Matches(msg, m.keys.GoBack):
			constants.KillProcessByNameWindows("vlc.exe")

			return constants.ReturnUI(m.parentUUID)
		}

	case constants.StreamResultData:
		constants.KillProcessByNameWindows("vlc.exe")

		pos := ""

		anime, err := database.GetAnimeProgress(m.details.AnimeId)
		if err == nil {
			pos = "--start-time=" + strconv.Itoa(anime.Position)
		}

		err = exec.Command("vlc", msg.Data.Sources[0].File, "--intf", "qt", "--extraintf", "http", "--http-password=amongus_is_funny", "--http-port=58000", pos).Start()
		if err != nil {
			return m, tea.Quit
		}

		m.episodeLoading = false

		return m, m.vlcUpdate
	}

	return m, nil
}
