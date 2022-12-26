package episode

import (
	"github.com/YusufOzmen01/otaku-cli/lib/anime"
	"github.com/YusufOzmen01/otaku-cli/lib/vlc"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
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
	return []key.Binding{k.Previous, k.Next, k.GoBack}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Previous, k.Next, k.GoBack},
	}
}

type VLCMsg struct {
	Data *vlc.VLCData
}

type UI struct {
	tea.Model
	UUID uuid.UUID
	vlc  vlc.VLC

	episodes            []*anime.Episode
	parentUUID          uuid.UUID
	details             *anime.Result
	currentEpisodeIndex int
	init                bool
	episodeLoading      bool
	receivedData        bool
	currentVLCData      *vlc.VLCData
	source              string

	keys      Keymap
	help      help.Model
	progress1 progress.Model
	progress2 progress.Model
}
