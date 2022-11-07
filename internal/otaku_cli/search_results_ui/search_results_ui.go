package search_results_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	tea "github.com/charmbracelet/bubbletea"
)

func NewUI(parentModel tea.Model, results []*constants.AnimeResult) UI {
	return UI{
		ParentModel: parentModel,
		Results:     results,
		keys:        keys,
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
