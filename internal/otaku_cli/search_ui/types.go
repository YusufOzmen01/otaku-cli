package search_ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

type Keymap struct {
	Enter key.Binding
	Quit  key.Binding
}

var keys = Keymap{
	Enter: key.NewBinding(
		key.WithKeys("enter")),
	Quit: key.NewBinding(
		key.WithKeys("esc")),
}

var (
	titleStyle = lipgloss.NewStyle().Background(lipgloss.Color("#007700"))
	textStyle  = lipgloss.NewStyle().Bold(true).Italic(true).MarginLeft(2)
)

type UI struct {
	tea.Model
	UUID      uuid.UUID
	spinner   spinner.Model
	textInput textinput.Model

	loading  bool
	nothing  bool
	switched bool
	init     bool

	httpErr error

	keys Keymap
}
