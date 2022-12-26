package anime

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"strconv"
)

type SearchResult struct {
	Data []*Result `json:"data"`
}

type Title struct {
	Native        string `json:"native"`
	Romaji        string `json:"romaji"`
	English       string `json:"english"`
	UserPreferred string `json:"userPreferred"`
}

type Result struct {
	list.Item

	Id         string `json:"id"`
	AnimeTitle *Title `json:"title"`
	Status     string `json:"status"`
}

func (i Result) Title() string {
	return i.AnimeTitle.Romaji
}

func (i Result) Description() string {
	return i.Id
}

func (i Result) FilterValue() string {
	return i.AnimeTitle.Romaji
}

type Source struct {
	Id     string `json:"id"`
	Target string `json:"target"`
}

type Details struct {
	list.Item

	Id          string     `json:"id"`
	Title       *Title     `json:"title"`
	AnilistId   int        `json:"anilistId"`
	MalId       int        `json:"malId"`
	Type        string     `json:"type"`
	ReleaseDate int        `json:"releaseDate"`
	Rating      int        `json:"rating"`
	Status      string     `json:"status"`
	Genres      []string   `json:"genre"`
	Episodes    []*Episode `json:"episodes"`
}

type Episode struct {
	list.Item

	EpisodeId  string    `json:"id"`
	EpisodeNum int       `json:"number"`
	Sources    []*Source `json:"sources"`
}

func (i Episode) EpisodeTitle() string {
	return fmt.Sprintf("Episode %d", i.EpisodeNum)
}

func (i Episode) FilterValue() string {
	return strconv.Itoa(i.EpisodeNum)
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
