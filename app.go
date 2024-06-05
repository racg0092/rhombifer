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

	if len(args) == 0 {
		if root != nil {
			if err := ExecCommand("help"); err != nil {
				return err
			}
			return nil
		}
	}

	// Expections are to change the control structure to allow for multiple roots
	if root == nil {
		return fmt.Errorf("Root command expected got %v", root)
	}

	if err := ExecCommand(args[0], args[1:]...); err != nil {
		return err
	}
	return nil
}
