package search_ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Keymap struct {
	Enter key.Binding
	Quit  key.Binding
}

var keys = Keymap{
	Enter: key.NewBinding(
		key.WithKeys("enter")),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c")),
}

type UI struct {
	tea.Model
	spinner   spinner.Model
	textInput textinput.Model

	loading  bool
	nothing  bool
	switched bool
	init     bool

	httpErr error

	keys Keymap
}
