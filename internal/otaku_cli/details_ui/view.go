package details_ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func (m UI) View() string {
	data := lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("%s\n", m.AnimeDetails.AnimeTitle))
	data += lipgloss.NewStyle().Italic(true).Render(fmt.Sprintf(""+
		"\nType: %s\n"+
		"Status: %s\n"+
		"Genres: %s\n"+
		"Other Name: %s\n"+
		"Episode Count: %d\n",
		m.AnimeDetails.Type,
		m.AnimeDetails.Status,
		m.AnimeDetails.Genres,
		m.AnimeDetails.OtherNames,
		len(m.AnimeDetails.EpisodesList)))
	//data += fmt.Sprintf("[w] Watch anime\n[q] Return to results")
	data += "\n" + m.help.View(m.keys)

	return data
}
