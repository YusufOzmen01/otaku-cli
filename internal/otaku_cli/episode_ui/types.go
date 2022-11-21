package episode_ui

import (
	"encoding/xml"
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
	return []key.Binding{k.Next, k.GoBack, k.Previous}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Next, k.GoBack, k.Previous},
	}
}

type VLCMsg struct {
	Data *Root
}

type Root struct {
	XMLName      xml.Name `xml:"root"`
	Text         string   `xml:",chardata"`
	Fullscreen   string   `xml:"fullscreen"`
	Aspectratio  string   `xml:"aspectratio"`
	Audiodelay   string   `xml:"audiodelay"`
	Apiversion   string   `xml:"apiversion"`
	Currentplid  string   `xml:"currentplid"`
	Time         string   `xml:"time"`
	Volume       string   `xml:"volume"`
	Length       string   `xml:"length"`
	Random       string   `xml:"random"`
	Audiofilters struct {
		Text    string `xml:",chardata"`
		Filter0 string `xml:"filter_0"`
	} `xml:"audiofilters"`
	Rate         string `xml:"rate"`
	Videoeffects struct {
		Text       string `xml:",chardata"`
		Hue        string `xml:"hue"`
		Saturation string `xml:"saturation"`
		Contrast   string `xml:"contrast"`
		Brightness string `xml:"brightness"`
		Gamma      string `xml:"gamma"`
	} `xml:"videoeffects"`
	State         string `xml:"state"`
	Loop          string `xml:"loop"`
	Version       string `xml:"version"`
	Position      string `xml:"position"`
	Repeat        string `xml:"repeat"`
	Subtitledelay string `xml:"subtitledelay"`
	Equalizer     string `xml:"equalizer"`
	Information   struct {
		Text     string `xml:",chardata"`
		Category []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
			Info []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"info"`
		} `xml:"category"`
	} `xml:"information"`
	Stats struct {
		Text                string `xml:",chardata"`
		Lostabuffers        string `xml:"lostabuffers"`
		Readpackets         string `xml:"readpackets"`
		Lostpictures        string `xml:"lostpictures"`
		Demuxreadbytes      string `xml:"demuxreadbytes"`
		Demuxbitrate        string `xml:"demuxbitrate"`
		Playedabuffers      string `xml:"playedabuffers"`
		Demuxcorrupted      string `xml:"demuxcorrupted"`
		Sendbitrate         string `xml:"sendbitrate"`
		Sentbytes           string `xml:"sentbytes"`
		Displayedpictures   string `xml:"displayedpictures"`
		Demuxreadpackets    string `xml:"demuxreadpackets"`
		Sentpackets         string `xml:"sentpackets"`
		Inputbitrate        string `xml:"inputbitrate"`
		Demuxdiscontinuity  string `xml:"demuxdiscontinuity"`
		Averagedemuxbitrate string `xml:"averagedemuxbitrate"`
		Decodedvideo        string `xml:"decodedvideo"`
		Averageinputbitrate string `xml:"averageinputbitrate"`
		Readbytes           string `xml:"readbytes"`
		Decodedaudio        string `xml:"decodedaudio"`
	} `xml:"stats"`
}

type UI struct {
	tea.Model
	UUID uuid.UUID

	episodes            []*constants.Episode
	parentUUID          uuid.UUID
	details             *constants.AnimeResult
	currentEpisodeIndex int
	init                bool
	episodeLoading      bool
	currentVLCData      *Root

	keys Keymap
	help help.Model
}
