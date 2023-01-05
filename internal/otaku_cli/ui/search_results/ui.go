package search_results

import (
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(results []*anime.Result, searchText string) UI {
	return UI{
		Results:    results,
		keys:       keys,
		UUID:       uuid.New(),
		SearchText: searchText,
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
