package anime

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
)

type Result struct {
	list.Item

	AnimeId    string `json:"animeId"`
	AnimeTitle string `json:"animeTitle"`
	AnimeUrl   string `json:"animeUrl"`
	AnimeImg   string `json:"animeImg"`
	Status     string `json:"status"`
}

func (i Result) Title() string {
	return i.AnimeTitle
}

func (i Result) Description() string {
	return i.AnimeId
}

func (i Result) FilterValue() string {
	return i.AnimeTitle
}

type Details struct {
	list.Item

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

type Episode struct {
	list.Item

	EpisodeId  string `json:"episodeId"`
	EpisodeNum string `json:"episodeNum"`
	EpisodeUrl string `json:"episodeUrl"`
}

func (i Episode) EpisodeTitle() string {
	return fmt.Sprintf("Episode %s", i.EpisodeNum)
}

func (i Episode) FilterValue() string {
	return i.EpisodeNum
}

type ResultMsg struct {
	Data []*Result
}

type SearchDoneMsg struct {
	Data []*Result
}

type DetailsMsg struct {
	Data *Details
}
