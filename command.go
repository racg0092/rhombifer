package rhombifer

import (
	"sync"
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
	Flags []Flag

	Subs map[string]Command

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

// Adds a sub command to a command
func (cmd *Command) AddSub(command Command) {
	if cmd == nil {
		panic("attempting to set sub command to a nil reference")
	}
	cmd.Subs[command.Name] = command
}

// Todo refactor execute should only exeute the current command Run function and pipe the output to its children
func (cmd Command) ExecCommand(name string, args []string) error {
	if cmd.Run != nil {
		cmd.Run()
	}
	return nil
}

// Takes in a pointer to `cmd` and sets it as the root command. The **root** command is only set once
// for the application runtime. It means that `root` will be set only the first time this funtion is call
func SetRoot(cmd *Command) {
	if cmd != nil {
		cmd.Root = true
		// for thread safety
		once.Do(func() {
			root = cmd
		})
	}
}

// Returns root command. If root has not been initialized it creates a new empty [Command]
// and returns the pointer
func Root() *Command {
	if root == nil {
		c := Command{
			Root: true,
			Subs: make(map[string]Command),
		}
		SetRoot(&c)
	}
	return root
}
