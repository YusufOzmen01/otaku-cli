package details

import (
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Keymap struct {
	Watch       key.Binding
	Finished    key.Binding
	Reset       key.Binding
	EpisodeList key.Binding
	GoBack      key.Binding
	Quit        key.Binding
}

var (
	keys = Keymap{
		Watch: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "start watching")),
		Finished: key.NewBinding(
			key.WithKeys("f"),
			key.WithHelp("f", "mark as finished")),
		Reset: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "reset progress")),
		EpisodeList: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "episodes")),
		GoBack: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "go back")),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c")),
	}
)

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.GoBack, k.EpisodeList, k.Watch, k.Finished, k.Reset}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.GoBack, k.EpisodeList, k.Watch, k.Finished, k.Reset},
	}
}

type UI struct {
	tea.Model
	UUID uuid.UUID

	*anime.Details
	*anime.Result

	keys     Keymap
	help     help.Model
	progress progress.Model
}
