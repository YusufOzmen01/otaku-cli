package template

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Keymap struct {
	GoBack key.Binding
	Quit   key.Binding
}

var (
	keys = Keymap{
		GoBack: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "go back")),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit the app")),
	}
)

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.GoBack}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.GoBack},
	}
}

type UI struct {
	tea.Model
	UUID uuid.UUID

	keys Keymap
	help help.Model
}
