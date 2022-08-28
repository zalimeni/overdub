# overdub
Create new commands from your shell history or terminal.

## TODO list

### Initial release

- Actually support making a dubbed function (basic)
- Install subcommand
  - Manage added commands in env-var-pathed separate file
  - Add post-execution alias to source file of added fns when main command is run
  - NTS: Using sourced fns takes care of auto-complete previously in this list
  - Maybe: Add note to README about recommended auto-complete options for supported shells


### Nice-to-haves
- Support parameterization 
- Filter out unlikely commands (e.g. package managers) from suggestions list
- Limit history pull to sane number for speed (500?)
- Dedupe commands to choose from
- Comment annotations to edit
  - Serialized original command?
  - Parse back persisted fn for edit 
