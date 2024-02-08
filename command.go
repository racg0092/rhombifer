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

func (cmd Command) ExecCommand(name string, args []string) {
	for _, sub := range cmd.Subs {
		if sub.Name == name {
			if sub.Run != nil {
				sub.Run(args)
			}
		}
	}
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

func UseBuiltInHelp(short, long *string) Command {
	help := Command{
		Name:      "help",
		ShortDesc: "Displays help information",
		LongDesc: `
		Displays help information for the specified command or the root command if no command is specified.
		`,
		Leaf: true,
		Run: func(args []string) error {

			if len(args) == 0 {
				fmt.Println(root.Name)
				if root.LongDesc != "" {
					fmt.Printf("\n%s", root.LongDesc)
				} else {
					fmt.Printf("\n%s", root.ShortDesc)
				}
			}

			return nil
		},
	}

	if short != nil {
		help.ShortDesc = *short
	}

	if long != nil {
		help.LongDesc = *long
	}

	return help
}
