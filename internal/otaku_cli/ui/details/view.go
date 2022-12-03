package details

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	"strings"
	"time"
)

func (m UI) View() string {
	currentEpisode := 0.0
	position := 0.0
	length := 1.0
	data := ""

	anime, err := database.GetAnimeProgress(m.AnimeResult.AnimeId)
	if err == nil {
		currentEpisode = float64(anime.CurrentEpisode.Number + 1)
		position = float64(anime.CurrentEpisode.Position)
		length = float64(anime.CurrentEpisode.Length)
	}

	data += styles.TitleStyle.Render("About the Anime")
	data += styles.DetailStyle.Render(fmt.Sprintf("\nTitle: %s", m.AnimeDetails.AnimeTitle))
	data += styles.DetailStyle.Render("\nStatus: ")
	switch m.AnimeDetails.Status {
	case "Completed":
		data += styles.CompletedStyle.Render(m.AnimeDetails.Status)
	case "Ongoing":
		data += styles.OngoingStyle.Render(m.AnimeDetails.Status)
	default:
		data += styles.DetailStyle.Render(m.AnimeDetails.Status)
	}
	data += styles.DetailStyle.Render(fmt.Sprintf("\nGenres: %s", strings.Join(m.AnimeDetails.Genres, ",")))

	data += styles.TitleStyle.Render("\n\nProgress")
	data += styles.DetailStyle.Render("\nEpisode: ") + m.progress.ViewAs(currentEpisode/float64(len(m.EpisodesList))) + fmt.Sprintf(" %d/%d", int(currentEpisode), len(m.EpisodesList))
	data += styles.DetailStyle.Render("\nTime:    ") + m.progress.ViewAs(position/length) + fmt.Sprintf(" %s/%s", time.Time{}.Add(time.Duration(position)*time.Second).Format("04:05"), time.Time{}.Add(time.Duration(length)*time.Second).Format("04:05"))

	data += "\n\n" + m.help.View(m.keys)

	return data
}
