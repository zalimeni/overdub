package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zalimeni/overdub/ui"
)

func main() {
	p := tea.NewProgram(ui.NewHistoryPickerModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Oops! Something went wrong: %v", err)
		os.Exit(1)
	}
}
