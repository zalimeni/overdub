package shell

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ParserFunc returns the tokenized command portion from a given raw
// line of history. If parsing fails, returns an error.
type ParserFunc func(string) ([]string, error)

type HistorySource interface {
	// GetHistory returns the last N lines of history from this source.
	// If 0 is provided, returns all lines.
	GetHistory(int) []string
}

func handleReadHistoryErr(err error) {
	if err != nil {
		panic(err)
	}
}

// readLines reads all available lines from a bufio.Scanner and
// returns them as a []string.
func readLines(scanner *bufio.Scanner) []string {
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	handleReadHistoryErr(scanner.Err())
	return lines
}

// getLastNLines gets the last `numLines` valid lines from a []string
// of lines. Validity is determined by the string->bool predicate.
// By default, the predicate used is isValidLine.
func getLastNLines(lines []string, numLines int, predicate func(string) bool) []string {
	lineCount := len(lines)
	var validLines []string
	for i := lineCount-1; i >= 0; i-- {
		if predicate(lines[i]) {
			validLines = append(validLines, lines[i])
			if len(validLines) == numLines {
				return validLines
			}
		}
	}
	return validLines
}

// isValidLine accepts only lines that have at least one argument.
// The assumption is that any plain 0-args command is not a candidate
// for auto-selection to create a new command alias.
func isValidLine(line string) bool {
	return len(strings.Fields(line)) > 1
}

func getHistoryFromFile(filePath string, numLines int) []string {
	file, err := os.Open(filePath)
	defer file.Close()
	handleReadHistoryErr(err)

	scanner := bufio.NewScanner(file)
	lines := readLines(scanner)

	return getLastNLines(lines, numLines, isValidLine)
}

//func getHistoryFromCommand(historyCmd *exec.Cmd, numLines int) [][]string {
//	historyCmd.Stderr = nil
//	r, err := historyCmd.StdoutPipe()
//	handleReadHistoryErr(err)
//
//	scanner := bufio.NewScanner(r)
//	handleReadHistoryErr(historyCmd.Start())
//	var lines = readSplitLines(scanner)
//	handleReadHistoryErr(historyCmd.Wait())
//
//	return getLastNLines(lines, numLines)
//}

func GetDefaultShell() (HistorySource, error) {
	shellPath := strings.TrimSpace(os.Getenv("SHELL"))
	if shellPath == "" {
		return nil, errors.New("$SHELL is not set, cannot detect shell type")
	}
	shellName := filepath.Base(shellPath)
	if shellName == "." || shellName == string(filepath.Separator) {
		return nil, errors.New("$SHELL could not be parsed to obtain a shell name")
	}

	fmt.Printf("Detected `%v` shell\n", shellName)

	switch shellName {
	case "bash":
		return &Bash{}, nil
	}
	return nil, errors.New("$SHELL contained an unsupported shell: " + shellName)
}
