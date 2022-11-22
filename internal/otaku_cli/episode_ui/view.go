package episode_ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func (m UI) View() string {
	str := fmt.Sprintf("Episode %d ", m.currentEpisodeIndex+1)

	if m.currentEpisodeIndex == len(m.episodes)-1 {
		str += lipgloss.NewStyle().Bold(true).Render("[Last Episode] ")
	}

	data := lipgloss.NewStyle().Bold(true).Render(str)

	if m.episodeLoading {
		data += lipgloss.NewStyle().Italic(true).Render("(Loading)")
	}

	return data + "\n\n" + m.help.View(m.keys)
}
