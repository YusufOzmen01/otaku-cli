package styles

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"strconv"
)

type Episode struct {
	list.Item

	EpisodeId  string `json:"episodeId"`
	EpisodeNum string `json:"episodeNum"`
	EpisodeUrl string `json:"episodeUrl"`
}

type AnimeEpisodesDelegate struct {
	AnimeID string
}

func (i Episode) EpisodeTitle() string {
	return fmt.Sprintf("Episode %s", i.EpisodeNum)
}

func (i Episode) FilterValue() string {
	return i.EpisodeNum
}

func (d AnimeEpisodesDelegate) Height() int {
	return 1
}

func (d AnimeEpisodesDelegate) Spacing() int {
	return 0
}

func (d AnimeEpisodesDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d AnimeEpisodesDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(*Episode)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i.EpisodeTitle())

	lastWatched, err := database.GetAnimeProgress(d.AnimeID)
	if err == nil {
		num1, err := strconv.Atoi(i.EpisodeNum)
		if err != nil {
			panic(err)
		}

		if num1-1 < lastWatched.CurrentEpisode.EpisodeNumber || lastWatched.Finished {
			str = WatchedStyle.Render(str)
		}

		if num1-1 == lastWatched.CurrentEpisode.EpisodeNumber && !lastWatched.Finished {
			str += " " + lipgloss.NewStyle().Italic(true).Bold(true).Render("<- Currently on this episode")
		}
	}

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(str string) string {
			return SelectionStyle.Render("→ " + str)
		}
	}

	fmt.Fprint(w, fn(str))
}
