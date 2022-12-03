package search

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
)

func (m UI) View() string {
	if m.loading {
		if m.searchDone {
			return fmt.Sprintf("%s Resolving data...", m.spinner.View())
		}

		return fmt.Sprintf("%s Searching for \"%s\"", m.spinner.View(), m.textInput.Value())
	}

	if m.nothing {
		return fmt.Sprintf("Nothing found. Please try again")
	}

	return styles.TitleStyle.Render("Search Anime") + "\n\n" + styles.DetailStyle.Render(m.textInput.View())
}
