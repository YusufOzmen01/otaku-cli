package episodes

import (
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(episodes []*styles.Episode, details *styles.AnimeResult) UI {
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
