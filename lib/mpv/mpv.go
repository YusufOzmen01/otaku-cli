package mpv

import (
	"fmt"
	"github.com/DexterLB/mpvipc"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"os"
	"os/exec"
	"path"
	"runtime"
)

type MPV interface {
	GetData() (*Progress, bool, bool)
	Kill()
}

type mpv struct {
	connection    *mpvipc.Connection
	events        chan *mpvipc.Event
	stopListening chan struct{}
	process       *exec.Cmd
	pipePath      string
	file          string
	subtitles     string
	time          int
	currentData   *Progress
	loading       bool
	closed        bool
}

func NewMPV(file, subtitles string, start int) MPV {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	pipePath := path.Join(home, "mpv-pipe")

	if runtime.GOOS == "windows" {
		pipePath = "\\\\.\\pipe\\pipe"
	}

	class := &mpv{
		pipePath:  pipePath,
		file:      file,
		subtitles: subtitles,
		time:      start,
		loading:   true,
	}

	go class.start()

	return class
}

func (m *mpv) Kill() {
	constants.KillProcessByName("mpv")
}

func (m *mpv) GetData() (*Progress, bool, bool) {
	return m.currentData, m.loading, m.closed
}

func (m *mpv) start() {
	params := []string{fmt.Sprintf("--playlist=%s", m.file), fmt.Sprintf("--start=%d", m.time), fmt.Sprintf("--input-ipc-server=%s", m.pipePath)}

	if len(m.subtitles) > 0 {
		params = append(params, fmt.Sprintf("--sub-file=%s", m.subtitles))
	}

	process := exec.Command("mpv", params...)

	if err := process.Start(); err != nil {
		panic(err)
	}

	conn := mpvipc.NewConnection(m.pipePath)

	for {
		if err := conn.Open(); err == nil {
			break
		}
	}

	for {
		if _, err := conn.Get("duration"); err == nil {
			break
		}
	}

	events, stopListening := conn.NewEventListener()

	_, err := conn.Call("observe_property", 1, "time-pos")
	if err != nil {
		fmt.Print(err)
	}

	_, err = conn.Call("observe_property", 2, "pause")
	if err != nil {
		fmt.Print(err)
	}

	data, err := conn.Get("duration")
	if err != nil {
		panic(err)
	}

	m.connection = conn
	m.events = events
	m.stopListening = stopListening
	m.currentData = &Progress{
		Time:   m.time,
		Paused: false,
		Length: int(data.(float64)),
	}
	m.process = process

	go m.closeHandler()
	go m.pullHandler()
}

func (m *mpv) closeHandler() {
	m.connection.WaitUntilClosed()
	m.stopListening <- struct{}{}
	m.closed = true
}

func (m *mpv) pullHandler() {
	m.loading = false

	for event := range m.events {
		switch event.ID {
		case 1:
			if event.Data != nil {
				m.currentData.Time = int(event.Data.(float64))
			}

		case 2:
			if event.Data != nil {
				m.currentData.Paused = event.Data.(bool)
			}
		}
	}
}
