package database

import (
	"encoding/json"
	"fmt"
	"github.com/akrylysov/pogreb"
	"strconv"
)

var (
	DB *pogreb.DB
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

func WatchAnime(anime *Anime) error {
	key := fmt.Sprintf("a[%s]", anime.ID)

	a, err := GetData(key)
	if err == nil {
		data := new(Anime)
		if err := json.Unmarshal(a, data); err != nil {
			return err
		}

		num1, err := strconv.Atoi(data.LastWatchedEpisode)
		if err != nil {
			return err
		}

		num2, err := strconv.Atoi(anime.LastWatchedEpisode)
		if err != nil {
			return err
		}

		if num1 > num2 {
			anime.LastWatchedEpisode = data.LastWatchedEpisode
		}

		if anime.Position == 0 && data.LastWatchedEpisode == anime.LastWatchedEpisode {
			anime.Position = data.Position
		}
	}

	return UpsertData(key, anime)
}

func GetAnimeProgress(animeID string) (*Anime, error) {
	lastWatched, err := GetData(fmt.Sprintf("a[%s]", animeID))
	if err != nil {
		return nil, err
	}

	data := new(Anime)
	if err := json.Unmarshal(lastWatched, data); err != nil {
		return nil, err
	}

	return data, nil
}
