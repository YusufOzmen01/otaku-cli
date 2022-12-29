package constants

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"math/rand"
	"os/exec"
	"runtime"
)

var (
	uiMap   = make(map[uuid.UUID]tea.Model)
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SwitchUI(self tea.Model, targetModel tea.Model, targetUUID uuid.UUID) (tea.Model, tea.Cmd) {
	uiMap[targetUUID] = self

	return targetModel, func() tea.Msg {
		return nil
	}
}

func ReturnUI(selfUUID uuid.UUID) (tea.Model, tea.Cmd) {
	parent := uiMap[selfUUID]
	delete(uiMap, selfUUID)

	return parent, func() tea.Msg {
		return nil
	}
}

func KillProcessByName(processName string) error {
	if runtime.GOOS == "windows" {
		kill := exec.Command("taskkill", "/im", processName+".exe", "/T", "/F")
		err := kill.Run()
		if err != nil {
			return err
		}

		return nil
	} else {
		kill := exec.Command("killall", processName)
		err := kill.Run()
		if err != nil {
			return err
		}

		return nil
	}
}
