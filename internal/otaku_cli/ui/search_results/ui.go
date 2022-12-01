package search_results

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(results []*constants.AnimeResult) UI {
	return UI{
		Results: results,
		keys:    keys,
		UUID:    uuid.New(),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
