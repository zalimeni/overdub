package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zalimeni/overdub/history"
	"os"
	"strconv"
)

type model struct {
	choices  []string
	cursor   int
	selected int
}

func loadCommandHistory() []string {
	//TODO support options for reading more/less history
	commands, err := history.ReadLocalHistory(5)
	if err != nil {
		fmt.Println("Error reading shell history: ", err)
		// Return empty commands list (allow Custom command entry)
		return nil
	}
	return commands
}

func initialModel() model {
	// Recent commands, and "*Custom*" for direct input
	initialChoices := append(loadCommandHistory(), "*Custom*")

	return model{
		choices: initialChoices,
		// Default to first choice
		cursor: 0,
		// Default to first choice
		selected: 0,
	}
}

func (m model) Init() tea.Cmd {
	// No I/O to start
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		// Enter + Space toggle the choice list selection
		case "enter", " ":
			m.selected = m.cursor
		default:
			// Check for numeric key -> select corresponding option
			// Only support numeric options 1-9 (single key)
			firstKey := msg.Runes[0]
			if firstKey >= '1' && firstKey <= '9' && len(msg.Runes) == 1 {
				choice, _ := strconv.Atoi(string(firstKey))
				// If the pressed numeric key is <= the number of choices available,
				// select and move the cursor to that number using the proper index.
				if choice <= len(m.choices) {
					m.cursor = choice - 1
					m.selected = choice - 1
				}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Select command to overdub:\n\n"

	// Iterate over choices
	for i, choice := range m.choices {
		// Support selection by number on keypad
		number := i+1
		numberShortcut := strconv.Itoa(number) + ")"
		if number > 9 {
			numberShortcut = "  "
		}

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // display cursor
		}

		// Is this choice selected?
		checked := " " // not selected
		if m.selected == i {
			checked = "x" // selected
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s %s\n", cursor, checked, numberShortcut, choice)
	}

	s += "\nPress `q` to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Oops! Something went wrong: %v", err)
		os.Exit(1)
	}
}
