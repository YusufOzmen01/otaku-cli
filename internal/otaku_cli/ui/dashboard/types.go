package dashboard

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Keymap struct {
	Search key.Binding
	MyList key.Binding
	Quit   key.Binding
}

var (
	keys = Keymap{
		Search: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "search")),
		MyList: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "my list")),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "q"),
			key.WithHelp("q/ctrl+c", "quit the app")),
	}
)

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Search, k.MyList, k.Quit}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Search, k.MyList, k.Quit},
	}
}

type UI struct {
	tea.Model
	UUID uuid.UUID

	keys Keymap
	help help.Model
}
