package rhombifer

import "os"

type App struct {
}

// var apponce sync.Once

// in case the application does not go the route of a single root command and has multiple root commands
// this will implement a way to have all those command join somehow

func Start() {
	args := os.Args[1:]

	if len(args) == 0 {
		if root != nil {
			root.ExecCommand("help", nil)
		}
	}

}
