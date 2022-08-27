# overdub
Create new commands from your shell history or terminal.

## TODO list for initial release
- Filter out unlikely commands (e.g. package managers) from suggestions list
- Actually support making a dubbed function (basic)
- Parsing of commands to aid creation of parameterized aliases (shell fns? command lib?)
- Install subcommand
  - Manage added commands in env-var-pathed separate file
  - Add post-execution alias to source file of added fns when main command is run
  - NTS: Using sourced fns takes care of auto-complete previously in this list
  - Maybe: Add note to README about recommended auto-complete options for supported shells
