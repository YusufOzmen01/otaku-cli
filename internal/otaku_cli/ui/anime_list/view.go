package anime_list

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"github.com/charmbracelet/lipgloss"
	"time"
)

func (m UI) View() string {
	currentEpisode := 1
	position := 0
	finished := false

	anime, err := database.GetAnimeProgress(m.AnimeResult.AnimeId)
	if err == nil {
		currentEpisode = anime.EpisodeProgress.CurrentEpisodeNumber + 1
		position = anime.EpisodeProgress.CurrentPositionInEpisode
		finished = anime.Finished
	}

	data := lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("%s\n", m.AnimeDetails.AnimeTitle))
	data += lipgloss.NewStyle().Italic(true).Render(fmt.Sprintf(""+
		"\nType: %s\n"+
		"Status: %s\n"+
		"Genres: %s\n"+
		"Other Name: %s\n"+
		"Episode Count: %d\n"+
		"Current Episode: %d\n"+
		"Position: %s\n"+
		"Finished: %t",
		m.AnimeDetails.Type,
		m.AnimeDetails.Status,
		m.AnimeDetails.Genres,
		m.AnimeDetails.OtherNames,
		len(m.AnimeDetails.EpisodesList),
		currentEpisode,
		time.Time{}.Add(time.Duration(position)*time.Second).Format("04:05"),
		finished))

	//data += fmt.Sprintf("[w] Watch anime\n[q] Return to results")
	data += "\n" + m.help.View(m.keys)

	return data
}
