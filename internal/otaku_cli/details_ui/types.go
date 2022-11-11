package details_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Keymap struct {
	Watch  key.Binding
	GoBack key.Binding
}

var (
	keys = Keymap{
		Watch: key.NewBinding(
			key.WithKeys("w"),
			key.WithHelp("w", "watch")),
		GoBack: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "go back")),
	}
)

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Watch, k.GoBack}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Watch, k.GoBack},
	}
}

type UI struct {
	tea.Model
	ParentModel tea.Model

	*constants.AnimeDetails
	*constants.AnimeResult

	keys Keymap
	help help.Model
}
