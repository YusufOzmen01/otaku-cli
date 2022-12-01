package details

import (
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI(details *styles.AnimeDetails, result *styles.AnimeResult) UI {
	return UI{
		AnimeDetails: details,
		AnimeResult:  result,
		keys:         keys,
		help:         help.New(),
		UUID:         uuid.New(),
		progress:     progress.New(progress.WithScaledGradient("#024f0d", "#05a11b")),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
