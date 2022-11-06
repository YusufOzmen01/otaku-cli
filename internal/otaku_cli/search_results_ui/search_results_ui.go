package search_results_ui

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewUI(parentModel tea.Model, results []*constants.AnimeResult) UI {
	return UI{
		ParentModel: parentModel,
		Results:     results,
		keys:        keys,
	}
}

func (m UI) Init() tea.Cmd {
	items := make([]list.Item, 0)

	for _, result := range m.Results {
		items = append(items, list.Item(result))
	}

	m.list = list.New(items, constants.AnimeResultDelegate{}, 0, 20)
	m.list.Title = "Anime Search Result"
	m.list.SetShowStatusBar(true)
	m.list.SetFilteringEnabled(true)
	m.list.Styles.Title = lipgloss.NewStyle().MarginLeft(0)
	m.list.Styles.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(2)
	m.list.Styles.HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(2).PaddingBottom(0)

	return nil
}
