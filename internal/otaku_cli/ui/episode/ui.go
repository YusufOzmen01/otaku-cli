package episode

import (
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/progress"
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
		progress1:           progress.New(progress.WithScaledGradient("#024f0d", "#05a11b")),
		progress2:           progress.New(progress.WithScaledGradient("#676907", "#cbcf0e")),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
