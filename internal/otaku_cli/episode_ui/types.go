package episode_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Keymap struct {
	Previous key.Binding
	Next     key.Binding
	GoBack   key.Binding
	Quit     key.Binding
}

var (
	keys = Keymap{
		Next: key.NewBinding(
			key.WithKeys("right"),
			key.WithHelp("→", "next episode")),
		Previous: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("←", "previous episode")),
		GoBack: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "go back")),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit the app")),
	}
)

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Next, k.Previous, k.GoBack}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Next, k.Previous, k.GoBack},
	}
}

type UI struct {
	tea.Model
	UUID uuid.UUID

	episodes            []*constants.Episode
	parentUUID          uuid.UUID
	currentEpisodeIndex int
	init                bool
	episodeLoading      bool

	keys Keymap
	help help.Model
}
