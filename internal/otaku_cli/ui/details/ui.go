package details

import (
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(details *anime.Details, result *anime.Result) UI {
	return UI{
		Details:  details,
		Result:   result,
		keys:     keys,
		help:     help.New(),
		UUID:     uuid.New(),
		progress: progress.New(progress.WithScaledGradient("#024f0d", "#05a11b"), progress.WithoutPercentage()),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
