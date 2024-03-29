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

	anime, err := database.GetAnimeProgress(m.Result.Id)
	if err == nil {
		currentEpisode = float64(anime.CurrentEpisode.Number + 1)
		position = float64(anime.CurrentEpisode.Position)
		length = float64(anime.CurrentEpisode.Length)
	}

	data += styles.TitleStyle.Render("About the Anime")
	data += styles.DetailStyle.Render(fmt.Sprintf("\nTitle: %s", m.Details.Title))
	data += styles.DetailStyle.Render("\nStatus: ")
	switch m.Details.Status {
	case "FINISHED":
		data += styles.CompletedStyle.Render(m.Details.Status)
	case "ONGOING":
		data += styles.OngoingStyle.Render(m.Details.Status)
	default:
		data += styles.DetailStyle.Render(m.Details.Status)
	}
	data += styles.DetailStyle.Render(fmt.Sprintf("\nGenres: %s", strings.Join(m.Details.Genres, ",")))

	data += styles.TitleStyle.Render("\n\nProgress")
	data += styles.DetailStyle.Render("\nEpisode: ") + m.progress.ViewAs(currentEpisode/float64(len(m.Episodes))) + fmt.Sprintf(" %d/%d", int(currentEpisode), len(m.Episodes))
	data += styles.DetailStyle.Render("\nTime:    ") + m.progress.ViewAs(position/length) + fmt.Sprintf(" %s/%s", time.Time{}.Add(time.Duration(position)*time.Second).Format("04:05"), time.Time{}.Add(time.Duration(length)*time.Second).Format("04:05"))

	data += "\n\n" + m.help.View(m.keys)

	return data
}
