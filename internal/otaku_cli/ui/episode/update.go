package episode

import (
	"context"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/lib/cmds"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
	"github.com/YusufOzmen01/otaku-cli/lib/vlc"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"strconv"
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

		time, err := strconv.Atoi(m.currentVLCData.Time)
		if err != nil {
			panic(err)
		}

		pos = time
	}

	ui := NewUI(m.parentUUID, m.episodes, episodeIndex, m.details, false)

	length, err := strconv.Atoi(m.currentVLCData.Length)
	if err != nil {
		panic(err)
	}

	anime := &database.Anime{
		ID:   m.details.Id,
		Name: m.details.AnimeTitle.Romaji,
		CurrentEpisode: &database.Episode{
			Number:   episodeIndex,
			Position: pos,
			Length:   length,
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

	ui := NewUI(m.parentUUID, m.episodes, m.currentEpisodeIndex-1, m.details, false)

	return constants.SwitchUI(m, ui, ui.UUID)
}

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.init {
		m.init = true
		m.episodeLoading = true

		m.source = "gogoanime"

		return m, cmds.GetAnimeStreamingUrls(m.episodes[m.currentEpisodeIndex].Sources[0].Id)
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

		if length > 0 {
			m.receivedData = true

			if pos+2 >= length {
				return m.NextEpisode()
			}
		}

		ep := &database.Episode{
			Number:   m.currentEpisodeIndex,
			Position: pos,
			Length:   length,
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

		m.currentVLCData = msg.Data

		return m, m.vlcUpdate(m.vlc)

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

		args := make([]string, 0)
		episode, err := database.GetEpisodeProgress(m.currentEpisodeIndex, m.details.Id)
		if err == nil {
			args = append(args, fmt.Sprintf("--start-time=%d", episode.Position))
		}

		if len(msg.Data.Subtitle) > 0 {
			args = append(args, fmt.Sprintf("--input-slave=%s", msg.Data.Subtitle))
		}

		v, err := vlc.NewVLC(msg.Data.Url, args)
		if err != nil {
			return m, tea.Quit
		}

		m.vlc = v

		m.episodeLoading = false

		return m, m.vlcUpdate(v)
	}

	return m, nil
}
