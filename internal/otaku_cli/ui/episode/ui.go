package episode

import (
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(parentUUID uuid.UUID, episodes []*styles.Episode, currentEpisodeIndex int, details *styles.AnimeResult) UI {
	return UI{
		parentUUID:          parentUUID,
		episodes:            episodes,
		currentEpisodeIndex: currentEpisodeIndex,
		keys:                keys,
		help:                help.New(),
		UUID:                uuid.New(),
		details:             details,
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
