package episode

import (
	"github.com/YusufOzmen01/otaku-cli/lib/mpv"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func (m UI) progressUpdate(mp mpv.MPV) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Millisecond * 100)

		if mp != nil {
			data, loading, closed := mp.GetData()
			if closed {
				return nil
			}

			if data != nil {
				return ProgressUpdate{
					Data: data,
				}
			}

			if loading {
				return ProgressUpdate{
					Data: &mpv.Progress{
						Loading: true,
					},
				}
			}
		}

		return nil
	}
}
