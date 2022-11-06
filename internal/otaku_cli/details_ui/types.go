package details_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Keymap struct {
	Quit   key.Binding
	Watch  key.Binding
	Cancel key.Binding
}

var keys = Keymap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c")),
	Watch: key.NewBinding(
		key.WithKeys("w")),
	Cancel: key.NewBinding(
		key.WithKeys("q")),
}

type UI struct {
	tea.Model
	ParentModel tea.Model

	*constants.AnimeDetails
	*constants.AnimeResult

	keys Keymap
}
