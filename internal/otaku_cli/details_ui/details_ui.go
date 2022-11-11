package details_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

func NewUI(parentModel tea.Model, details *constants.AnimeDetails, result *constants.AnimeResult) UI {
	return UI{
		ParentModel:  parentModel,
		AnimeDetails: details,
		AnimeResult:  result,
		keys:         keys,
		help:         help.New(),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
