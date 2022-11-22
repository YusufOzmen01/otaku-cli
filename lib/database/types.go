package database

type EpisodeProgress struct {
	CurrentEpisodeIndex      int
	MaxEpisodes              int
	CurrentPositionInEpisode int
}

type Anime struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Finished        bool
	EpisodeProgress *EpisodeProgress
}
