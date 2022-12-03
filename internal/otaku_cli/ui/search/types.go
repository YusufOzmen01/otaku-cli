package search

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Keymap struct {
	Enter  key.Binding
	GoBack key.Binding
	Quit   key.Binding
}

var keys = Keymap{
	Enter: key.NewBinding(
		key.WithKeys("enter")),
	GoBack: key.NewBinding(
		key.WithKeys("esc")),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c")),
}

type UI struct {
	tea.Model
	UUID      uuid.UUID
	spinner   spinner.Model
	textInput textinput.Model

	loading    bool
	searchDone bool
	nothing    bool
	switched   bool
	init       bool

	httpErr error

	keys Keymap
}
