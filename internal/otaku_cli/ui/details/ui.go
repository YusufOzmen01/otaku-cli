package details

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(details *constants.AnimeDetails, result *constants.AnimeResult) UI {
	return UI{
		AnimeDetails: details,
		AnimeResult:  result,
		keys:         keys,
		help:         help.New(),
		UUID:         uuid.New(),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
