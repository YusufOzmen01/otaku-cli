package episode

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"time"
)

func (m UI) View() string {
	currentEpisode := float64(m.currentEpisodeIndex + 1)
	maxEpisodes := float64(len(m.episodes))
	data := ""
	pos := 0.0
	length := 0.1

	if m.currentProgress != nil {
		pos = float64(m.currentProgress.Time)
		length = float64(m.currentProgress.Length) + 0.0001
	}

	if m.episodeLoading {
		data += styles.OngoingStyle.Render("Starting up mpv...")
	} else if m.mpvLoading {
		data += styles.OngoingStyle.Render(fmt.Sprintf("[%s] Waiting for response", m.source))
	} else {
		data += styles.CompletedStyle.Render(fmt.Sprintf("[%s] Playing", m.source))
	}

	data += styles.DetailStyle.Render(fmt.Sprintf("\nEpisode:  ")) + m.progress1.ViewAs(currentEpisode/maxEpisodes) + fmt.Sprintf(" %d/%d", int(currentEpisode), int(maxEpisodes))
	data += styles.DetailStyle.Render(fmt.Sprintf("\nPosition: "))
	if m.currentProgress == nil || !m.currentProgress.Paused {
		data += m.progress1.ViewAs(pos/length) + fmt.Sprintf(" %s/%s", time.Time{}.Add(time.Duration(pos)*time.Second).Format("04:05"), time.Time{}.Add(time.Duration(length)*time.Second).Format("04:05"))
	} else if m.currentProgress.Paused {
		data += m.progress2.ViewAs(pos/length) + fmt.Sprintf(" %s/%s", time.Time{}.Add(time.Duration(pos)*time.Second).Format("04:05"), time.Time{}.Add(time.Duration(length)*time.Second).Format("04:05"))
	}

	return data + "\n\n" + m.help.View(m.keys)
}
