package search

import (
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	tea "github.com/charmbracelet/bubbletea"
	"sync"
)

func updateAnimes(results []*anime.Result) tea.Cmd {
	return func() tea.Msg {
		wg := &sync.WaitGroup{}
		for _, result := range results {
			wg.Add(1)
			go database.UpdateAnimeData(result.Id, wg)
		}

		wg.Wait()

		return anime.SearchDoneMsg{Data: results}
	}
}
