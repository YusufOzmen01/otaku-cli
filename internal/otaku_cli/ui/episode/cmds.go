package episode

import (
	"github.com/YusufOzmen01/otaku-cli/lib/vlc"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func (m UI) vlcUpdate(v vlc.VLC) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Millisecond * 100)

		if v != nil {
			data, err := v.GetVLCData()
			if err == nil {
				return VLCMsg{Data: data}
			}
		}

		return nil
	}
}
