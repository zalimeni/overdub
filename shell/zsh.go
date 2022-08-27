package shell

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Zsh struct{}

func (sh *Zsh) parseFileHistoryLine(line string) (string, error) {
	_, cmd, _ := strings.Cut(line, ";")
	if len(cmd) == 0 {
		return "", errors.New("encountered empty line while parsing")
	}
	return cmd, nil
}

func (sh *Zsh) GetHistory(numLines int) []string {
	histFilePath, ok := os.LookupEnv("HISTFILE")
	if !ok {
		histFilePath = filepath.Join(os.Getenv("HOME"), ".zsh_history")
	}
	rawLines := getHistoryFromFile(histFilePath, numLines)
	var parsedLines []string
	for i := 1; i < len(rawLines); i++ {
		parsed, err := sh.parseFileHistoryLine(rawLines[i])
		if err != nil {
			log.Println("Warn: " + err.Error())
		} else {
			parsedLines = append(parsedLines, parsed)
		}
	}
	return parsedLines
}
