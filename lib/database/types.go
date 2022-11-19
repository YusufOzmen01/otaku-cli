package database

type Anime struct {
	ID                   string `clover:"id"`
	Name                 string `clover:"name"`
	LastWatchedEpisodeID string `clover:"last_watched_episode"`
}
