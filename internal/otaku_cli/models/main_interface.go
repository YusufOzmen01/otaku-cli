package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os/exec"
	"time"
)

const ApiUrl = "https://gogoanime.consumet.org"

type errMsg struct {
	err error
}

type resultMsg struct {
	data []*Result
}

type detailMsg struct {
	data *AnimeDetails
}

type waitMsg struct {
}

type MainInterfaceKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Quit   key.Binding
	Watch  key.Binding
	Cancel key.Binding
}

var keys = MainInterfaceKeyMap{
	Up: key.NewBinding(
		key.WithKeys("up")),
	Down: key.NewBinding(key.WithKeys("down")),
	Enter: key.NewBinding(
		key.WithKeys("enter")),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c")),
	Watch: key.NewBinding(
		key.WithKeys("w")),
	Cancel: key.NewBinding(
		key.WithKeys("c")),
}

type Result struct {
	AnimeId    string `json:"animeId"`
	AnimeTitle string `json:"animeTitle"`
	AnimeUrl   string `json:"animeUrl"`
	AnimeImg   string `json:"animeImg"`
	Status     string `json:"status"`
}

type AnimeDetails struct {
	AnimeTitle    string   `json:"animeTitle"`
	Type          string   `json:"type"`
	ReleasedDate  string   `json:"releasedDate"`
	Status        string   `json:"status"`
	Genres        []string `json:"genres"`
	OtherNames    string   `json:"otherNames"`
	Synopsis      string   `json:"synopsis"`
	AnimeImg      string   `json:"animeImg"`
	TotalEpisodes string   `json:"totalEpisodes"`
	EpisodesList  []struct {
		EpisodeId  string `json:"episodeId"`
		EpisodeNum string `json:"episodeNum"`
		EpisodeUrl string `json:"episodeUrl"`
	} `json:"episodesList"`
}

type MainInterface struct {
	tea.Model
	keys MainInterfaceKeyMap

	results        []*Result
	cursor         int
	selected       *Result
	selectedDetail *AnimeDetails
	searchLoading  bool
	detailLoading  bool
	showDetails    bool

	httpErr  error
	httpDone bool

	spinner   spinner.Model
	textInput textinput.Model
}

func InitMainInterface() MainInterface {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Blink(true)

	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return MainInterface{
		keys:      keys,
		textInput: ti,
		results:   make([]*Result, 0),
		spinner:   s,
	}
}

func (m MainInterface) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m MainInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			if m.httpDone && !m.showDetails {
				if m.cursor > 0 {
					m.cursor--
				}
			}

		case key.Matches(msg, m.keys.Down):
			if m.httpDone && !m.showDetails {
				if m.cursor < len(m.results)-1 {
					m.cursor++
				}
			}

		case key.Matches(msg, m.keys.Enter):
			if m.httpDone {
				m.detailLoading = true
				m.selected = m.results[m.cursor]

				return m, m.getAnimeDetails
			} else {
				m.searchLoading = true
				m.cursor = 0

				return m, m.searchAnime
			}

		case key.Matches(msg, m.keys.Watch):
			if m.showDetails {
				exec.Command("explorer", m.selected.AnimeUrl).Start()

				return m, tea.Quit
			}

		case key.Matches(msg, m.keys.Cancel):
			if m.showDetails {
				m.showDetails = false
				m.httpDone = false
				m.selected = nil
				m.results = nil
				m.textInput.SetValue("")
				m.cursor = 0

				return m, nil
			}

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}

		switch msg.String() {
		default:
			if !m.searchLoading && !m.httpDone {
				var cmd tea.Cmd
				m.textInput, cmd = m.textInput.Update(msg)

				return m, cmd
			}
		}

	case errMsg:
		m.httpErr = msg.err
		return m, tea.Quit

	case resultMsg:
		m.searchLoading = false
		m.httpDone = true
		m.results = msg.data
		if len(msg.data) == 0 {
			return m, Wait(time.Second * 3)
		}

		m.textInput.SetValue("")

		return m, nil

	case waitMsg:
		if len(m.results) == 0 {
			m.httpDone = false
			m.results = nil

			return m, nil
		}

	case detailMsg:
		m.selectedDetail = msg.data
		m.showDetails = true
		m.detailLoading = false

		return m, nil

	default:
		var cmd tea.Cmd

		if m.searchLoading || m.detailLoading {
			m.spinner, cmd = m.spinner.Update(msg)

			return m, cmd
		}
	}

	return m, nil
}

func (m MainInterface) View() string {
	if m.searchLoading {
		return fmt.Sprintf("%s Searching for \"%s\"", m.spinner.View(), m.textInput.Value())
	}

	if m.detailLoading {
		return fmt.Sprintf("%s Loading details, hang on...", m.spinner.View())
	}

	var data string

	if m.httpDone {
		if len(m.results) == 0 {
			return fmt.Sprintf("Nothing found. Please try again")
		}

		if m.showDetails {
			data = fmt.Sprintf("%s\n\n", m.selected.AnimeTitle)
			data += fmt.Sprintf("Title: %s\n", m.selectedDetail.AnimeTitle)
			data += fmt.Sprintf("Type: %s\n", m.selectedDetail.Type)
			data += fmt.Sprintf("Status: %s\n", m.selectedDetail.Status)
			data += fmt.Sprintf("Genres: %s\n", m.selectedDetail.Genres)
			data += fmt.Sprintf("Other Name: %s\n", m.selectedDetail.OtherNames)
			data += fmt.Sprintf("Episode Count: %d\n\n", len(m.selectedDetail.EpisodesList))
			data += fmt.Sprintf("[w] Watch anime\n[c] Return to search")

			return data
		}

		data = "Anime Search Result \n\n"

		for i, result := range m.results {
			if m.cursor == i {
				data += ">"
			}

			data += fmt.Sprintf(" %s (%s)\n", result.AnimeTitle, result.AnimeUrl)
		}

		return data
	}

	return fmt.Sprintf("Search an anime %s", m.textInput.View())
}

func Wait(duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(duration)

		return waitMsg{}
	}
}

func (m MainInterface) searchAnime() tea.Msg {
	url := fmt.Sprintf(ApiUrl+"/search?keyw=%s", m.textInput.Value())

	resp, status, err := network.ProcessGet(context.Background(), url, nil)
	if err != nil {
		return errMsg{err: err}
	}

	if status != 200 {
		return errMsg{err: fmt.Errorf("server returned %d", status)}
	}

	data := new([]*Result)

	if err := json.Unmarshal(resp, data); err != nil {
		return errMsg{err: err}
	}

	return resultMsg{data: *data}
}

func (m MainInterface) getAnimeDetails() tea.Msg {
	url := fmt.Sprintf(ApiUrl+"/anime-details/%s", m.selected.AnimeId)

	resp, status, err := network.ProcessGet(context.Background(), url, nil)
	if err != nil {
		return errMsg{err: err}
	}

	if status != 200 {
		return errMsg{err: fmt.Errorf("server returned %d", status)}
	}

	data := new(AnimeDetails)

	if err := json.Unmarshal(resp, data); err != nil {
		return errMsg{err: err}
	}

	return detailMsg{data: data}
}
