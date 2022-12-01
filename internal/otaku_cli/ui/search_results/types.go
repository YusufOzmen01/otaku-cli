package search_results

import (
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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

type UI struct {
	tea.Model
	UUID uuid.UUID
	list list.Model

	Results  []*styles.AnimeResult
	selected *styles.AnimeResult

	init     bool
	switched bool
	loading  bool

	keys Keymap
}
