package otaku_cli

import (
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/models"
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
		program: tea.NewProgram(models.InitMainInterface()),
	}
}

func (oc *otakuCli) Run() {
	if err := oc.program.Start(); err != nil {
		fmt.Printf("An error occured: %s", err)
		os.Exit(1)
	}
}
