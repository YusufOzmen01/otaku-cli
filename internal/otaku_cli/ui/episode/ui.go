package episode

import (
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(parentUUID uuid.UUID, episodes []*anime.Episode, currentEpisodeIndex int, details *anime.Result) UI {

	return UI{
		parentUUID:          parentUUID,
		episodes:            episodes,
		currentEpisodeIndex: currentEpisodeIndex,
		keys:                keys,
		help:                help.New(),
		UUID:                uuid.New(),
		details:             details,
		progress1:           progress.New(progress.WithScaledGradient("#024f0d", "#05a11b"), progress.WithoutPercentage()),
		progress2:           progress.New(progress.WithScaledGradient("#676907", "#cbcf0e"), progress.WithoutPercentage()),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
