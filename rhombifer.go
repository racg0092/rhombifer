// A flexible and simple unopinionated library for cli tools
package rhombifer

import (
	"fmt"
	"github.com/racg0092/rhombifer/pkg/models"
	"sync"
)

var root *Command

var once sync.Once

// Takes in a pointer to `cmd` and sets it as the root command. The **root** command is only set once
// for the application runtime. It means that `root` will be set only the first time this funtion is call.
// Use it if you need to define the `root` command before initialization
func SetRoot(cmd *Command) {
	if cmd != nil {
		cmd.Root = true
		if cmd.Subs == nil {
			cmd.Subs = make(map[string]Command)
		}
		if cmd.Flags == nil {
			cmd.Flags = make([]models.Flag, 0)
		}
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
		c := Command{}
		SetRoot(&c)
	}
	return root
}

// Executes command passed in. It expects [root] to be set
func ExecCommand(cmd string, args ...string) error {
	root := Root()
	if root == nil {
		return fmt.Errorf("Expected root command to be set found %v", root)
	}
	if len(args) == 0 && root.Run != nil && cmd == "" {
		root.Run()
		return nil
	}
	subcommand, found := root.Subs[cmd]
	if !found {
		return fmt.Errorf("Command %s was not found", cmd)
	}
	if subcommand.Run == nil {
		return fmt.Errorf("Sub command %s, does not have a valid function (Run)", subcommand.Name)
	}

	if subcommand.RequiredFlags != nil {
		if len(args) == 0 {
			return fmt.Errorf("This command (%s) requires flags. Please check the commands docs", subcommand.Name)
		}
		//todo: check if the flags present are valid

	}
	return subcommand.Run(args...)
}
