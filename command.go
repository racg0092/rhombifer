package rhombifer

import "github.com/racg0092/rhombifer/pkg/models"

type Run func(args ...string) error

type Command struct {
	// command name
	Name string

	// short description showed when the help command is run
	ShortDesc string

	// Long command description
	LongDesc string

	RequiredFlags []string

	// flags if any
	Flags []models.Flag

	// Flags found when parsing input. It holds a pointer to the flags in [Flags]
	FoundFlags []*models.Flag

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
func (cmd *Command) AddFlag(f models.Flag) {
	if f.RequiresValue && f.Values == nil {
		f.Values = make([]string, 0)
	}
	cmd.Flags = append(cmd.Flags, f)
}

// Adds a sub command to a command
func (cmd *Command) AddSub(command Command) {
	if cmd == nil {
		panic("attempting to set sub command to a nil reference")
	}
	cmd.Subs[command.Name] = command
}

// Sets the value of [FoundFlags] in [Command] to <nil>
func (cmd *Command) EmptyFoundFlags() {
	cmd.FoundFlags = nil
}
