package details

import (
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Keymap struct {
	Watch       key.Binding
	EpisodeList key.Binding
	GoBack      key.Binding
	Quit        key.Binding
}

var (
	keys = Keymap{
		Watch: key.NewBinding(
			key.WithKeys("w"),
			key.WithHelp("w", "start watching/continue watching where you left off")),
		EpisodeList: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "episode list")),
		GoBack: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "go back")),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit the app")),
	}
)

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.EpisodeList, k.GoBack, k.Watch}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.EpisodeList, k.GoBack, k.Watch},
	}
}

type UI struct {
	tea.Model
	UUID uuid.UUID

	*styles.AnimeDetails
	*styles.AnimeResult

	keys     Keymap
	help     help.Model
	progress progress.Model
}
