package episodes_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(episodes []*constants.Episode, details *constants.AnimeResult) UI {
	return UI{
		keys:     keys,
		help:     help.New(),
		episodes: episodes,
		UUID:     uuid.New(),
		details:  details,
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
