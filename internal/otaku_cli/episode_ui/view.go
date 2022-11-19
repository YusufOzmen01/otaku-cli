package episode_ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func (m UI) View() string {
	data := lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("Episode %s ", m.episodes[m.currentEpisodeIndex].EpisodeNum))

	if m.episodeLoading {
		data += lipgloss.NewStyle().Italic(true).Render("(Loading)")
	}

	return data + "\n\n" + m.help.View(m.keys)
}