package database

import (
	"encoding/json"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/akrylysov/pogreb"
	"sync"
)

var (
	DB *pogreb.DB
)

const (
	AnimeTrackingKey   = "anime/tracking/%s"
	EpisodeTrackingKey = "episode/tracking/%s/%d"
	AnimeKey           = "anime/details/%s"
)

func InitializeDatabase(dbPath string) error {
	var err error

	DB, err = pogreb.Open(dbPath, nil)
	return err
}

func UpsertData(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	exists, err := DB.Has([]byte(key))
	if err != nil {
		return err
	}

	if exists {
		if err := DB.Delete([]byte(key)); err != nil {
			return err
		}
	}

	return DB.Put([]byte(key), data)
}

func GetData(key string) ([]byte, error) {
	exists, err := DB.Has([]byte(key))
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("key not found")
	}

	return DB.Get([]byte(key))
}

func UpdateAnimeTracking(anime *Anime) error {
	key := fmt.Sprintf(AnimeTrackingKey, anime.ID)

	a, err := GetData(key)
	if err == nil {
		data := new(Anime)
		if err := json.Unmarshal(a, data); err != nil {
			return err
		}

		if data.CurrentEpisode.Number > anime.CurrentEpisode.Number {
			return nil
		}

		if data.CurrentEpisode.Position > anime.CurrentEpisode.Position && data.CurrentEpisode.Number == anime.CurrentEpisode.Number {
			anime.CurrentEpisode.Position = data.CurrentEpisode.Position
		}
	}

	go UpdateAnimeData(anime.ID, nil)

	return UpsertData(key, anime)
}

func UpdateAnimeData(id string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	key := fmt.Sprintf(AnimeKey, id)

	data, err := anime.GetAnimeDetails(id)
	if err == nil {
		if err := UpsertData(key, data); err != nil {
			return
		}
	}
}

func GetAnimeDetails(id string) (*anime.Details, error) {
	data, err := GetData(fmt.Sprintf(AnimeKey, id))
	if err != nil {
		return nil, err
	}

	details := new(anime.Details)
	if err := json.Unmarshal(data, details); err != nil {
		return nil, err
	}

	return details, nil
}

func UpdateEpisode(episode *Episode, animeID string) error {
	return UpsertData(fmt.Sprintf(EpisodeTrackingKey, animeID, episode.Number), episode)
}

func GetEpisodeProgress(episodeNumber int, animeID string) (*Episode, error) {
	episode, err := GetData(fmt.Sprintf(EpisodeTrackingKey, animeID, episodeNumber))
	if err != nil {
		return nil, err
	}

	data := new(Episode)
	if err := json.Unmarshal(episode, data); err != nil {
		return nil, err
	}

	return data, nil
}

func GetAnimeProgress(animeID string) (*Anime, error) {
	lastWatched, err := GetData(fmt.Sprintf(AnimeTrackingKey, animeID))
	if err != nil {
		return nil, err
	}

	data := new(Anime)
	if err := json.Unmarshal(lastWatched, data); err != nil {
		return nil, err
	}

	return data, nil
}

func ResetAnimeProgress(animeID string, episodeCount int) error {
	if err := DB.Delete([]byte(fmt.Sprintf(AnimeTrackingKey, animeID))); err != nil {
		return err
	}

	for i := 0; i < episodeCount; i++ {
		if err := DB.Delete([]byte(fmt.Sprintf(EpisodeTrackingKey, animeID, i))); err != nil {
			return err
		}
	}

	return nil
}

func GetAllAnimes() ([]*Anime, error) {
	it := DB.Items()
	animeList := make([]*Anime, 0)
	for {
		_, value, err := it.Next()
		if err == pogreb.ErrIterationDone {
			break
		}

		if err != nil {
			return nil, err
		}

		data := new(Anime)
		if err := json.Unmarshal(value, data); err == nil {
			animeList = append(animeList, data)
		}
	}

	return animeList, nil
}
