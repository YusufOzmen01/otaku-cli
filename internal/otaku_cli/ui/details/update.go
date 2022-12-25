package details

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/ui/episode"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/ui/episodes"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
)

func (m UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width
		if m.progress.Width > 80 {
			m.progress.Width = 80
		}
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Finished):
			anime := &database.Anime{
				ID:   m.Details.Id,
				Name: m.Details.Title.Romaji,
				CurrentEpisode: &database.Episode{
					Number:   len(m.Details.Episodes) - 1,
					Position: 1,
					Length:   1,
				},
				MaxEpisodes: len(m.Episodes),
				Finished:    m.Details.Status == "Completed",
			}

			if err := database.UpdateAnimeTracking(anime); err != nil {
				panic(err)
			}

			return m, nil

		case key.Matches(msg, m.keys.Reset):
			_, err := database.GetAnimeProgress(m.Details.Id)
			if err == nil {
				if err := database.ResetAnimeProgress(m.Details.Id, len(m.Details.Episodes)); err != nil {
					panic(err)
				}
			}

			return m, nil
		case key.Matches(msg, m.keys.EpisodeList):
			ui := episodes.NewUI(m.Episodes, m.Result)

			return constants.SwitchUI(m, ui, ui.UUID)

		case key.Matches(msg, m.keys.GoBack):
			return constants.ReturnUI(m.UUID)

		case key.Matches(msg, m.keys.Watch):
			index, pos := 0, 0

			a, err := database.GetAnimeProgress(m.Details.Id)
			if err == nil {
				index = a.CurrentEpisode.Number
				pos = a.CurrentEpisode.Position
			}

			ui := episode.NewUI(m.UUID, m.Episodes, index, m.Result)

			anime := &database.Anime{
				ID:   m.Details.Id,
				Name: m.Details.Title.Romaji,
				CurrentEpisode: &database.Episode{
					Number:   index,
					Position: pos,
				},
				MaxEpisodes: len(m.Episodes),
			}

			if err := database.UpdateAnimeTracking(anime); err != nil {
				panic(err)
			}

			return constants.SwitchUI(m, ui, ui.UUID)
		}

	case constants.StreamingUrlsMsg:
		err := exec.Command("vlc", msg.Data.Url).Start()
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
