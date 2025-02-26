package main

import (
	"fmt"
	"github.com/racg0092/rhombifer/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("REPL for Rhombifer. Feel Free to test command input %s\n", u.Username)
	repl.Start(os.Stdin, os.Stdout)
}
