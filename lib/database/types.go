package database

type Episode struct {
	EpisodeNumber int
	Position      int
	EpisodeLength int
}

type Anime struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Finished       bool
	CurrentEpisode *Episode
	MaxEpisodes    int
}
