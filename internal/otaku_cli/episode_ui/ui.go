package episode_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(parentUUID uuid.UUID, episodes []*constants.Episode, currentEpisodeIndex int) UI {
	return UI{
		parentUUID:          parentUUID,
		episodes:            episodes,
		currentEpisodeIndex: currentEpisodeIndex,
		keys:                keys,
		help:                help.New(),
		UUID:                uuid.New(),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
