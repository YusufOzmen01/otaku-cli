package styles

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
)

type AnimeResult struct {
	list.Item

	AnimeId    string `json:"animeId"`
	AnimeTitle string `json:"animeTitle"`
	AnimeUrl   string `json:"animeUrl"`
	AnimeImg   string `json:"animeImg"`
	Status     string `json:"status"`
}

type AnimeResultDelegate struct{}

func (i AnimeResult) Title() string {
	return i.AnimeTitle
}

func (i AnimeResult) Description() string {
	return i.AnimeId
}

func (i AnimeResult) FilterValue() string {
	return i.AnimeTitle
}

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
	i, ok := listItem.(*AnimeResult)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Title())

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(str string) string {
			return SelectionStyle.Render("→ " + str)
		}
	}

	_, err := fmt.Fprint(w, fn(str))
	if err != nil {
		return
	}
}
