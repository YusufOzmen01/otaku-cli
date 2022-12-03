package episode

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants/styles"
	"strconv"
	"time"
)

func (m UI) View() string {
	currentEpisode := float64(m.currentEpisodeIndex + 1)
	maxEpisodes := float64(len(m.episodes))
	data := ""
	pos := 0.0
	length := 0.1

	if m.currentVLCData != nil {
		p, err := strconv.Atoi(m.currentVLCData.Time)
		if err != nil {
			panic(err)
		}

		l, err := strconv.Atoi(m.currentVLCData.Length)
		if err != nil {
			panic(err)
		}

		pos = float64(p)
		length = float64(l) + 0.0001
	}

	if m.episodeLoading {
		data += styles.OngoingStyle.Render("Starting up VLC ")
	} else if !m.receivedData {
		data += styles.OngoingStyle.Render("Waiting for response")
	} else {
		data += styles.CompletedStyle.Render("Playing")
	}

	data += styles.DetailStyle.Render(fmt.Sprintf("\nEpisode: ")) + m.progress1.ViewAs(currentEpisode/maxEpisodes) + fmt.Sprintf(" %d/%d", int(currentEpisode), int(maxEpisodes))
	data += styles.DetailStyle.Render(fmt.Sprintf("\nPosition: "))
	if m.currentVLCData == nil || m.currentVLCData.State == "playing" {
		data += m.progress1.ViewAs(pos/length) + fmt.Sprintf(" %s/%s", time.Time{}.Add(time.Duration(pos)*time.Second).Format("04:05"), time.Time{}.Add(time.Duration(length)*time.Second).Format("04:05"))
	} else if m.currentVLCData.State == "paused" {
		data += m.progress2.ViewAs(pos/length) + fmt.Sprintf(" %s/%s", time.Time{}.Add(time.Duration(pos)*time.Second).Format("04:05"), time.Time{}.Add(time.Duration(length)*time.Second).Format("04:05"))
	}

	return data + "\n\n" + m.help.View(m.keys)
}
