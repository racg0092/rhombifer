package rhombifer

import (
	"github.com/racg0092/rhombifer/pkg/models"
	"strings"
)

type Run func(args ...string) error

type Command struct {
	// command name
	Name string

	// short description showed when the help command is run
	ShortDesc string

	// Long command description
	LongDesc string

	// flags if any
	Flags []*models.Flag

	// Pointers to required flags if any
	requiredFlags []*models.Flag

	// Sub commands for this command
	Subs map[string]*Command

	// Action perform by the command
	Run Run

	// Signifies if this is the root command
	Root bool

	// Signifies if there are no more commands after this one
	Leaf bool
}

// Adds a flag to the a command
func (cmd *Command) AddFlags(flags ...*models.Flag) {
	for _, f := range flags {
		if f.Required {
			if cmd.requiredFlags == nil {
				cmd.requiredFlags = make([]*models.Flag, 0)
			}
			cmd.requiredFlags = append(cmd.requiredFlags, f)
		}
		if f.RequiresValue && f.Values == nil {
			f.Values = make([]string, 0)
		}
		cmd.Flags = append(cmd.Flags, f)
	}
}

// Adds a sub command to a command
func (cmd *Command) AddSub(command *Command) {
	if cmd == nil {
		panic("attempting to set sub command to a nil reference")
	}
	cmd.Subs[command.Name] = command
}

// Validates if required flags are found in the input string. If any required flag is missing it returns false
// otherwise true. If no flags are required it returns true.
func (cmd *Command) ValidateRequiredFlags(args []string) bool {
	if len(cmd.requiredFlags) <= 0 {
		return true
	}

	if len(args) == 0 {
		return false
	}

	var missing bool = false
	joinArgs := strings.Join(args, " ")
	for _, f := range cmd.requiredFlags {
		if !strings.Contains(joinArgs, "--"+f.Name) && !strings.Contains(joinArgs, "-"+f.ShortFormat) {
			missing = true
			break
		}
	}

	return !missing
}

// Get required flags
func (cmd *Command) RequiredFlags() *[]*models.Flag {
	return &cmd.requiredFlags
}
