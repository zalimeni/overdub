// Package history supports reading and parsing history of the host shell
package history

import (
	"github.com/zalimeni/overdub/shell"
)

func ReadLocalHistory(numLines int) ([]string, error) {
	s, e := shell.GetDefaultShell()
	if e != nil {
		return nil, e
	}

	//TODO parse command history into subcommands, flags, args, and other tokens
	//  - Maybe use https://github.com/alecthomas/kong or some such
	return s.GetHistory(numLines), nil
}
