package episode

import (
	"context"
	"encoding/xml"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
	tea "github.com/charmbracelet/bubbletea"
)

func (m UI) vlcUpdate() tea.Msg {
	body, status, err := network.ProcessGet(context.Background(), "http://localhost:58000/requests/status.xml", map[string]string{"Authorization": "Basic OmFtb25ndXNfaXNfZnVubnk="})
	if err != nil {
		return m.vlcUpdate()
	}

	if status != 200 {
		return m.vlcUpdate()
	}

	data := new(Root)

	if err := xml.Unmarshal(body, data); err != nil {
		return m.vlcUpdate()
	}

	return VLCMsg{Data: data}
}
