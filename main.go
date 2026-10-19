package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zalimeni/overdub/ui"
)

func main() {
	p := tea.NewProgram(ui.NewHistoryPickerModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops! Something went wrong: %v", err)
		os.Exit(1)
	}
}
