package rhombifer

import "sync"

type Run func(args []string) error

type Command struct {
	// command name
	Name string

	// short description showed when the help command is run
	ShortDesc string

	// Long command description
	LongDesc string

	// flags if any
	Flags []Flag
	Subs  []Command
	Run   Run

	Root bool
	Leaf bool
}

var root *Command

var once sync.Once

// Adds a flag to the a command
func (cmd *Command) AddFlag(f Flag) {
	cmd.Flags = append(cmd.Flags, f)
}

func (cmd *Command) AddSub(command Command) {
	cmd.Subs = append(cmd.Subs, command)
}

// Sets the root command
func SetRoot(cmd *Command) *Command {
	if cmd != nil {
		// for thread safety
		once.Do(func() {
			root = cmd
		})
	}
	return root
}

func Root() *Command {
	return root
}
