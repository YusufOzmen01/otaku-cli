package episodes_ui

func (m UI) View() string {
	return m.list.View() + "\n" + m.help.View(m.keys)
}
