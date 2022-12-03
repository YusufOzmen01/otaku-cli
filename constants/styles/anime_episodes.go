package styles

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"strconv"
	"time"
)

type AnimeEpisodesDelegate struct {
	AnimeID string
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
	i, ok := listItem.(*anime.Episode)
	if !ok {
		return
	}

	size := len(strconv.Itoa(len(m.VisibleItems())))

	str := fmt.Sprintf("%s", i.EpisodeTitle())
	pos := 0.0
	length := 0.1

	currentEpisode := false

	data, err := database.GetAnimeProgress(d.AnimeID)
	if err == nil {
		num1, err := strconv.Atoi(i.EpisodeNum)
		if err != nil {
			panic(err)
		}

		if num1-1 < data.CurrentEpisode.Number || data.Finished {
			str = WatchedStyle.Render(str)
		}

		if num1-1 == data.CurrentEpisode.Number && !data.Finished {
			currentEpisode = true
		}
	}

	progressBar := progress.New(progress.WithScaledGradient("#024f0d", "#05a11b"), progress.WithoutPercentage(), progress.WithWidth(20))
	for i := 0; i < size-len(strconv.Itoa(index+1))+1; i++ {
		str += " "
	}

	if data != nil {
		episodeID, err := strconv.Atoi(i.EpisodeNum)
		if err != nil {
			panic(err)
		}

		episode, err := database.GetEpisodeProgress(episodeID-1, data.ID)
		if err == nil {
			pos = float64(episode.Position)
			length = float64(episode.Length)
		}
	}

	str += fmt.Sprintf(" %s %s/%s", progressBar.ViewAs(pos/length), time.Time{}.Add(time.Duration(pos)*time.Second).Format("04:05"), time.Time{}.Add(time.Duration(length)*time.Second).Format("04:05"))

	if currentEpisode {
		str += " " + lipgloss.NewStyle().Italic(true).Bold(true).Render("<- Currently on this episode")
	}

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(str string) string {
			return SelectionStyle.Render("→ " + str)
		}
	}

	fmt.Fprint(w, fn(str))
}
