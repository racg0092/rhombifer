package rhombifer

import (
	"fmt"
	"sync"
)

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

	Subs []Command

	Run Run

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
	if cmd == nil {
		c := Command{}
		SetRoot(&c)
		cmd.Subs = append(cmd.Subs, command)
		return
	}
	cmd.Subs = append(cmd.Subs, command)
}

func (cmd Command) ExecCommand(name string, args []string) error {
	found := false
	for _, sub := range cmd.Subs {
		if sub.Name == name {
			found = true
			if sub.Run == nil {
				return fmt.Errorf("Command %s does not have a run function", name)
			}

			err := sub.Run(args)

			if err != nil {
				return err
			}

		}
	}

	if !found {
		return fmt.Errorf("Command %s does not exist", name)
	}

	return nil
}

// Sets the root command
func SetRoot(cmd *Command) {
	if cmd != nil {
		cmd.Root = true
		// for thread safety
		once.Do(func() {
			root = cmd
		})
	}
}

// Get Root Command
func Root() *Command {
	return root
}
