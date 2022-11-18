package search_ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewUI(parent tea.Model) UI {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#00aa00"))

	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 32
	ti.Width = 32

	return UI{
		keys:        keys,
		textInput:   ti,
		spinner:     s,
		ParentModel: parent,
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
