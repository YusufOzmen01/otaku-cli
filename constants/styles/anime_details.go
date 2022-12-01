package styles

import "github.com/charmbracelet/bubbles/list"

type AnimeDetails struct {
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
