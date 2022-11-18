package dashboard_ui

import "github.com/charmbracelet/lipgloss"

func (m UI) View() string {
	data := lipgloss.NewStyle().Bold(true).Render("Dashboard (TODO)\n")

	return data + "\n" + m.help.View(m.keys)
}
