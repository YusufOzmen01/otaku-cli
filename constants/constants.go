package constants

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
)

const ApiUrl = "https://gogoanime.consumet.org"

type ErrMsg struct {
	Err error
}

type ResultMsg struct {
	Data []*AnimeResult
}

type DetailMsg struct {
	Data *AnimeDetails
}

type StreamResultData struct {
	Data *StreamData
}

type WaitMsg struct {
}

type AnimeResult struct {
	AnimeId    string `json:"animeId"`
	AnimeTitle string `json:"animeTitle"`
	AnimeUrl   string `json:"animeUrl"`
	AnimeImg   string `json:"animeImg"`
	Status     string `json:"status"`
}

type AnimeDetails struct {
	AnimeTitle    string   `json:"animeTitle"`
	Type          string   `json:"type"`
	ReleasedDate  string   `json:"releasedDate"`
	Status        string   `json:"status"`
	Genres        []string `json:"genres"`
	OtherNames    string   `json:"otherNames"`
	Synopsis      string   `json:"synopsis"`
	AnimeImg      string   `json:"animeImg"`
	TotalEpisodes string   `json:"totalEpisodes"`
	EpisodesList  []struct {
		EpisodeId  string `json:"episodeId"`
		EpisodeNum string `json:"episodeNum"`
		EpisodeUrl string `json:"episodeUrl"`
	} `json:"episodesList"`
}

type Stream struct {
	File  string `json:"file"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type StreamData struct {
	Referer string `json:"Referer"`
	Sources []struct {
		File  string `json:"file"`
		Label string `json:"label"`
		Type  string `json:"type"`
	} `json:"sources"`
	SourcesBk []struct {
		File  string `json:"file"`
		Label string `json:"label"`
		Type  string `json:"type"`
	} `json:"sources_bk"`
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

func (d AnimeResultDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
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
			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#00aa00")).Render(fmt.Sprintf("%d. %s", index+1, i.Title()))
		}
	}

	fmt.Fprint(w, fn(str))
}
