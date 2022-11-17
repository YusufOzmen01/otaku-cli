package dashboard_ui

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

func NewUI() UI {
	return UI{
		keys: keys,
		help: help.New(),
	}
}

func (m UI) Init() tea.Cmd {
	return nil
}
