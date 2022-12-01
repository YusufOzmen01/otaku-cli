package database

type EpisodeProgress struct {
	CurrentEpisodeNumber     int
	MaxEpisodes              int
	CurrentPositionInEpisode int
	EpisodeLength            int
}

type Anime struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	EpisodeProgress *EpisodeProgress
}
