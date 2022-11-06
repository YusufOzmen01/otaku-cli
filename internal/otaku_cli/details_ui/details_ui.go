package details_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	tea "github.com/charmbracelet/bubbletea"
)

func NewUI(parentModel tea.Model, details *constants.AnimeDetails, result *constants.AnimeResult) UI {
	return UI{
		ParentModel:  parentModel,
		AnimeDetails: details,
		AnimeResult:  result,
		keys:         keys,
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
