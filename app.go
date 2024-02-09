package rhombifer

import (
	"os"
)

type App struct {
}

// var apponce sync.Once

// in case the application does not go the route of a single root command and has multiple root commands
// this will implement a way to have all those command join somehow

func Start() error {
	args := os.Args[1:]

	if len(args) == 0 {
		if root != nil {
			if err := root.ExecCommand("help", nil); err != nil {
				return err
			}
		}
		return nil
	}

	if root == nil {
		return nil
	}

	//todo: multiple roots command path structure needs to be implemented

	if root.Run == nil {
		return nil
	}

	if err := root.Run(args); err != nil {
		return err
	}

	return nil
}
