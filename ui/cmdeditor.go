package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Assignment int

const (
	keep Assignment = iota
	positional_param
	previous
	skip
)

type TokenAssignment struct {
	tokenIndex int
	assignment Assignment
}

type CommandPart struct {
	value string
}

type CommandEditorModel struct {
	tokens      []string
	assignments []Assignment
	name        string
	cursor      int
}

// initialCommandName returns the suggested name of a command based on the input
// command being dubbed. The value returned is an "acronym" based on the first letter
// of each token of the command. For flag tokens, the first non-`-` is used.
func initialCommandName(fields []string) string {
	var ret string
	for _, f := range fields {
		for _, c := range f {
			if c != '-' {
				ret += string(c)
				break
			}
		}
	}
	return ret
}

func NewCommandEditorModel(cmdStr string) *CommandEditorModel {
	fields := strings.Fields(cmdStr)
	return &CommandEditorModel{
		tokens: fields,
		name:   initialCommandName(fields),
	}
}

func (m CommandEditorModel) Init() tea.Cmd {
	// No I/O to start
	return nil
}

func (m CommandEditorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			//TODO: change assignment
		case "down", "j":
			//TODO: change assignment
		case " ":
			//TODO: toggle skip vs. default assignment
		case "enter":
			//TODO: name command dialog (esc -> return to edit, enter -> save)
		default:
			//TODO: ignore? error on some?
		}
	}

	return m, nil
}

func (m CommandEditorModel) View() string {
	s := "Coming soon: command editor!\n\n"

	//TODO: implement
	s += "Suggested command name: " + m.name

	s += "\nPress `q` to quit.\n"

	return s
}
