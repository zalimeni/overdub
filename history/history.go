// Package history supports reading and parsing history of the host shell
package history

func ReadLocal() (error, []string) {
	//TODO load command history based on $SHELL
	//TODO parse command history into subcommands, flags, args, and other tokens
	//  - Maybe use https://github.com/alecthomas/kong or some such
	return nil, []string{"This", "is -a test"}
}


