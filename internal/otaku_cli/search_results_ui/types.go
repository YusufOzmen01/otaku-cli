package search_results_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Keymap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Return key.Binding
	Quit   key.Binding
}

var keys = Keymap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up")),
	Down: key.NewBinding(key.WithKeys("down"),
		key.WithHelp("↓", "move down")),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("⏎", "select")),
	Return: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "return to search")),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c")),
}

var titleStyle = lipgloss.NewStyle().Background(lipgloss.Color("#007700"))

type UI struct {
	tea.Model
	ParentModel tea.Model
	list        list.Model

	Results  []*constants.AnimeResult
	selected *constants.AnimeResult

	init     bool
	switched bool
	loading  bool

	keys Keymap
}
