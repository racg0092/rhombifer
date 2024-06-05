package rhombifer

type Run func(args ...string) error

type Command struct {
	// command name
	Name string

	// short description showed when the help command is run
	ShortDesc string

	// Long command description
	LongDesc string

	// flags if any
	Flags []Flag

	// Sub commands for this command
	Subs map[string]Command

	// Action perform by the command
	Run Run

	// Signifies if this is the root command
	Root bool

	// Signifies if there are no more commands after this one
	Leaf bool
}

// Adds a flag to the a command
func (cmd *Command) AddFlag(f Flag) {
	cmd.Flags = append(cmd.Flags, f)
}

// Adds a sub command to a command
func (cmd *Command) AddSub(command Command) {
	if cmd == nil {
		panic("attempting to set sub command to a nil reference")
	}
	cmd.Subs[command.Name] = command
}
