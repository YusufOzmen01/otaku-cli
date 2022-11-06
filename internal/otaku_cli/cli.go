package otaku_cli

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/search_ui"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type OtakuCli interface {
	Run()
}

type otakuCli struct {
	program *tea.Program
}

func NewOtakuCli() OtakuCli {
	return &otakuCli{
		program: tea.NewProgram(search_ui.NewUI(), tea.WithAltScreen()),
	}
}

func (oc *otakuCli) Run() {
	if err := oc.program.Start(); err != nil {
		fmt.Printf("An error occured: %s", err)
		os.Exit(1)
	}
}
