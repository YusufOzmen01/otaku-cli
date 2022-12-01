package styles

import "github.com/charmbracelet/lipgloss"

var (
	SelectionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00dd00"))
	TitleStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#00dd00")).Bold(true)

	OngoingStyle   = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#cbcf0e"))
	CompletedStyle = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#00ad57"))
	DetailStyle    = lipgloss.NewStyle().Italic(true).PaddingLeft(2)
	WatchedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#4d4d4d"))
)
