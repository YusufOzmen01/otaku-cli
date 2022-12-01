package episodes

import (
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Keymap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Return key.Binding
	Quit   key.Binding
}

var (
	keys = Keymap{
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
)

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Enter, k.Return}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Enter, k.Return},
	}
}

type UI struct {
	tea.Model
	UUID uuid.UUID

	details *styles.AnimeResult

	keys Keymap
	help help.Model
	list list.Model

	init     bool
	episodes []*styles.Episode
}
