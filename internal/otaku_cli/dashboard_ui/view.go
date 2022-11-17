package dashboard_ui

func (m UI) View() string {
	return m.help.View(m.keys)
}
