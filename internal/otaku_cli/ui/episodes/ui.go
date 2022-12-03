package episodes

import (
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(episodes []*anime.Episode, details *anime.Result) UI {
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
