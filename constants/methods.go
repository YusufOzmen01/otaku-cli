package constants

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"os/exec"
)

var uiMap = make(map[uuid.UUID]tea.Model)

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

func KillProcessByNameWindows(processName string) int {
	kill := exec.Command("taskkill", "/im", processName, "/T", "/F")
	err := kill.Run()
	if err != nil {
		return -1
	}

	return 0
}

func ReverseSlice[T any](original []T) (reversed []T) {
	reversed = make([]T, len(original))
	copy(reversed, original)

	for i := len(reversed)/2 - 1; i >= 0; i-- {
		tmp := len(reversed) - 1 - i
		reversed[i], reversed[tmp] = reversed[tmp], reversed[i]
	}

	return
}
