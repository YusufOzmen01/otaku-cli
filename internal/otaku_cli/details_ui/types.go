package details_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Keymap struct {
	EpisodeList key.Binding
	GoBack      key.Binding
}

var (
	keys = Keymap{
		EpisodeList: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "episode list")),
		GoBack: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "go back")),
	}
)

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.EpisodeList, k.GoBack}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.EpisodeList, k.GoBack},
	}
}

type UI struct {
	tea.Model
	UUID uuid.UUID

	*constants.AnimeDetails
	*constants.AnimeResult

	keys Keymap
	help help.Model
}
