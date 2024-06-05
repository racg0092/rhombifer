package builtin

import (
	"github.com/racg0092/rhombifer"
	"github.com/racg0092/rhombifer/pkg/builtin"
)

func ExampleHelpCommand() {
	// get root command if any. Warning calling this function
	// creates root command
	root := rhombifer.Root()
	// add built in. can be added to any command
	root.AddSub(builtin.HelpCommand(nil, nil))
	// run help command
	root.Subs["help"].Run()
}
