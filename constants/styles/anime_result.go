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
)

type AnimeResultDelegate struct{}

func (d AnimeResultDelegate) Height() int {
	return 1
}

func (d AnimeResultDelegate) Spacing() int {
	return 0
}

func (d AnimeResultDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (d AnimeResultDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(*anime.Result)
	if !ok {
		return
	}

	current, max := 0.0, 0.1

	size := 0
	for _, item := range m.Items() {
		if len(item.(*anime.Result).AnimeTitle) > size {
			size = len(item.(*anime.Result).AnimeTitle)
		}
	}

	str := fmt.Sprintf("%s", i.Title())

	progressBar := progress.New(progress.WithScaledGradient("#024f0d", "#05a11b"), progress.WithoutPercentage(), progress.WithWidth(20))
	for j := 0; j < size-len(i.AnimeTitle)+1; j++ {
		str += " "
	}

	data, err := database.GetAnimeProgress(i.AnimeId)
	if err == nil {
		current = float64(data.CurrentEpisode.Number) + 1
	}

	details, err := database.GetAnimeDetails(i.AnimeId)
	if err == nil {
		max = float64(len(details.EpisodesList))
	}

	str += fmt.Sprintf(" %s %d/%d", progressBar.ViewAs(current/max), int(current), int(max))

	a, err := database.GetAnimeProgress(i.AnimeId)
	if err == nil {
		if a.Finished {
			str = CompletedStyle.Render(str)
		} else {
			str = OngoingStyle.Render(str)
		}
	}

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(str string) string {
			return SelectionStyle.Render("→ " + str)
		}
	}

	_, err = fmt.Fprint(w, fn(str))
	if err != nil {
		return
	}
}
