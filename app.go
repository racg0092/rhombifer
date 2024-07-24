package rhombifer

import (
	"fmt"
	"os"
)

type App struct {
}

// var apponce sync.Once

// Kick starts the CLI with certain expectations. For more flexebility handle the start of the application your self
// This may change if future versions
func Start() error {
	args := os.Args[1:]
	var cmd string

	// Expections are to change the control structure to allow for multiple roots. At that point this would change
	// to accomodate for multiple roots (Honestly not sure if this is even worth it)
	if root == nil {
		return fmt.Errorf("Root command expected got %v", root)
	}

	if len(args) == 0 && config.RunHelpIfNoInput {
		cmd = "help"
	} else if len(args) == 0 {
		cmd = ""
	} else if len(args) > 0 && IsFirstArgFlag(args[0]) {
		if !config.AllowFlagsInRoot {
			return fmt.Errorf("Root command does not allow flags. Please use a subcommand")
		}
		cmd = ""
	} else if len(args) > 0 {
		cmd = args[0]
		args = args[1:]
	}

	if err := ExecCommand(cmd, args...); err != nil {
		return err
	}

	return nil
}
