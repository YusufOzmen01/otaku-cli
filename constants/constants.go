package constants

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	"io"
	"os/exec"
	"strconv"
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

type Episode struct {
	list.Item

	EpisodeId  string `json:"episodeId"`
	EpisodeNum string `json:"episodeNum"`
	EpisodeUrl string `json:"episodeUrl"`
}

type AnimeDetails struct {
	AnimeTitle    string     `json:"animeTitle"`
	Type          string     `json:"type"`
	ReleasedDate  string     `json:"releasedDate"`
	Status        string     `json:"status"`
	Genres        []string   `json:"genres"`
	OtherNames    string     `json:"otherNames"`
	Synopsis      string     `json:"synopsis"`
	AnimeImg      string     `json:"animeImg"`
	TotalEpisodes string     `json:"totalEpisodes"`
	EpisodesList  []*Episode `json:"episodesList"`
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

var uiMap = make(map[uuid.UUID]tea.Model)

func SwitchUI(self tea.Model, targetModel tea.Model, targetUUID uuid.UUID) (tea.Model, tea.Cmd) {
	uiMap[targetUUID] = self

	return targetModel, func() tea.Msg {
		return nil
	}
}

func ReturnUI(selfUUID uuid.UUID) (tea.Model, tea.Cmd) {
	parent := uiMap[selfUUID]
	delete(uiMap, selfUUID)

	return parent, func() tea.Msg {
		return nil
	}
}

func KillProcessByNameWindows(processName string) int {
	kill := exec.Command("taskkill", "/im", processName, "/T", "/F")
	err := kill.Run()
	if err != nil {
		return -1
	}

	return 0
}

type AnimeResultDelegate struct{}
type AnimeEpisodesDelegate struct {
	AnimeID string
}

func (i AnimeResult) Title() string {
	return i.AnimeTitle
}

func (i Episode) EpisodeTitle() string {
	return fmt.Sprintf("Episode %s", i.EpisodeNum)
}

func (i Episode) FilterValue() string {
	return i.EpisodeNum
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

func (d AnimeEpisodesDelegate) Height() int {
	return 1
}

func (d AnimeResultDelegate) Spacing() int {
	return 0
}

func (d AnimeEpisodesDelegate) Spacing() int {
	return 0
}

func (d AnimeResultDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d AnimeEpisodesDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
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

func (d AnimeEpisodesDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(*Episode)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i.EpisodeTitle())

	lastWatched, err := database.GetAnimeProgress(d.AnimeID)
	if err == nil {
		num1, _ := strconv.Atoi(i.EpisodeNum)
		num2, _ := strconv.Atoi(lastWatched.LastWatchedEpisode)

		if num1 <= num2 {
			str = lipgloss.NewStyle().Foreground(lipgloss.Color("#4d4d4d")).Render(str)
		}

		if i.EpisodeNum == lastWatched.LastWatchedEpisode {
			str += " " + lipgloss.NewStyle().Italic(true).Render("(Currently on this episode)")
		}
	}

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(str string) string {
			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#00aa00")).Render(str)
		}
	}

	fmt.Fprint(w, fn(str))
}
