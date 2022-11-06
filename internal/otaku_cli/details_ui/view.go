package details_ui

import "fmt"

func (m UI) View() string {
	data := fmt.Sprintf("%s\n\n", m.AnimeDetails.AnimeTitle)
	data += fmt.Sprintf("Title: %s\n", m.AnimeDetails.AnimeTitle)
	data += fmt.Sprintf("Type: %s\n", m.AnimeDetails.Type)
	data += fmt.Sprintf("Status: %s\n", m.AnimeDetails.Status)
	data += fmt.Sprintf("Genres: %s\n", m.AnimeDetails.Genres)
	data += fmt.Sprintf("Other Name: %s\n", m.AnimeDetails.OtherNames)
	data += fmt.Sprintf("Episode Count: %d\n\n", len(m.AnimeDetails.EpisodesList))
	data += fmt.Sprintf("[w] Watch anime\n[q] Return to results")

	return data
}
