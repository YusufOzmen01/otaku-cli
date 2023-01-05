package episode

import (
	"context"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/lib/cmds"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/YusufOzmen01/otaku-cli/lib/mpv"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m UI) NextEpisode() (tea.Model, tea.Cmd) {
	if m.episodeLoading {
		return m, nil
	}

	episodeIndex := m.currentEpisodeIndex + 1
	finished := false
	done := false
	pos := 0

	if episodeIndex == len(m.episodes) {
		episodeIndex--
		done = true
		if m.details.Status == "FINISHED" {
			finished = true
		}

		pos = m.currentProgress.Time
	}

	ui := NewUI(m.parentUUID, m.episodes, episodeIndex, m.details, false)

	m.mpv.Kill()

	anime := &database.Anime{
		ID:   m.details.Id,
		Name: m.details.AnimeTitle.Romaji,
		CurrentEpisode: &database.Episode{
			Number:   episodeIndex,
			Position: pos,
			Length:   m.currentProgress.Length,
		},
		MaxEpisodes: len(m.episodes),
		Finished:    finished,
	}

	if err := database.UpdateAnimeTracking(anime); err != nil {
		panic(err)
	}

	if done {
		return constants.ReturnUI(m.parentUUID)
	}

	return constants.SwitchUI(m, ui, ui.UUID)
}

func (m UI) PreviousEpisode() (tea.Model, tea.Cmd) {
	if m.episodeLoading {
		return m, nil
	}

	if m.currentEpisodeIndex-1 < 0 {
		return m, nil
	}

	if m.mpv != nil {
		m.mpv.Kill()
	}

	ui := NewUI(m.parentUUID, m.episodes, m.currentEpisodeIndex-1, m.details, false)

	return constants.SwitchUI(m, ui, ui.UUID)
}

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.init {
		m.init = true
		m.episodeLoading = true

		m.source = "gogoanime"

		return m, cmds.GetAnimeStreamingUrls(m.episodes[m.currentEpisodeIndex].Sources[1].Id)
	}

	switch msg := msg.(type) {
	case constants.ErrMsg:
		fmt.Println(msg.Err)

		return m, tea.Quit

	case ProgressUpdate:
		m.mpvLoading = msg.Data.Loading

		if m.mpvLoading {
			return m, m.progressUpdate(m.mpv)
		}

		if msg.Data.Time+2 >= msg.Data.Length {
			return m.NextEpisode()
		}

		ep := &database.Episode{
			Number:   m.currentEpisodeIndex,
			Position: msg.Data.Time,
			Length:   msg.Data.Length,
		}

		anime := &database.Anime{
			ID:             m.details.Id,
			Name:           m.details.AnimeTitle.Romaji,
			CurrentEpisode: ep,
			MaxEpisodes:    len(m.episodes),
		}

		if err := database.UpdateAnimeTracking(anime); err != nil {
			panic(err)
		}

		if err := database.UpdateEpisode(ep, m.details.Id); err != nil {
			panic(err)
		}

		m.currentProgress = msg.Data

		return m, m.progressUpdate(m.mpv)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Next):
			return m.NextEpisode()

		case key.Matches(msg, m.keys.Previous):
			return m.PreviousEpisode()

		case key.Matches(msg, m.keys.GoBack):
			if m.mpv != nil {
				m.mpv.Kill()
			}

			return constants.ReturnUI(m.parentUUID)
		}

	case constants.StreamingUrlsMsg:
		_, status, err := network.ProcessGet(context.Background(), msg.Data.Url, nil)
		if err != nil {
			m.source = "zoro.to"

			return m, cmds.GetAnimeStreamingUrls(m.episodes[m.currentEpisodeIndex].Sources[1].Id)
		}

		if status != 200 {
			m.source = "zoro.to"

			return m, cmds.GetAnimeStreamingUrls(m.episodes[m.currentEpisodeIndex].Sources[1].Id)
		}

		time := 0
		episode, err := database.GetEpisodeProgress(m.currentEpisodeIndex, m.details.Id)
		if err == nil {
			time = episode.Position
		}

		m.mpv = mpv.NewMPV(msg.Data.Url, msg.Data.Subtitle, time)
		m.episodeLoading = false

		return m, m.progressUpdate(m.mpv)
	}

	return m, nil
}
