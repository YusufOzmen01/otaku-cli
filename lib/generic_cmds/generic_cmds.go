package generic_cmds

import (
	"github.com/YusufOzmen01/otaku-cli/constants"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func Wait(duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(duration)

		return constants.WaitMsg{}
	}
}
