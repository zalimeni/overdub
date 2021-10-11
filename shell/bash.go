package shell

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Bash struct {}

func (sh *Bash) parseFileHistoryLine(line string) (string, error) {
	parsed := strings.TrimSpace(line)
	if len(parsed) == 0 {
		return "", errors.New("encountered empty line while parsing")
	}
	return parsed, nil
}

func (sh *Bash) GetHistory(numLines int) []string {
	histFilePath, ok := os.LookupEnv("HISTFILE"); if !ok {
		histFilePath = filepath.Join(os.Getenv("HOME"), ".bash_history")
	}
	// Fetch twice as many lines because bash history is 2 lines per-entry (timestamp, command).
	// When parsing, skip every other line beginning w/ first line for the same reason.
	rawLines := getHistoryFromFile(histFilePath, numLines*2)

	var parsedLines []string
	for i := 1; i < len(rawLines); i+=2 {
		parsed, err := sh.parseFileHistoryLine(rawLines[i])
		if err != nil {
			log.Println("Warn: " + err.Error())
		} else {
			parsedLines = append(parsedLines, parsed)
		}
	}
	return parsedLines
}
