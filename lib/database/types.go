package database

type Episode struct {
	Number   int
	Position int
	Length   int
}

type Anime struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Finished       bool
	CurrentEpisode *Episode
	MaxEpisodes    int
}
