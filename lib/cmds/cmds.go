package cmds

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func Wait(duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(duration)

		return constants.WaitMsg{}
	}
}

func GetAnimeStreamingUrls(animeId string) tea.Cmd {
	return func() tea.Msg {
		data, err := anime.GetAnimeStreamingUrls(animeId)
		if err != nil {
			return constants.ErrMsg{Err: err}
		}

		return constants.StreamingUrlsMsg{Data: data}
	}
}

func SearchAnime(query string) tea.Cmd {
	return func() tea.Msg {
		data, err := anime.SearchAnime(query)
		if err != nil {
			return constants.ErrMsg{Err: err}
		}

		return anime.ResultMsg{Data: data}
	}
}

func GetAnimeDetails(animeId string) tea.Cmd {
	return func() tea.Msg {
		data, err := anime.GetAnimeDetails(animeId)
		if err != nil {
			return constants.ErrMsg{Err: err}
		}

		return anime.DetailsMsg{Data: data}
	}
}
