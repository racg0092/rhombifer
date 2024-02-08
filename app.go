package rhombifer

import "sync"

type App struct {
}

var apponce sync.Once

// in case the application does not go the route of a single root command and has multiple root commands
// this will implement a way to have all those command join somehow
