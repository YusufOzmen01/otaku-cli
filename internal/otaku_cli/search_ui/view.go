package search_ui

import "fmt"

func (m UI) View() string {
	if m.loading {
		return fmt.Sprintf("%s Searching for \"%s\"", m.spinner.View(), m.textInput.Value())
	}

	if m.nothing {
		return fmt.Sprintf("Nothing found. Please try again")
	}

	return fmt.Sprintf("Search an anime %s", m.textInput.View())
}