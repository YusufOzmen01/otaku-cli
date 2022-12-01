package search_results

import (
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(results []*styles.AnimeResult) UI {
	return UI{
		Results: results,
		keys:    keys,
		UUID:    uuid.New(),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
