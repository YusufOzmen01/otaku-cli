package dashboard_ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Keymap struct {
	Search      key.Binding
	LastWatched key.Binding
	Quit        key.Binding
}

var (
	keys = Keymap{
		Search: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "search")),
		LastWatched: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "last watched")),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "q"),
			key.WithHelp("q/ctrl+c", "quit the app")),
	}
)

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Search, k.LastWatched, k.Quit}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Search, k.LastWatched, k.Quit},
	}
}

type UI struct {
	tea.Model

	keys Keymap
	help help.Model
}
