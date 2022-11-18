package search_results_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

type Keymap struct {
	Enter  key.Binding
	Return key.Binding
	Quit   key.Binding
}

var keys = Keymap{
	Enter: key.NewBinding(
		key.WithKeys("enter")),
	Return: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "return to search")),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c")),
}

var titleStyle = lipgloss.NewStyle().Background(lipgloss.Color("#007700"))

type UI struct {
	tea.Model
	UUID uuid.UUID
	list list.Model

	Results  []*constants.AnimeResult
	selected *constants.AnimeResult

	init     bool
	switched bool
	loading  bool

	keys Keymap
}
