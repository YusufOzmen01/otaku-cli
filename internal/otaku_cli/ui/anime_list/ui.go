package anime_list

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func NewUI() UI {
	return UI{
		keys: keys,
		help: help.New(),
		UUID: uuid.New(),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
