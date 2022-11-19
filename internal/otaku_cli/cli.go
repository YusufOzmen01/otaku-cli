package otaku_cli

import (
	"github.com/YusufOzmen01/otaku-cli/internal/otaku_cli/dashboard_ui"
	"github.com/YusufOzmen01/otaku-cli/lib/database"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"path"
)

type OtakuCli interface {
	Run()
}

type otakuCli struct {
	program *tea.Program
}

func NewOtakuCli() OtakuCli {
	return &otakuCli{
		program: tea.NewProgram(dashboard_ui.NewUI(), tea.WithAltScreen()),
	}
}

func (oc *otakuCli) Run() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	if err := database.InitializeDatabase(path.Join(home, ".otaku-cli")); err != nil {
		panic(err)
	}

	defer database.DB.Close()

	if err := oc.program.Start(); err != nil {
		panic(err)
	}
}
